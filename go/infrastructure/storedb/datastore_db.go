package storedb

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

type DataStoreService struct {
	client *datastore.Client
	ctx    context.Context
}

type Entity struct {
	Value     interface{}
	CreatedAt time.Duration
}

func NewDataStoreDB(projectID string) (*DataStoreService, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_ID"))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return &DataStoreService{client: client, ctx: ctx}, nil
}

func (ds *DataStoreService) NewKey(kind, name string) *datastore.Key {
	return datastore.NameKey(kind, name, nil)
}

// 分単位
// データを登録
func (ds *DataStoreService) Set(name string, value interface{}, ttl time.Duration) error {
	key := ds.NewKey(name, name)

	entity := Entity{
		Value:     value,
		CreatedAt: ttl,
	}

	_, err := ds.client.Put(ds.ctx, key, entity)
	if err != nil {
		return err
	}

	return nil
}

// データを取得
func (ds *DataStoreService) Get(name string) (interface{}, error) {
	key := ds.NewKey(name, name)
	entity := new(Entity)
	if err := ds.client.Get(ds.ctx, key, &entity); err != nil {
		return nil, err
	}

	return entity.Value, nil
}

// データを削除
func (ds *DataStoreService) Delete(name string) error {
	key := ds.NewKey(name, name)
	if err := ds.client.Delete(ds.ctx, key); err != nil {
		return err
	}
	return nil
}
