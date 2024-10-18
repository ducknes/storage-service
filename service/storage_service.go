package service

import (
	"storage-service/database"
	"storage-service/domain"
	"storage-service/domain/mappings"
	"storage-service/tools/storagecontext"
)

type Storage interface {
	GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) (domain.Products, error)
	GetProduct(ctx storagecontext.StorageContext, productId string) (domain.Product, error)
	SaveProducts(ctx storagecontext.StorageContext, products []domain.Product) error
	UpdateProducts(ctx storagecontext.StorageContext, products []domain.Product) error
	RemoveProducts(ctx storagecontext.StorageContext, productIds []string) error
}

type StorageServiceImpl struct {
	storageRepository database.StorageRepository
}

func NewStorageService(repo database.StorageRepository) Storage {
	return &StorageServiceImpl{
		storageRepository: repo,
	}
}

func (s *StorageServiceImpl) GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) (domain.Products, error) {
	dbProducts, err := s.storageRepository.GetProducts(ctx, limit, cursor)
	if err != nil {
		return domain.Products{}, err
	}

	nextCursor := ""
	if len(dbProducts) > 0 {
		nextCursor = dbProducts[len(dbProducts)-1].Id
	}

	return domain.Products{
		Items:      mappings.ToDomainProducts(dbProducts),
		Limit:      limit,
		Cursor:     cursor,
		NextCursor: nextCursor,
	}, err
}

func (s *StorageServiceImpl) GetProduct(ctx storagecontext.StorageContext, productId string) (domain.Product, error) {
	dbProduct, err := s.storageRepository.GetProduct(ctx, productId)
	return mappings.ToDomainProduct(dbProduct), err
}

func (s *StorageServiceImpl) SaveProducts(ctx storagecontext.StorageContext, products []domain.Product) error {
	return s.storageRepository.AddProducts(ctx, mappings.ToDbProducts(products))
}

func (s *StorageServiceImpl) UpdateProducts(ctx storagecontext.StorageContext, products []domain.Product) error {
	return s.storageRepository.UpdateProducts(ctx, mappings.ToDbProducts(products))
}

func (s *StorageServiceImpl) RemoveProducts(ctx storagecontext.StorageContext, productIds []string) error {
	return s.storageRepository.DeleteProducts(ctx, productIds)
}
