package cache

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	shortTerm = 1000 * time.Millisecond
	midTerm   = 3000 * time.Millisecond
	longTerm  = 5000 * time.Millisecond
)

// CachedV2Adapter provides access to a Redis caching database.
type CachedV2Adapter interface {
	Get(ctx context.Context, key string, v interface{}) error
	MGet(ctx context.Context, keys []string) ([]interface{}, error)
	Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	MSet(ctx context.Context, m map[string]interface{}) error
	MSetNX(ctx context.Context, m map[string]interface{}) (bool, error)
	Del(ctx context.Context, keys ...string) error
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, quantity int64) (int64, error)
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	GetInt64(ctx context.Context, key string) (int64, error)
	SAdd(ctx context.Context, key string, list []string) error
	SMem(ctx context.Context, key string, value string) bool
	SMMem(ctx context.Context, key string, list []string) []bool
	SCard(ctx context.Context, key string) int64
	SRem(ctx context.Context, key string, list []string) error
	PFAdd(ctx context.Context, key string, list []string) error
	PFCount(ctx context.Context, key string) int64
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HMSet(ctx context.Context, key string, fields interface{}) error
	HMGet(ctx context.Context, key string, field []string) ([]interface{}, error)
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key string, member string, v interface{}) error
	HDel(ctx context.Context, key, member string) error
	HLen(ctx context.Context, key string) (int64, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	HSetInt64(ctx context.Context, key, member string, value int64) error
	HGetInt64(ctx context.Context, key, member string) (int64, error)
	HExists(ctx context.Context, key, field string) bool
	HExistsV2(ctx context.Context, key, field string) (bool, error)
	Exists(ctx context.Context, key string) bool
	ExistsV2(ctx context.Context, keys ...string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	HIncr(ctx context.Context, key string, member string, quantity int64) (int64, error)
	Flush(ctx context.Context) error
	ZAdd(ctx context.Context, key string, v *redis.Z) error
	ZScan(ctx context.Context, key string, match string) ([]string, error)
	ZAdds(ctx context.Context, key string, v ...*redis.Z) error
	ZRem(ctx context.Context, key string, member interface{}) error
	ZIncr(ctx context.Context, key string, member *redis.Z) error
	ZRevRange(ctx context.Context, key string, min, max int64) ([]redis.Z, error)
	ZRevRangeByScore(ctx context.Context, key string, min string, max string, offset, count int64) ([]redis.Z, error)
	ZScore(ctx context.Context, key string, member string) (float64, error)
	ZMScore(ctx context.Context, key string, members ...string) ([]float64, error)
	ZRank(ctx context.Context, key string, member string) (int64, error)
	ZRevRank(ctx context.Context, key string, member string) (int64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key string, min, max int64) (int64, error)
	LPush(ctx context.Context, key string, val ...interface{}) error
	LLength(ctx context.Context, key string) (int64, error)
	LRem(ctx context.Context, key string, val interface{}) (int64, error)
	RPop(ctx context.Context, key string) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	BitField(ctx context.Context, key string, args []BitFieldModel) ([]int64, error)
	BitCount(ctx context.Context, key string) (int64, error)
	BitCountAll(ctx context.Context, key string) (int64, error)
	BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpNot(ctx context.Context, destKey string, key string) (int64, error)
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error)
	StrLen(ctx context.Context, key string) (int64, error)
	GetBytes(ctx context.Context, key string) ([]byte, error)
	RPush(ctx context.Context, key string, val ...interface{}) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	SetBit(ctx context.Context, key string, offset int64, value int) error
	GetBit(ctx context.Context, key string, offset int64) (int64, error)
	RSetBit(ctx context.Context, key string, offset int64, value int) (int64, error)
	RGetBit(ctx context.Context, key string, offset int64) (int64, error)
	RGetBits(ctx context.Context, keys []string, offset int64) ([]int, error)
	RBitCount(ctx context.Context, key string) (int64, error)
	RSetIntArray(ctx context.Context, key string, value []int64) (bool, error)
	RGetIntArray(ctx context.Context, key string) ([]int64, error)
	RAppendIntArray(ctx context.Context, key string, value []int64) (bool, error)
	AppendIntArray(ctx context.Context, key string, value []int64) (bool, error)
	ROptimize(ctx context.Context, key string) (bool, error)
}

type cachedV2Adapter struct {
	client *redis.Client
}

