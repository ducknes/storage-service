package database

import (
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"storage-service/database/pbmodels/pb"
	"storage-service/tools/storagecontext"
	"time"
)

const (
	_productsKey = "products-info"
	_defaultTTL  = 60 * time.Second
)

type CacheRepository interface {
	CheckData(ctx storagecontext.StorageContext) bool
	Get(ctx storagecontext.StorageContext) ([]Product, error)
	Set(ctx storagecontext.StorageContext, items []Product) error
}

type CacheRepositoryImpl struct {
	redisClient *redis.Client
}

func NewCacheRepository(redisClient *redis.Client) CacheRepository {
	return &CacheRepositoryImpl{
		redisClient: redisClient,
	}
}

func (c *CacheRepositoryImpl) CheckData(ctx storagecontext.StorageContext) bool {
	value := c.redisClient.Exists(ctx.Ctx(), _productsKey)
	exists, err := value.Result()
	if err != nil {
		ctx.Log().Error(err.Error())
		return false
	}

	return exists == 1
}

func (c *CacheRepositoryImpl) Get(ctx storagecontext.StorageContext) ([]Product, error) {
	dataBytes, err := c.redisClient.Get(ctx.Ctx(), _productsKey).Bytes()
	if err != nil {
		return nil, err
	}

	products := &pb.MapProducts{}
	if err = proto.Unmarshal(dataBytes, products); err != nil {
		return nil, err
	}

	return ToDatabase(products), nil
}

func (c *CacheRepositoryImpl) Set(ctx storagecontext.StorageContext, items []Product) error {
	value, err := proto.Marshal(ToProto(items))
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx.Ctx(), _productsKey, value, _defaultTTL).Err()
}
