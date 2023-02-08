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

// ClusterAdapter provides access to a Redis caching database.
type ClusterAdapter interface {
	Get(ctx context.Context, key string, v interface{}) error
	Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, quantity int64) (int64, error)
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	GetInt64(ctx context.Context, key string) (int64, error)
	SAdd(ctx context.Context, key string, list []string) error
	SMem(ctx context.Context, key string, value string) bool
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
	HIncrBy(ctx context.Context, key, field string, incr int64) error
	HExists(ctx context.Context, key, field string) bool
	HExistsV2(ctx context.Context, key, field string) (bool, error)
	Exists(ctx context.Context, key string) bool
	Expire(ctx context.Context, key string, expiration time.Duration) error
	HIncr(ctx context.Context, key string, member string, quantity int64) (int64, error)
	Flush(ctx context.Context) error
	ZAdd(ctx context.Context, key string, v *redis.Z) error
	ZAdds(ctx context.Context, key string, v ...*redis.Z) error
	ZIncr(ctx context.Context, key string, member *redis.Z) error
	ZRevRange(ctx context.Context, key string, max int64) ([]redis.Z, error)
	ZScore(ctx context.Context, key string, member string) (float64, error)
	ZRank(ctx context.Context, key string, member string) (int64, error)
	ZRevRank(ctx context.Context, key string, member string) (int64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	LPush(ctx context.Context, key string, val ...interface{}) error
	LLength(ctx context.Context, key string) (int64, error)
	RPop(ctx context.Context, key string) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	BitField(ctx context.Context, key string, args []BitFieldModel) ([]int64, error)
	BitCount(ctx context.Context, key string) (int64, error)
	BitCountAll(ctx context.Context, key string) (int64, error)
	GetBytes(ctx context.Context, key string) ([]byte, error)
	RPush(ctx context.Context, key string, val ...interface{}) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	SetBit(ctx context.Context, key string, offset int64, value int) error
	GetBit(ctx context.Context, key string, offset int64) (int64, error)
	RSetBit(ctx context.Context, key string, offset int64, value int) (int64, error)
	RGetBit(ctx context.Context, key string, offset int64) (int64, error)
	RBitCount(ctx context.Context, key string) (int64, error)
	RSetIntArray(ctx context.Context, key string, value []int64) (bool, error)
	RGetIntArray(ctx context.Context, key string) ([]int64, error)
	RAppendIntArray(ctx context.Context, key string, value []int64) (bool, error)
	ROptimize(ctx context.Context, key string) (bool, error)
}

type cacheClusterAdapter struct {
	client *redis.ClusterClient
}

// NewCacheClusterAdapter returns a new instance of ClusterAdapter.
func NewCacheClusterAdapter(client *redis.ClusterClient) ClusterAdapter {
	return &cacheClusterAdapter{
		client: client,
	}
}

func (r *cacheClusterAdapter) Get(ctx context.Context, key string, v interface{}) error {

	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return msgpack.Unmarshal(data, v)
}

func (r *cacheClusterAdapter) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	data, err := msgpack.Marshal(v)
	if err != nil {
		return nil
	}

	_, err = r.client.Set(ctx, key, data, expiration).Result()
	return err
}

func (r *cacheClusterAdapter) Del(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}

