package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TBD54566975/did-dht-method/config"
	"github.com/TBD54566975/did-dht-method/internal/did"
	"github.com/TBD54566975/did-dht-method/pkg/dht"
	"github.com/TBD54566975/did-dht-method/pkg/service"
	"github.com/TBD54566975/did-dht-method/pkg/storage"
)

func TestPKARRRouter(t *testing.T) {
	pkarrSvc := testPKARRService(t)
	pkarrRouter, err := NewPkarrRouter(&pkarrSvc)
	require.NoError(t, err)
	require.NotEmpty(t, pkarrRouter)

	t.Run("test put record", func(t *testing.T) {
		didID, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		suffix, err := did.DHT(didID).Suffix()
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.True(t, is2xxResponse(w.Code), "unexpected %s", w.Result().Status)
	})

	t.Run("test get record", func(t *testing.T) {
		didID, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		suffix, err := did.DHT(didID).Suffix()
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.True(t, is2xxResponse(w.Code), "unexpected %s", w.Result().Status)

		w = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", testServerURL, suffix), nil)
		c = newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.GetRecord(c)
		assert.True(t, is2xxResponse(w.Code), "unexpected %s", w.Result().Status)

		resp, err := io.ReadAll(w.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
		assert.Equal(t, reqData, resp)
	})

	t.Run("test get no ID", func(t *testing.T) {
		didID, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		suffix, err := did.DHT(didID).Suffix()
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.True(t, is2xxResponse(w.Code), "unexpected %s", w.Result().Status)

		w = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", testServerURL, suffix), nil)
		c = newRequestContextWithParams(w, req, map[string]string{})

		pkarrRouter.GetRecord(c)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})

	t.Run("test put no ID", func(t *testing.T) {
		_, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/", testServerURL), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{})

		pkarrRouter.PutRecord(c)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})

	t.Run("test put undecodable ID", func(t *testing.T) {
		_, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		suffix := "----"
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})

	t.Run("test put invalid record signature", func(t *testing.T) {
		didID, reqData := generateDIDPutRequest(t)

		reqData = append(reqData, []byte{1, 2, 3, 4, 5}...) // append some garbage to the request body, making the signature invalid

		w := httptest.NewRecorder()
		suffix, err := did.DHT(didID).Suffix()
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})

	t.Run("test put invalid key ID", func(t *testing.T) {
		_, reqData := generateDIDPutRequest(t)

		w := httptest.NewRecorder()
		suffix := "aaaa"
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", testServerURL, suffix), bytes.NewReader(reqData))
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})

		pkarrRouter.PutRecord(c)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})

	t.Run("test get not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		suffix := "aaa"
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", testServerURL, suffix), nil)
		c := newRequestContextWithParams(w, req, map[string]string{IDParam: suffix})
		pkarrRouter.GetRecord(c)
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode, "unexpected %s", w.Result().Status)
	})
}

func testPKARRService(t *testing.T) service.PkarrService {
	defaultConfig := config.GetDefaultConfig()

	db, err := storage.NewStorage(defaultConfig.ServerConfig.StorageURI)
	require.NoError(t, err)
	require.NotEmpty(t, db)

	d := dht.NewTestDHT(t)

	pkarrService, err := service.NewPkarrService(&defaultConfig, db, d)
	require.NoError(t, err)
	require.NotEmpty(t, pkarrService)

	return *pkarrService
}

func generateDIDPutRequest(t *testing.T) (string, []byte) {
	// generate a DID Document
	sk, doc, err := did.GenerateDIDDHT(did.CreateDIDDHTOpts{})
	require.NoError(t, err)
	require.NotEmpty(t, doc)

	packet, err := did.DHT(doc.ID).ToDNSPacket(*doc, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, packet)

	bep44Put, err := dht.CreatePKARRPublishRequest(sk, *packet)
	assert.NoError(t, err)
	assert.NotEmpty(t, bep44Put)

	// prepare request as sig:seq:v
	var seqBuf [8]byte
	binary.BigEndian.PutUint64(seqBuf[:], uint64(bep44Put.Seq))
	return doc.ID, append(bep44Put.Sig[:], append(seqBuf[:], bep44Put.V.([]byte)...)...)
}
