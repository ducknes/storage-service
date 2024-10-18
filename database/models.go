package database

import "github.com/samber/lo"

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
}

type ProductItem struct {
	StockCount int     `bson:"stock_count" json:"StockCount"`
	Size       int     `bson:"size" json:"Size"`
	Weight     float64 `bson:"weight" json:"Weight"`
	Color      string  `bson:"color" json:"Color"`
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
	}
}

func ToInsertItems(products []Product) []ProductForInsert {
	return lo.Map(products, func(item Product, _ int) ProductForInsert {
		return ToInsertItem(item)
	})
}
