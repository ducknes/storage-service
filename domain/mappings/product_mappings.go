package mappings

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
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
		Price:       info.Price.InexactFloat64(),
		Items:       ToDbProductItems(info.Items),
		Materials:   info.Materials,
		Images:      info.Images,
		Approver:    info.Approver,
		Status:      string(info.Status),
	}
}

func ToDbAddingProduct(info domain.AddingProduct) database.Product {
	return database.Product{
		BrandName:   info.BrandName,
		FactoryName: info.FactoryName,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price.InexactFloat64(),
		Items:       ToDbProductItems(info.Items),
		Materials:   info.Materials,
		Images:      info.Images,
		Approver:    info.Approver,
		Status:      string(info.Status),
	}
}

func ToDomainProduct(info database.Product) domain.Product {
	return domain.Product{
		Id:          info.Id,
		BrandName:   info.BrandName,
		FactoryName: info.FactoryName,
		Name:        info.Name,
		Description: info.Description,
		Price:       decimal.NewFromFloat(info.Price),
		Items:       ToDomainProductItems(info.Items),
		Materials:   info.Materials,
		Images:      info.Images,
		Approver:    info.Approver,
		Status:      domain.ProductStatus(info.Status),
	}
}

func ToDbProducts(products []domain.Product) []database.Product {
	return lo.Map(products, func(item domain.Product, _ int) database.Product {
		return ToDbProduct(item)
	})
}

func ToDbAddingProducts(products []domain.AddingProduct) []database.Product {
	return lo.Map(products, func(item domain.AddingProduct, _ int) database.Product {
		return ToDbAddingProduct(item)
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
		Weight:     item.Weight.InexactFloat64(),
		Color:      item.Color,
	}
}

func ToDomainProductItem(item database.ProductItem) domain.ProductItem {
	return domain.ProductItem{
		StockCount: item.StockCount,
		Size:       item.Size,
		Weight:     decimal.NewFromFloat(item.Weight),
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