func (r *cacheClusterAdapter) Incr(ctx context.Context, key string) (int64, error) {
	cmd := r.client.Incr(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (r *cacheClusterAdapter) IncrBy(ctx context.Context, key string, quantity int64) (int64, error) {
	return r.client.IncrBy(ctx, key, quantity).Result()
}

func (r *cacheClusterAdapter) SetString(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	_, err := r.client.Set(ctx, key, v, expiration).Result()
	return err
}

func (r *cacheClusterAdapter) GetString(ctx context.Context, key string) (string, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	return string(data), err
}

func (r *cacheClusterAdapter) GetInt64(ctx context.Context, key string) (int64, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(string(data), 10, 64)
}

func (r *cacheClusterAdapter) SAdd(ctx context.Context, key string, list []string) error {
	_, err := r.client.SAdd(ctx, key, list).Result()
	return err
}

func (r *cacheClusterAdapter) SMem(ctx context.Context, key string, value string) bool {
	found, _ := r.client.SIsMember(ctx, key, value).Result()
	return found
}

func (r *cacheClusterAdapter) SCard(ctx context.Context, key string) int64 {
	data, _ := r.client.SCard(ctx, key).Result()
	return data
}

func (r *cacheClusterAdapter) SRem(ctx context.Context, key string, list []string) error {
	_, err := r.client.SRem(ctx, key, list).Result()
	return err
}

func (r *cacheClusterAdapter) PFAdd(ctx context.Context, key string, list []string) error {
	_, err := r.client.PFAdd(ctx, key, list).Result()
	return err
}

func (r *cacheClusterAdapter) PFCount(ctx context.Context, key string) int64 {
	data, _ := r.client.PFCount(ctx, key).Result()
	return data
}

// HGet function
func (r *cacheClusterAdapter) HGet(ctx context.Context, key string, member string, v interface{}) error {
	data, err := r.client.HGet(ctx, key, member).Bytes()
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(data, v)
}

func (r *cacheClusterAdapter) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (r *cacheClusterAdapter) HMSet(ctx context.Context, key string, fields interface{}) error {
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

func (r *cacheClusterAdapter) HMGet(ctx context.Context, key string, field []string) ([]interface{}, error) {
	result, err := r.client.HMGet(ctx, key, field...).Result()
	return result, err
}

func (r *cacheClusterAdapter) HSet(ctx context.Context, key, field string, value interface{}) error {
	data, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	_, err = r.client.HSet(ctx, key, field, data).Result()
	return err
}

func (r *cacheClusterAdapter) HIncrBy(ctx context.Context, key, field string, incr int64) error {
	return r.client.HIncrBy(ctx, key, field, incr).Err()
}

func (r *cacheClusterAdapter) HExists(ctx context.Context, key, field string) bool {
	return r.client.HExists(ctx, key, field).Val()
}

func (r *cacheClusterAdapter) HExistsV2(ctx context.Context, key, field string) (bool, error) {
	return r.client.HExists(ctx, key, field).Result()
}

func (r *cacheClusterAdapter) Exists(ctx context.Context, key string) bool {
	return r.client.Exists(ctx, []string{key}...).Val() > 0
}

func (r *cacheClusterAdapter) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

func (r *cacheClusterAdapter) HIncr(ctx context.Context, key string, member string, quantity int64) (int64, error) {
	return r.client.HIncrBy(ctx, key, member, quantity).Result()
}

func (r *cacheClusterAdapter) Flush(ctx context.Context) error {
	_, err := r.client.FlushAll(ctx).Result()
	return err
}

func (r *cacheClusterAdapter) ZAdd(ctx context.Context, key string, v *redis.Z) error {
	_, err := r.client.ZAdd(ctx, key, v).Result()
	return err
}

func (r *cacheClusterAdapter) ZAdds(ctx context.Context, key string, v ...*redis.Z) error {
	_, err := r.client.ZAdd(ctx, key, v...).Result()
	return err
}

func (r *cacheClusterAdapter) ZIncr(ctx context.Context, key string, member *redis.Z) error {
	_, err := r.client.ZIncr(ctx, key, member).Result()
	return err
}

func (r *cacheClusterAdapter) ZRevRange(ctx context.Context, key string, max int64) ([]redis.Z, error) {
	data, err := r.client.ZRevRangeWithScores(ctx, key, 0, max).Result()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *cacheClusterAdapter) ZScore(ctx context.Context, key string, member string) (float64, error) {
	return r.client.ZScore(ctx, key, member).Result()
}

func (r *cacheClusterAdapter) ZRank(ctx context.Context, key string, member string) (int64, error) {
	return r.client.ZRank(ctx, key, member).Result()
}

func (r *cacheClusterAdapter) ZRevRank(ctx context.Context, key string, member string) (int64, error) {
	return r.client.ZRevRank(ctx, key, member).Result()
}

func (r *cacheClusterAdapter) ZCard(ctx context.Context, key string) (int64, error) {
	return r.client.ZCard(ctx, key).Result()
}

// LRange return a list with range in List
func (r *cacheClusterAdapter) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	stuff := r.client.LRange(ctx, key, start, stop)
	return stuff.Val(), stuff.Err()
}

// RPop to remove last element in List
func (r *cacheClusterAdapter) RPop(ctx context.Context, key string) error {
	return r.client.RPop(ctx, key).Err()
}

// LLength function return length a List
func (r *cacheClusterAdapter) LLength(ctx context.Context, key string) (int64, error) {
	length := r.client.LLen(ctx, key)
	return length.Val(), length.Err()
}

// LPush push values to a List
func (r *cacheClusterAdapter) LPush(ctx context.Context, key string, val ...interface{}) error {
	return r.client.LPush(ctx, key, val...).Err()
}

func (r *cacheClusterAdapter) SetBit(ctx context.Context, key string, offset int64, value int) error {
	return r.client.SetBit(ctx, key, offset, value).Err()
}

func (r *cacheClusterAdapter) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return r.client.GetBit(ctx, key, offset).Result()
}

func (r *cacheClusterAdapter) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return r.client.Get(ctx, key).Bytes()
}