// NewCachedAdapter returns a new instance of CachedAdapter.
func NewCacheV2Adapter(client *redis.Client) CachedV2Adapter {
	return &cachedV2Adapter{
		client: client,
	}
}

func (r *cachedV2Adapter) Get(ctx context.Context, key string, v interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return msgpack.Unmarshal(data, v)
}

func (r *cachedV2Adapter) MGet(ctx context.Context, keys []string) ([]interface{}, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.MGet(ctx, keys...).Result()
}

func (r *cachedV2Adapter) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	data, err := msgpack.Marshal(v)
	if err != nil {
		return nil
	}

	_, err = r.client.Set(ctx, key, data, expiration).Result()
	return err
}

func (r *cachedV2Adapter) MSet(ctx context.Context, m map[string]interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()

	return r.client.MSet(ctx, m).Err()
}

func (r *cachedV2Adapter) MSetNX(ctx context.Context, m map[string]interface{}) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()

	return r.client.MSetNX(ctx, m).Result()
}

func (r *cachedV2Adapter) Del(ctx context.Context, keys ...string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	_, err := r.client.Del(ctx, keys...).Result()
	return err
}

func (r *cachedV2Adapter) Incr(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	cmd := r.client.Incr(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (r *cachedV2Adapter) IncrBy(ctx context.Context, key string, quantity int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.IncrBy(ctx, key, quantity).Result()
}

func (r *cachedV2Adapter) SetString(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	_, err := r.client.Set(ctx, key, v, expiration).Result()
	return err
}

func (r *cachedV2Adapter) GetString(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	data, err := r.client.Get(ctx, key).Bytes()
	return string(data), err
}

func (r *cachedV2Adapter) GetInt64(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(string(data), 10, 64)
}

func (r *cachedV2Adapter) SAdd(ctx context.Context, key string, list []string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.SAdd(ctx, key, list).Result()
	return err
}

func (r *cachedV2Adapter) SMem(ctx context.Context, key string, value string) bool {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	found, _ := r.client.SIsMember(ctx, key, value).Result()
	return found
}

func (r *cachedV2Adapter) SMMem(ctx context.Context, key string, listValue []string) []bool {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	found, _ := r.client.SMIsMember(ctx, key, listValue).Result()
	return found
}

func (r *cachedV2Adapter) SCard(ctx context.Context, key string) int64 {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	data, _ := r.client.SCard(ctx, key).Result()
	return data
}

func (r *cachedV2Adapter) SRem(ctx context.Context, key string, list []string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.SRem(ctx, key, list).Result()
	return err
}

func (r *cachedV2Adapter) PFAdd(ctx context.Context, key string, list []string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.PFAdd(ctx, key, list).Result()
	return err
}

func (r *cachedV2Adapter) PFCount(ctx context.Context, key string) int64 {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	data, _ := r.client.PFCount(ctx, key).Result()
	return data
}

// HGet function
func (r *cachedV2Adapter) HGet(ctx context.Context, key string, member string, v interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	data, err := r.client.HGet(ctx, key, member).Bytes()
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(data, v)
}

func (r *cachedV2Adapter) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.HGetAll(ctx, key).Result()
}

func (r *cachedV2Adapter) HMSet(ctx context.Context, key string, fields interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	var redisData []interface{}
	v := reflect.ValueOf(fields)
	data := map[string]interface{}{}

	if v.Kind() == reflect.Map {
		for _, item := range v.MapKeys() {
			myStruct := v.MapIndex(item)
			myKey := item.String()
			data[myKey] = myStruct.Interface()
		}
	} else {
		return errors.New("fields not map type")
	}

	for key, item := range data {
		redisData = append(redisData, key)
		bytes, err := msgpack.Marshal(item)
		if err != nil {
			return err
		}
		redisData = append(redisData, bytes)

	}

	return r.client.HMSet(ctx, key, redisData...).Err()
}

func (r *cachedV2Adapter) HMGet(ctx context.Context, key string, field []string) ([]interface{}, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	result, err := r.client.HMGet(ctx, key, field...).Result()
	return result, err
}

func (r *cachedV2Adapter) HSet(ctx context.Context, key, field string, value interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	data, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	_, err = r.client.HSet(ctx, key, field, data).Result()
	return err
}

func (r *cachedV2Adapter) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return r.client.HIncrBy(ctx, key, field, incr).Result()
}

