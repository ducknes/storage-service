package database

import "github.com/shopspring/decimal"

type Product struct {
	Id          int             `bson:"_id" json:"Idd"`
	BrandName   string          `bson:"brand_name" json:"BrandName"`
	FactoryName string          `bson:"factory_name" json:"FactoryName"`
	Name        string          `bson:"name" json:"Name"`
	Description string          `bson:"description" json:"Description"`
	Price       decimal.Decimal `bson:"price" json:"Price"`
	Items       []ProductItem   `bson:"items" json:"Items"`
	Materials   []string        `bson:"materials" json:"Materials"`
	Images      []string        `bson:"images" json:"Images"`
}

type ProductItem struct {
	StockCount int             `bson:"stock_count" json:"StockCount"`
	Size       int             `bson:"size" json:"Size"`
	Weight     decimal.Decimal `bson:"weight" json:"Weight"`
	Color      string          `bson:"color" json:"Color"`
}