func (r *cacheClusterAdapter) SetBytes(ctx context.Context, key string, bytes []byte, duration time.Duration) error {
	return r.client.Set(ctx, key, bytes, duration).Err()
}

func (r *cacheClusterAdapter) SetMultiBit(ctx context.Context, key string, args ...interface{}) error {
	return r.client.BitField(ctx, key, args).Err()
}

func (r *cacheClusterAdapter) BitField(ctx context.Context, key string, args []BitFieldModel) ([]int64, error) {
	var data []interface{}
	for _, item := range args {
		data = append(data, item.convertCommand()...)
	}

	return r.client.BitField(ctx, key, data...).Result()
}

func (r *cacheClusterAdapter) BitCount(ctx context.Context, key string) (int64, error) {
	return r.client.BitCount(ctx, key, nil).Result()
}

func (r *cacheClusterAdapter) BitCountAll(ctx context.Context, key string) (int64, error) {
	return r.client.BitCount(ctx, key, nil).Result()
}

func (r *cacheClusterAdapter) HDel(ctx context.Context, key, member string) error {
	return r.client.HDel(ctx, key, member).Err()
}

func (r *cacheClusterAdapter) RPush(ctx context.Context, key string, val ...interface{}) error {
	return r.client.RPush(ctx, key, val...).Err()
}

func (r *cacheClusterAdapter) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}

func (r *cacheClusterAdapter) RSetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {

	cmd := redis.NewIntCmd(ctx,
		"R.SETBIT",
		key,
		offset,
		value,
	)

	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cacheClusterAdapter) RGetBit(ctx context.Context, key string, offset int64) (int64, error) {
	cmd := redis.NewIntCmd(ctx,
		"R.GETBIT",
		key,
		offset,
	)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cacheClusterAdapter) RBitCount(ctx context.Context, key string) (int64, error) {
	cmd := redis.NewIntCmd(ctx,
		"R.BITCOUNT",
		key,
	)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()

}

func (r *cacheClusterAdapter) RSetIntArray(ctx context.Context, key string, value []int64) (bool, error) {
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

func (r *cacheClusterAdapter) RAppendIntArray(ctx context.Context, key string, value []int64) (bool, error) {
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

func (r *cacheClusterAdapter) RGetIntArray(ctx context.Context, key string) ([]int64, error) {
	cmd := redis.NewIntSliceCmd(ctx, "R.GETINTARRAY", key)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}

func (r *cacheClusterAdapter) ROptimize(ctx context.Context, key string) (bool, error) {
	cmd := redis.NewBoolCmd(ctx, "R.OPTIMIZE", key)
	_ = r.client.Process(ctx, cmd)
	return cmd.Result()
}
