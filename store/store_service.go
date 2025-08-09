package store

import(
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type storageService struct{
	redisClient *redis.Client
}

var (
	storeService =&StrorageSerivce{}
	ctx=context.Background()
)

const CacheDuration = 6 * time.Hour