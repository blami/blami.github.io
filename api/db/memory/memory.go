// Non-production adapter for database implemented using string map in memory.
package memory

import "context"

type db struct {
	data map[string]string
}

func New() *db {
	return &db{
		data: map[string]string{},
	}
}

func (d *db) Set(ctx context.Context, key string, value string) error {
	d.data[key] = value
	return nil
}

func (d *db) Get(ctx context.Context, key string) (string, error) {
	if val, ok := d.data[key]; ok {
		return val, nil
	}
	return "", nil
}
