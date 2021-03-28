// Package db defines interface of key-value DB adapter.
package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/blami/blami.github.io/api/db/datastore"
	"github.com/blami/blami.github.io/api/db/memory"
)

// Interface that must be implemented by each storage adapter so it can be used
// to store data by apisvc. Most of methods take Context as first argument. It
// is used to pass correlation values (e.g. requestId from web API, etc.)
type DB interface {
	// Set given key to given value, if key doesn't exist create it otherwise
	// overwrite it with new value. Return error in case of failure.
	Set(ctx context.Context, key string, value string) error

	// Get value of given key. If key doesn't exist return an error.
	Get(ctx context.Context, key string) (string, error)
}

// Given connection string creates appropriate DB adapter that implements DB
// interface.
func MakeDB(conn string) (DB, error) {
	opts := strings.SplitN(conn, ":", 1)

	var db DB
	switch opts[0] {
	case "memory":
		db = memory.New()
		return db, nil
	case "datastore":
		db = datastore.New()
		return db, nil
	}

	return nil, fmt.Errorf("unknown database driver: %s", opts[0])
}
