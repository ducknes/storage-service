package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id          int             `json:"id"`           // Id продукта
	BrandName   string          `json:"brand_name"`   // Бренд кроссовок
	FactoryName string          `json:"factory_name"` // Завод изготовитель
	Name        string          `json:"name"`         // Название модели кроссовка
	Description string          `json:"description"`  // Описание модели кроссовка
	Price       decimal.Decimal `json:"price"`        // Цена продукта
	Items       []ProductItem   `json:"items"`        // Варианты кроссовок
	Materials   []string        `json:"materials"`    // Материалы изготовления
	Images      []string        `json:"images"`       // Картинки
}

type ProductItem struct {
	StockCount int             `json:"stock_count"` // Кол-во на складе
	Size       int             `json:"size"`        // Размер
	Weight     decimal.Decimal `json:"weight"`      // Вес
	Color      string          `json:"color"`       // Цвет
}
