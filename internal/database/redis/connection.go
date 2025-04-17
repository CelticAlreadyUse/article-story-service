package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)
type redisClient struct {
	redisClient *redis.Client
}
func InitConnectRedis() model.RedisClient {
	addr := fmt.Sprintf(config.RedisHost() + ":" + config.RedisPort())
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisPass(),
		DB:       0,
		Protocol: 2,
	}) 
	
	logrus.Info("redis connected")
	return &redisClient{redisClient: client}
}
func (r *redisClient) Set(ctx context.Context, key string, data any, exp time.Duration) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("failed convert data to byte in redis set, error: %v", err)
		return err
	}
	err = r.redisClient.Set(ctx, key, byteData, exp).Err()
	if err != nil {
		logrus.Errorf("failed insert data to redis, error: %v", err)
	}
	return err
}
func (r *redisClient) Get(ctx context.Context, key string, data any) error {
	value, err := r.redisClient.Get(ctx, key).Result()
	switch err {
	case nil:
	case redis.Nil:
		return nil
	default:
		logrus.Errorf("failed get data from redis, error: %v", err)
		return err
	}
	return json.Unmarshal([]byte(value), &data)
}
func (r *redisClient) Del(ctx context.Context, keys ...string) error {
	err := r.redisClient.Del(ctx, keys...).Err()
	if err != nil {
		logrus.Errorf("failed delete data from redis, error: %v", err)
	}
	return err
}
func (r *redisClient) HSet(ctx context.Context, bucketKey, key string, data any, exp time.Duration) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("failed convert data to byte in redis set, error: %v", err)
		return err
	}
	err = r.redisClient.HSet(ctx, bucketKey, key, byteData).Err()
	if err != nil {
		logrus.Errorf("failed insert data to redis, error: %v", err)
	}
	return err
}
func (r *redisClient) HGet(ctx context.Context, bucketKey, key string, data any) error {
	value, err := r.redisClient.HGet(ctx, bucketKey, key).Result()
	switch err {
	case nil:
	case redis.Nil:
		return nil
	default:
		logrus.Errorf("failed get data from redis, error: %v", err)
		return err
	}
	return json.Unmarshal([]byte(value), &data)
}

func (r *redisClient) HDelByBucketKey(ctx context.Context, bucket string) error {
	fields, err := r.redisClient.HKeys(ctx, bucket).Result()
	if err != nil {
		return err
	}
	if len(fields) == 0 {
		return nil
	}
	return r.redisClient.HDel(ctx, bucket, fields...).Err()
}

func (r *redisClient) HDelByBucketKeyAndKeys(ctx context.Context, bucketKey string, keys ...string) error {
	err := r.redisClient.HDel(ctx, bucketKey, keys...).Err()
	if err != nil {
		logrus.Errorf("failed delete data from redis, error: %v", err)
	}
	return err
}
