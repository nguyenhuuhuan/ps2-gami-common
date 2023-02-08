package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitlab.id.vin/gami/go-agent/v3/newrelic"

	"github.com/dgraph-io/ristretto"
	"github.com/vmihailenco/msgpack/v5"
)

type LocalCacheAdapter interface {
	Get(ctx context.Context, key string, v interface{}) error
	Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	Del(ctx context.Context, key []string) error
}
type localCacheAdapter struct {
	cache *ristretto.Cache
}

// NewLocalCacheAdapter returns a new instance of LocalCacheAdapter.
func NewLocalCacheAdapter() LocalCacheAdapter {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e6,     // number of keys to track frequency of (1M).
		MaxCost:     1 << 24, // maximum cost of cache (128 MB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	return &localCacheAdapter{
		cache: cache,
	}
}

func (r *localCacheAdapter) Get(ctx context.Context, key string, v interface{}) error {
	txn := newrelic.FromContext(ctx)
	if txn != nil {
		seg := txn.StartSegment("db.local-cache")
		seg.AddAttribute("query", fmt.Sprintf("GET key %v", key))
		defer seg.End()
	}

	data, found := r.cache.Get(key)
	if !found {
		return errors.New("value not found")
	}

	dataBytes, ok := data.([]byte)
	if ok {
		return msgpack.Unmarshal(dataBytes, v)
	}

	return errors.New("values not valid byte type")

}

func (r *localCacheAdapter) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	txn := newrelic.FromContext(ctx)
	if txn != nil {
		seg := txn.StartSegment("db.local-cache")
		seg.AddAttribute("query", fmt.Sprintf("SET key %v", key))
		defer seg.End()
	}

	data, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}

	isValid := r.cache.SetWithTTL(key, data, 1, expiration)
	if isValid {
		return nil
	}
	return errors.New("set key errors")
}

func (r *localCacheAdapter) Del(ctx context.Context, key []string) error {
	txn := newrelic.FromContext(ctx)
	if txn != nil {
		seg := txn.StartSegment("db.local-cache")
		seg.AddAttribute("query", fmt.Sprintf("SET key %v", key))
		defer seg.End()
	}

	for _, v := range key {
		r.cache.Del(v)
	}
	return nil
}