func (r *cachedV2Adapter) HGetInt64(ctx context.Context, key, member string) (int64, error) {
	data, err := r.client.HGet(ctx, key, member).Int64()
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (r *cachedV2Adapter) HSetInt64(ctx context.Context, key, member string, value int64) error {
	err := r.client.HSet(ctx, key, member, value).Err()
	return err
}

func (r *cachedV2Adapter) HExists(ctx context.Context, key, field string) bool {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.HExists(ctx, key, field).Val()
}

func (r *cachedV2Adapter) HExistsV2(ctx context.Context, key, field string) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.HExists(ctx, key, field).Result()
}

func (r *cachedV2Adapter) HLen(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.HLen(ctx, key).Result()
}

func (r *cachedV2Adapter) Exists(ctx context.Context, key string) bool {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.Exists(ctx, []string{key}...).Val() > 0
}

func (r *cachedV2Adapter) ExistsV2(ctx context.Context, keys ...string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.Exists(ctx, keys...).Result()
}

func (r *cachedV2Adapter) Expire(ctx context.Context, key string, expiration time.Duration) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.Expire(ctx, key, expiration).Err()
}

func (r *cachedV2Adapter) HIncr(ctx context.Context, key string, member string, quantity int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.HIncrBy(ctx, key, member, quantity).Result()
}

func (r *cachedV2Adapter) Flush(ctx context.Context) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	_, err := r.client.FlushAll(ctx).Result()
	return err
}

func (r *cachedV2Adapter) ZAdd(ctx context.Context, key string, v *redis.Z) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.ZAdd(ctx, key, v).Result()
	return err
}

func (r *cachedV2Adapter) ZScan(ctx context.Context, key string, match string) ([]string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	data, _, err := r.client.ZScan(ctx, key, 0, match, 1).Result()
	return data, err
}

func (r *cachedV2Adapter) ZAdds(ctx context.Context, key string, v ...*redis.Z) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.ZAdd(ctx, key, v...).Result()
	return err
}

func (r *cachedV2Adapter) ZRem(ctx context.Context, key string, member interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZRem(ctx, key, member).Err()
}

func (r *cachedV2Adapter) ZIncr(ctx context.Context, key string, member *redis.Z) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	_, err := r.client.ZAddArgsIncr(ctx, key, redis.ZAddArgs{
		Members: []redis.Z{*member},
	}).Result()
	return err
}

func (r *cachedV2Adapter) ZRevRange(ctx context.Context, key string, min, max int64) ([]redis.Z, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	data, err := r.client.ZRevRangeWithScores(ctx, key, min, max).Result()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *cachedV2Adapter) ZRevRangeByScore(ctx context.Context, key string, min string, max string, offset int64, count int64) ([]redis.Z, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	data, err := r.client.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *cachedV2Adapter) ZScore(ctx context.Context, key string, member string) (float64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZScore(ctx, key, member).Result()
}

func (r *cachedV2Adapter) ZMScore(ctx context.Context, key string, members ...string) ([]float64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZMScore(ctx, key, members...).Result()
}

func (r *cachedV2Adapter) ZRank(ctx context.Context, key string, member string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZRank(ctx, key, member).Result()
}

func (r *cachedV2Adapter) ZRevRank(ctx context.Context, key string, member string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZRevRank(ctx, key, member).Result()
}

func (r *cachedV2Adapter) ZCard(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.ZCard(ctx, key).Result()
}

func (r *cachedV2Adapter) ZCount(ctx context.Context, key string, min, max int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	smin := strconv.FormatInt(min, 10)
	smax := strconv.FormatInt(max, 10)
	return r.client.ZCount(ctx, key, smin, smax).Result()
}

// LRange return a list with range in List
func (r *cachedV2Adapter) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	stuff := r.client.LRange(ctx, key, start, stop)
	return stuff.Val(), stuff.Err()
}

// RPop to remove last element in List
func (r *cachedV2Adapter) RPop(ctx context.Context, key string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.RPop(ctx, key).Err()
}

// LLength function return length a List
func (r *cachedV2Adapter) LLength(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	length := r.client.LLen(ctx, key)
	return length.Val(), length.Err()
}

// LPush push values to a List
func (r *cachedV2Adapter) LPush(ctx context.Context, key string, val ...interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.LPush(ctx, key, val...).Err()
}

func (r *cachedV2Adapter) LRem(ctx context.Context, key string, val interface{}) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.LRem(ctx, key, 0, val).Result()
}

