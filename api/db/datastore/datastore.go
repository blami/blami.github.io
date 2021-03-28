// Adapter for Firebase in Datastore mode with map cache in front of read to
// better manage free quota in case only one or zero instances are running.
package datastore

import (
	"context"

	"github.com/blami/blami.github.io/api/log"

	"cloud.google.com/go/datastore"
)

type Value struct {
	Value string
}

type db struct {
	client *datastore.Client
}

func New() *db {
	// Create datastore client
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to datastore")
	}
	return &db{
		client: client,
	}
}

func (d *db) Set(ctx context.Context, key string, value string) error {
	k := datastore.NameKey("Value", key, nil)
	v := &Value{
		Value: value,
	}

	if _, err := d.client.Put(ctx, k, v); err != nil {
		log.WithError(err).Error("unable to set")
		return err
	}
	return nil
}

func (d *db) Get(ctx context.Context, key string) (string, error) {
	k := datastore.NameKey("Value", key, nil)
	v := new(Value)

	if err := d.client.Get(ctx, k, v); err != nil {
		log.WithError(err).Error("unable to get")
		return "", err
	}
	return v.Value, nil
}
