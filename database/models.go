package database

import (
	"github.com/samber/lo"
	"storage-service/database/pbmodels/pb"
	"time"
)

type Product struct {
	Id          string        `bson:"_id" json:"Id"`
	BrandName   string        `bson:"brand_name" json:"BrandName"`
	FactoryName string        `bson:"factory_name" json:"FactoryName"`
	Name        string        `bson:"name" json:"Name"`
	Description string        `bson:"description" json:"Description"`
	Price       float64       `bson:"price" json:"Price"`
	Items       []ProductItem `bson:"items" json:"Items"`
	Materials   []string      `bson:"materials" json:"Materials"`
	Images      []string      `bson:"images" json:"Images"`
	Approver    string        `bson:"approver" json:"Approver"`
	Status      string        `bson:"status" json:"Status"`
	ApproveTime time.Time     `bson:"approve_time"`
}

type ProductForInsert struct {
	BrandName   string        `bson:"brand_name"`
	FactoryName string        `bson:"factory_name"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Price       float64       `bson:"price"`
	Items       []ProductItem `bson:"items"`
	Materials   []string      `bson:"materials"`
	Images      []string      `bson:"images"`
	Approver    string        `bson:"approver"`
	Status      string        `bson:"status"`
	ApproveTime time.Time     `bson:"approve_time"`
}

type ProductItem struct {
	StockCount int     `bson:"stock_count" json:"StockCount"`
	Size       int     `bson:"size" json:"Size"`
	Weight     float64 `bson:"weight" json:"Weight"`
	Color      string  `bson:"color" json:"Color"`
}

type ApproveMessage struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type ApprovedItem struct {
	ProductId   string    `json:"product_id"`
	ApproveTime time.Time `json:"approve_time"`
}

func ToInsertItem(product Product) ProductForInsert {
	return ProductForInsert{
		BrandName:   product.BrandName,
		FactoryName: product.FactoryName,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Items:       product.Items,
		Materials:   product.Materials,
		Images:      product.Images,
		Approver:    product.Approver,
		Status:      product.Status,
		ApproveTime: product.ApproveTime,
	}
}

func ToInsertItems(products []Product) []ProductForInsert {
	return lo.Map(products, func(item Product, _ int) ProductForInsert {
		return ToInsertItem(item)
	})
}

func ToDatabase(products *pb.MapProducts) []Product {
	return lo.MapToSlice(products.Items, func(_ string, value *pb.Product) Product {
		return ToDatabaseProduct(value)
	})
}

func ToDatabaseProduct(product *pb.Product) Product {
	return Product{
		Id:          product.Id,
		BrandName:   product.BrandName,
		FactoryName: product.FactoryName,
		Name:        product.Name,
		Description: product.Description,
		Price:       float64(product.Price),
		Items:       ToDatabaseProductItems(product.Items),
		Materials:   product.Materials,
		Images:      product.Images,
	}
}

func ToDatabaseProductItem(item *pb.ProductItem) ProductItem {
	return ProductItem{
		StockCount: int(item.StockCount),
		Size:       int(item.Size),
		Weight:     float64(item.Weight),
		Color:      item.Color,
	}
}

func ToDatabaseProductItems(items []*pb.ProductItem) []ProductItem {
	return lo.Map(items, func(item *pb.ProductItem, _ int) ProductItem {
		return ToDatabaseProductItem(item)
	})
}

func ToProto(items []Product) *pb.MapProducts {
	result := &pb.MapProducts{Items: make(map[string]*pb.Product)}
	for _, item := range items {
		result.Items[item.Id] = ToProtoProduct(item)
	}
	return result
}

func ToProtoProduct(product Product) *pb.Product {
	return &pb.Product{
		Id:          product.Id,
		BrandName:   product.BrandName,
		FactoryName: product.FactoryName,
		Name:        product.Name,
		Description: product.Description,
		Price:       float32(product.Price),
		Items:       ToProtoProductItems(product.Items),
		Materials:   product.Materials,
		Images:      product.Images,
	}
}

func ToProtoProductItem(item ProductItem) *pb.ProductItem {
	return &pb.ProductItem{
		StockCount: int32(item.StockCount),
		Size:       int32(item.Size),
		Weight:     float32(item.Weight),
		Color:      item.Color,
	}
}

func ToProtoProductItems(items []ProductItem) []*pb.ProductItem {
	return lo.Map(items, func(item ProductItem, _ int) *pb.ProductItem {
		return ToProtoProductItem(item)
	})
}
