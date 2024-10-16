package database

import "github.com/shopspring/decimal"

type Product struct {
	Id          int             `bson:"_id"`
	BrandName   string          `bson:"brand_id"`
	FactoryName string          `bson:"factory_id"`
	Name        string          `bson:"name"`
	Description string          `bson:"description"`
	Price       decimal.Decimal `bson:"price"`
	Items       []ProductItem   `bson:"items"`
	Materials   []string        `bson:"materials"`
	Images      []string        `bson:"images"`
}

type ProductItem struct {
	StockCount int             `bson:"stock_count"`
	Size       int             `bson:"size"`
	Weight     decimal.Decimal `bson:"weight"`
	Color      string          `bson:"color"`
}
