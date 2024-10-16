package mappings

import (
	"github.com/samber/lo"
	"storage-service/database"
	"storage-service/domain"
)

func ToDbProduct(info domain.Product) database.Product {
	return database.Product{
		Id:          info.Id,
		BrandName:   info.BrandName,
		FactoryName: info.FactoryName,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Items:       ToDbProductItems(info.Items),
		Materials:   info.Materials,
		Images:      info.Images,
	}
}

func ToDomainProduct(info database.Product) domain.Product {
	return domain.Product{
		Id:          info.Id,
		BrandName:   info.BrandName,
		FactoryName: info.FactoryName,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Items:       ToDomainProductItems(info.Items),
		Materials:   info.Materials,
		Images:      info.Images,
	}
}

func ToDbProducts(products []domain.Product) []database.Product {
	return lo.Map(products, func(item domain.Product, _ int) database.Product {
		return ToDbProduct(item)
	})
}

func ToDomainProducts(products []database.Product) []domain.Product {
	return lo.Map(products, func(item database.Product, _ int) domain.Product {
		return ToDomainProduct(item)
	})
}

func ToDbProductItem(item domain.ProductItem) database.ProductItem {
	return database.ProductItem{
		StockCount: item.StockCount,
		Size:       item.Size,
		Weight:     item.Weight,
		Color:      item.Color,
	}
}

func ToDomainProductItem(item database.ProductItem) domain.ProductItem {
	return domain.ProductItem{
		StockCount: item.StockCount,
		Size:       item.Size,
		Weight:     item.Weight,
		Color:      item.Color,
	}
}

func ToDbProductItems(items []domain.ProductItem) []database.ProductItem {
	return lo.Map(items, func(item domain.ProductItem, _ int) database.ProductItem {
		return ToDbProductItem(item)
	})
}

func ToDomainProductItems(items []database.ProductItem) []domain.ProductItem {
	return lo.Map(items, func(item database.ProductItem, _ int) domain.ProductItem {
		return ToDomainProductItem(item)
	})
}

func ToAny[T any](items []T) (result []any) {
	for _, item := range items {
		result = append(result, item)
	}
	return
}
