package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id          int             `bson:"id"`          // Id продукта
	BrandName   string          `bson:"brand_id"`    // Бренд кроссовок
	FactoryName string          `bson:"factory_id"`  // Завод изготовитель
	Name        string          `bson:"name"`        // Название модели кроссовка
	Description string          `bson:"description"` // Описание модели кроссовка
	Price       decimal.Decimal `bson:"price"`       // Цена продукта
	Items       []ProductItem   `json:"items"`       // Варианты кроссовок
	Materials   []string        `json:"materials"`   // Материалы изготовления
	Images      []string        `json:"images"`      // Картинки
}

type ProductItem struct {
	StockCount int             `json:"stock_count"` // Кол-во на складе
	Size       int             `json:"size"`        // Размер
	Weight     decimal.Decimal `json:"weight"`      // Вес
	Color      string          `json:"color"`       // Цвет
}
