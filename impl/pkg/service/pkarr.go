package service

import (
	"context"
	"errors"
	"time"

	"github.com/goccy/go-json"

	"github.com/TBD54566975/did-dht-method/internal/util"
	ssiutil "github.com/TBD54566975/ssi-sdk/util"
	"github.com/allegro/bigcache/v3"
	"github.com/anacrolix/dht/v2/bep44"
	"github.com/anacrolix/torrent/bencode"
	"github.com/sirupsen/logrus"

	"github.com/TBD54566975/did-dht-method/config"
	dhtint "github.com/TBD54566975/did-dht-method/internal/dht"
	"github.com/TBD54566975/did-dht-method/pkg/dht"
	"github.com/TBD54566975/did-dht-method/pkg/storage"
	"github.com/TBD54566975/did-dht-method/pkg/storage/pkarr"
)

const recordSizeLimit = 1000

// PkarrService is the Pkarr service responsible for managing the Pkarr DHT and reading/writing records
type PkarrService struct {
	cfg       *config.Config
	db        storage.Storage
	dht       *dht.DHT
	cache     *bigcache.BigCache
	scheduler *dhtint.Scheduler
}

// NewPkarrService returns a new instance of the Pkarr service
func NewPkarrService(cfg *config.Config, db storage.Storage) (*PkarrService, error) {
	if cfg == nil {
		return nil, ssiutil.LoggingNewError("config is required")
	}

	d, err := dht.NewDHT(cfg.DHTConfig.BootstrapPeers)
	if err != nil {
		return nil, ssiutil.LoggingErrorMsg(err, "failed to instantiate dht")
	}

	// create and start cache and scheduler
	cacheTTL := time.Duration(cfg.PkarrConfig.CacheTTLSeconds) * time.Second
	cacheConfig := bigcache.DefaultConfig(cacheTTL)
	cacheConfig.MaxEntrySize = recordSizeLimit
	cacheConfig.HardMaxCacheSize = cfg.PkarrConfig.CacheSizeLimitMB
	cacheConfig.CleanWindow = cacheTTL / 2
	cache, err := bigcache.New(context.Background(), cacheConfig)
	if err != nil {
		return nil, ssiutil.LoggingErrorMsg(err, "failed to instantiate cache")
	}
	scheduler := dhtint.NewScheduler()
	service := PkarrService{
		cfg:       cfg,
		db:        db,
		dht:       d,
		cache:     cache,
		scheduler: &scheduler,
	}
	if err = scheduler.Schedule(cfg.PkarrConfig.RepublishCRON, service.republish); err != nil {
		return nil, ssiutil.LoggingErrorMsg(err, "failed to start republisher")
	}
	return &service, nil
}

// PublishPkarrRequest is the request to publish a Pkarr record
type PublishPkarrRequest struct {
	V   []byte   `validate:"required"`
	K   [32]byte `validate:"required"`
	Sig [64]byte `validate:"required"`
	Seq int64    `validate:"required"`
}

// isValid returns an error if the request is invalid; also validates the signature
func (p PublishPkarrRequest) isValid() error {
	if err := ssiutil.IsValidStruct(p); err != nil {
		return err
	}
	// validate the signature
	bv, err := bencode.Marshal(p.V)
	if err != nil {
		return err
	}
	if !bep44.Verify(p.K[:], nil, p.Seq, bv, p.Sig[:]) {
		return errors.New("signature is invalid")
	}
	return nil
}

func (p PublishPkarrRequest) toRecord() pkarr.Record {
	return pkarr.Record{
		V:   p.V,
		K:   p.K[:],
		Sig: p.Sig[:],
		Seq: p.Seq,
	}
}

// PublishPkarr stores the record in the db, publishes the given Pkarr record to the DHT, and returns the z-base-32 encoded ID
func (s *PkarrService) PublishPkarr(ctx context.Context, id string, request PublishPkarrRequest) error {
	if err := request.isValid(); err != nil {
		return err
	}

	// write to db and cache
	record := request.toRecord()
	if err := s.db.WriteRecord(ctx, record); err != nil {
		return err
	}
	recordBytes, err := json.Marshal(GetPkarrResponse{
		V:   request.V,
		Seq: request.Seq,
		Sig: request.Sig,
	})
	if err != nil {
		return err
	}

	if err = s.cache.Set(id, recordBytes); err != nil {
		return err
	}

	// return here and put it in the DHT asynchronously
	// TODO(gabe): consider a background process to monitor failures
	go func() {
		_, err := s.dht.Put(ctx, bep44.Put{
			V:   request.V,
			K:   &request.K,
			Sig: request.Sig,
			Seq: request.Seq,
		})
		if err != nil {
			logrus.WithError(err).Error("error from dht.Put")
		}
	}()

	return nil
}

