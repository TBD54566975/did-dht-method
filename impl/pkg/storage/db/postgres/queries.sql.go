// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package postgres

import (
	"context"
)

const listRecords = `-- name: ListRecords :many
SELECT id, key, value, sig FROM pkarr_records
`

func (q *Queries) ListRecords(ctx context.Context) ([]PkarrRecord, error) {
	rows, err := q.db.Query(ctx, listRecords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PkarrRecord
	for rows.Next() {
		var i PkarrRecord
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Value,
			&i.Sig,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readRecord = `-- name: ReadRecord :one
SELECT id, key, value, sig FROM pkarr_records WHERE key = $1 LIMIT 1
`

func (q *Queries) ReadRecord(ctx context.Context, key string) (PkarrRecord, error) {
	row := q.db.QueryRow(ctx, readRecord, key)
	var i PkarrRecord
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Sig,
	)
	return i, err
}

const writeRecord = `-- name: WriteRecord :exec
INSERT INTO pkarr_records(key, value, sig) VALUES($1, $2, $3)
`

type WriteRecordParams struct {
	Key   string
	Value string
	Sig   string
}

func (q *Queries) WriteRecord(ctx context.Context, arg WriteRecordParams) error {
	_, err := q.db.Exec(ctx, writeRecord, arg.Key, arg.Value, arg.Sig)
	return err
}