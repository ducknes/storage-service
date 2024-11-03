package service

import (
	"fmt"
	"storage-service/database"
	"storage-service/domain"
	"storage-service/domain/mappings"
	"storage-service/tools/storagecontext"
)

type Storage interface {
	GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) (domain.Products, error)
	GetProduct(ctx storagecontext.StorageContext, productId string) (domain.Product, error)
	SaveProducts(ctx storagecontext.StorageContext, products []domain.AddingProduct) error
	UpdateProducts(ctx storagecontext.StorageContext, products []domain.Product) error
	RemoveProducts(ctx storagecontext.StorageContext, productIds []string) error
}

type StorageServiceImpl struct {
	storageRepository database.StorageRepository
	cacheRepository   database.CacheRepository
}

func NewStorageService(repo database.StorageRepository, cache database.CacheRepository) Storage {
	return &StorageServiceImpl{
		storageRepository: repo,
		cacheRepository:   cache,
	}
}

func (s *StorageServiceImpl) GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) (domain.Products, error) {
	if s.cacheRepository.CheckData(ctx, cursor) {
		cacheProducts, err := s.cacheRepository.Get(ctx, cursor)
		if err != nil {
			return domain.Products{}, err
		}

		return domain.Products{
			Items:      mappings.ToDomainProducts(cacheProducts),
			Limit:      limit,
			Cursor:     cursor,
			NextCursor: getNextCursor(cacheProducts),
			FromCache:  true,
		}, err
	}

	dbProducts, err := s.storageRepository.GetProducts(ctx, limit, cursor)
	if err != nil {
		return domain.Products{}, err
	}

	defer func() {
		err = s.cacheRepository.Set(ctx, dbProducts, cursor)
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось сохранить данные в кэш: %v", err))
		}
	}()

	return domain.Products{
		Items:      mappings.ToDomainProducts(dbProducts),
		Limit:      limit,
		Cursor:     cursor,
		NextCursor: getNextCursor(dbProducts),
	}, err
}

func (s *StorageServiceImpl) GetProduct(ctx storagecontext.StorageContext, productId string) (domain.Product, error) {
	dbProduct, err := s.storageRepository.GetProduct(ctx, productId)
	return mappings.ToDomainProduct(dbProduct), err
}

func (s *StorageServiceImpl) SaveProducts(ctx storagecontext.StorageContext, products []domain.AddingProduct) error {
	return s.storageRepository.AddProducts(ctx, mappings.ToDbAddingProducts(products))
}

func (s *StorageServiceImpl) UpdateProducts(ctx storagecontext.StorageContext, products []domain.Product) error {
	return s.storageRepository.UpdateProducts(ctx, mappings.ToDbProducts(products))
}

func (s *StorageServiceImpl) RemoveProducts(ctx storagecontext.StorageContext, productIds []string) error {
	return s.storageRepository.DeleteProducts(ctx, productIds)
}

func getNextCursor(items []database.Product) string {
	nextCursor := ""
	if len(items) > 0 {
		nextCursor = items[len(items)-1].Id
	}

	return nextCursor
}
