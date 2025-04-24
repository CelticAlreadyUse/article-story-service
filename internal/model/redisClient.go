package model

import (
	"context"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key string, data any, exp time.Duration) error
	Get(ctx context.Context, key string, data any) error
	Del(ctx context.Context, keys ...string) error
	HSet(ctx context.Context, bucketKey, key string, data any, exp time.Duration) error
	HGet(ctx context.Context, bucketKey, key string, data any) error
	HDelByBucketKey(ctx context.Context, bucketKey string) error
	HDelByBucketKeyAndKeys(ctx context.Context, bucketKey string, keys ...string) error
}