// GetPkarrResponse is the response to a get Pkarr request
type GetPkarrResponse struct {
	V   []byte   `validate:"required"`
	Seq int64    `validate:"required"`
	Sig [64]byte `validate:"required"`
}

func fromPkarrRecord(record pkarr.Record) (*GetPkarrResponse, error) {
	return &GetPkarrResponse{
		V:   record.V,
		Seq: record.Seq,
		Sig: [64]byte(record.Sig),
	}, nil
}

// GetPkarr returns the full Pkarr record (including sig data) for the given z-base-32 encoded ID
func (s *PkarrService) GetPkarr(ctx context.Context, id string) (*GetPkarrResponse, error) {
	// first do a cache lookup
	if got, err := s.cache.Get(id); err == nil {
		var resp GetPkarrResponse
		if err = json.Unmarshal(got, &resp); err != nil {
			return nil, err
		}
		logrus.Debugf("resolved pkarr record[%s] from cache", id)
		return &resp, nil
	}

	// next do a dht lookup
	got, err := s.dht.GetFull(ctx, id)
	if err != nil {
		// try to resolve from storage before returning and error
		logrus.WithError(err).WithField("record", id).Warn("failed to get pkarr record from dht, attempting to resolve from storage")

		rawID, err := util.Z32Decode(id)
		if err != nil {
			return nil, err
		}

		record, err := s.db.ReadRecord(ctx, rawID)
		if err != nil || record == nil {
			logrus.WithError(err).WithField("record", id).Error("failed to resolve pkarr record from storage")
			return nil, err
		}

		logrus.WithField("record", id).Debug("resolved pkarr record from storage")
		resp, err := fromPkarrRecord(*record)
		if err == nil {
			if err = s.addRecordToCache(id, *resp); err != nil {
				logrus.WithError(err).WithField("record", id).Error("failed to set pkarr record in cache")
			}
		}

		return resp, err
	}

	// prepare the record for return
	bBytes, err := got.V.MarshalBencode()
	if err != nil {
		return nil, err
	}
	var payload string
	if err = bencode.Unmarshal(bBytes, &payload); err != nil {
		return nil, err
	}
	resp := GetPkarrResponse{
		V:   []byte(payload),
		Seq: got.Seq,
		Sig: got.Sig,
	}

	// add the record to cache, do it here to avoid duplicate calculations
	if err = s.addRecordToCache(id, resp); err != nil {
		logrus.WithError(err).Errorf("failed to set pkarr record[%s] in cache", id)
	}

	return &resp, nil
}

func (s *PkarrService) addRecordToCache(id string, resp GetPkarrResponse) error {
	recordBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	if err = s.cache.Set(id, recordBytes); err != nil {
		return err
	}
	return nil
}

// TODO(gabe) make this more efficient. create a publish schedule based on each individual record, not all records
func (s *PkarrService) republish() {
	var nextPageToken []byte
	var allRecords []pkarr.Record
	var err error
	errCnt := 0
	successCnt := 0
	for {
		allRecords, nextPageToken, err = s.db.ListRecords(context.Background(), nextPageToken, 1000)
		if err != nil {
			logrus.WithError(err).Error("failed to list record(s) for republishing")
			return
		}

		if len(allRecords) == 0 {
			logrus.Info("No records to republish")
			return
		}

		logrus.WithField("record_count", len(allRecords)).Info("Republishing record")

		for _, record := range allRecords {
			put, err := recordToBEP44Put(record)
			if err != nil {
				logrus.WithError(err).Error("failed to convert record to bep44 put")
				errCnt++
				continue
			}

			if _, err = s.dht.Put(context.Background(), *put); err != nil {
				logrus.WithError(err).Error("failed to republish record")
				errCnt++
				continue
			}

			successCnt++
		}

		if nextPageToken == nil {
			break
		}
	}

	logrus.WithField("success", successCnt).WithField("errors", errCnt).Info("Republishing complete")
}

func recordToBEP44Put(record pkarr.Record) (*bep44.Put, error) {
	return &bep44.Put{
		V:   record.V,
		K:   (*[32]byte)(record.K),
		Sig: [64]byte(record.Sig),
		Seq: record.Seq,
	}, nil
}
