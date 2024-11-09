package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"storage-service/database/pbmodels/pb"
	"storage-service/tools/storagecontext"
	"time"
)

const (
	_productsKey = "products-info"
	_defaultTTL  = 10 * time.Minute
)

type CacheRepository interface {
	CheckData(ctx storagecontext.StorageContext, limitKey string) bool
	Get(ctx storagecontext.StorageContext, limitKey string) ([]Product, error)
	Set(ctx storagecontext.StorageContext, items []Product, limitKey string) error
	Clear(ctx storagecontext.StorageContext)
}

type CacheRepositoryImpl struct {
	redisClient *redis.Client
}

func NewCacheRepository(redisClient *redis.Client) CacheRepository {
	return &CacheRepositoryImpl{
		redisClient: redisClient,
	}
}

func (c *CacheRepositoryImpl) CheckData(ctx storagecontext.StorageContext, limitKey string) bool {
	value := c.redisClient.Exists(ctx.Ctx(), createItemKey(limitKey))
	exists, err := value.Result()
	if err != nil {
		ctx.Log().Error(err.Error())
		return false
	}

	return exists == 1
}

func (c *CacheRepositoryImpl) Get(ctx storagecontext.StorageContext, limitKey string) ([]Product, error) {
	dataBytes, err := c.redisClient.Get(ctx.Ctx(), createItemKey(limitKey)).Bytes()
	if err != nil {
		return nil, err
	}

	products := &pb.MapProducts{}
	if err = proto.Unmarshal(dataBytes, products); err != nil {
		return nil, err
	}

	return ToDatabase(products), nil
}

func (c *CacheRepositoryImpl) Set(ctx storagecontext.StorageContext, items []Product, limitKey string) error {
	value, err := proto.Marshal(ToProto(items))
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx.Ctx(), createItemKey(limitKey), value, _defaultTTL).Err()
}

func (c *CacheRepositoryImpl) Clear(ctx storagecontext.StorageContext) {
	var currentCursor uint64
	for {
		keys, cursor, err := c.redisClient.Scan(ctx.Ctx(), currentCursor, "*", 100).Result()
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("redis scan error: %s", err.Error()))
			continue
		}

		currentCursor = cursor

		if len(keys) > 0 {
			if delErr := c.redisClient.Del(ctx.Ctx(), keys...).Err(); delErr != nil {
				ctx.Log().Error(fmt.Sprintf("redis del error: %s", delErr.Error()))
			}
		}

		if currentCursor == 0 {
			break
		}
	}
}

func createItemKey(limitKey string) string {
	return fmt.Sprintf("%s:%s", limitKey, _productsKey)
}
