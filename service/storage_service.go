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

	return domain.Products{
		Items:      mappings.ToDomainProducts(dbProducts),
		Limit:      limit,
		NextCursor: dbProducts[len(dbProducts)-1].Id,
	}, err
}

func (s *StorageServiceImpl) GetProduct(ctx storagecontext.StorageContext, productId string) (domain.Product, error) {
	dbProduct, err := s.storageRepository.GetProduct(ctx, productId)
	if err != nil {
		return domain.Product{}, err
	}

	return mappings.ToDomainProduct(dbProduct), nil
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
