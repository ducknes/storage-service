package domain

import "github.com/shopspring/decimal"

type Products struct {
	Items      []Product `json:"items"`      // Список продуктов
	Limit      int64     `json:"limit"`      // Кол-во элементов
	Cursor     string    `json:"cursor"`     // Текущий курсор
	NextCursor string    `json:"nextCursor"` // Курсор для запроса след страницы
}

type Product struct {
	Id          string          `json:"id"`          // Id продукта
	BrandName   string          `json:"brandName"`   // Бренд кроссовок
	FactoryName string          `json:"factoryName"` // Завод изготовитель
	Name        string          `json:"name"`        // Название модели кроссовка
	Description string          `json:"description"` // Описание модели кроссовка
	Price       decimal.Decimal `json:"price"`       // Цена продукта
	Items       []ProductItem   `json:"items"`       // Варианты кроссовок
	Materials   []string        `json:"materials"`   // Материалы изготовления
	Images      []string        `json:"images"`      // Картинки
}

type ProductItem struct {
	StockCount int             `json:"stockCount"` // Кол-во на складе
	Size       int             `json:"size"`       // Размер
	Weight     decimal.Decimal `json:"weight"`     // Вес
	Color      string          `json:"color"`      // Цвет
}