func (r *cachedV2Adapter) SetBit(ctx context.Context, key string, offset int64, value int) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.SetBit(ctx, key, offset, value).Err()
}

func (r *cachedV2Adapter) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	return r.client.GetBit(ctx, key, offset).Result()
}

func (r *cachedV2Adapter) GetBytes(ctx context.Context, key string) ([]byte, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.Get(ctx, key).Bytes()
}

func (r *cachedV2Adapter) SetBytes(ctx context.Context, key string, bytes []byte, duration time.Duration) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.Set(ctx, key, bytes, duration).Err()
}

func (r *cachedV2Adapter) SetMultiBit(ctx context.Context, key string, args ...interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitField(ctx, key, args).Err()
}

func (r *cachedV2Adapter) BitField(ctx context.Context, key string, args []BitFieldModel) ([]int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	var data []interface{}
	for _, item := range args {
		data = append(data, item.convertCommand()...)
	}

	return r.client.BitField(ctx, key, data...).Result()
}

func (r *cachedV2Adapter) BitCount(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitCount(ctx, key, nil).Result()
}

func (r *cachedV2Adapter) BitCountAll(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitCount(ctx, key, nil).Result()
}

func (r *cachedV2Adapter) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitOpAnd(ctx, destKey, keys...).Result()
}

func (r *cachedV2Adapter) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitOpOr(ctx, destKey, keys...).Result()
}

func (r *cachedV2Adapter) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitOpXor(ctx, destKey, keys...).Result()
}

func (r *cachedV2Adapter) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitOpNot(ctx, destKey, key).Result()
}

func (r *cachedV2Adapter) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.BitPos(ctx, key, bit, pos...).Result()
}

func (r *cachedV2Adapter) StrLen(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	return r.client.StrLen(ctx, key).Result()
}

func (r *cachedV2Adapter) HDel(ctx context.Context, key, member string) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	return r.client.HDel(ctx, key, member).Err()
}

func (r *cachedV2Adapter) RPush(ctx context.Context, key string, val ...interface{}) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	return r.client.RPush(ctx, key, val...).Err()
}

func (r *cachedV2Adapter) Keys(ctx context.Context, pattern string) ([]string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(shortTerm))
	defer cancel()
	return r.client.Keys(ctx, pattern).Result()
}

func (r *cachedV2Adapter) RSetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	cmd := redis.NewIntCmd(ctx,
		"R.SETBIT",
		key,
		offset,
		value,
	)

	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) RGetBit(ctx context.Context, key string, offset int64) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	cmd := redis.NewIntCmd(ctx,
		"R.GETBIT",
		key,
		offset,
	)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) RGetBits(ctx context.Context, keys []string, offset int64) ([]int, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	pipe := r.client.Pipeline()
	cmds := []*redis.Cmd{}
	for _, key := range keys {
		cmds = append(cmds, pipe.Do(ctx, "R.GETBIT", key, offset))
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	ans := []int{}
	for _, v := range cmds {
		rs, _ := v.Int()
		ans = append(ans, rs)
	}
	return ans, nil
}

func (r *cachedV2Adapter) RBitCount(ctx context.Context, key string) (int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	cmd := redis.NewIntCmd(ctx,
		"R.BITCOUNT",
		key,
	)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()

}

func (r *cachedV2Adapter) RSetIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	args := make([]interface{}, 2+len(value))
	args[0] = "R.SETINTARRAY"
	args[1] = key

	for i, key := range value {
		args[2+i] = key
	}

	cmd := redis.NewBoolCmd(ctx, args...)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) RAppendIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	args := make([]interface{}, 2+len(value))
	args[0] = "R.APPENDINTARRAY"
	args[1] = key

	for i, key := range value {
		args[2+i] = key
	}

	cmd := redis.NewBoolCmd(ctx, args...)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) AppendIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(longTerm))
	defer cancel()
	args := make([]interface{}, 2+len(value))
	args[0] = "APPEND"
	args[1] = key

	for i, key := range value {
		args[2+i] = key
	}

	cmd := redis.NewBoolCmd(ctx, args...)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) RGetIntArray(ctx context.Context, key string) ([]int64, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	cmd := redis.NewIntSliceCmd(ctx, "R.GETINTARRAY", key)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cachedV2Adapter) ROptimize(ctx context.Context, key string) (bool, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(midTerm))
	defer cancel()
	cmd := redis.NewBoolCmd(ctx, "R.OPTIMIZE", key)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}
