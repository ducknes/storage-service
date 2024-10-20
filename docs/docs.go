// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "Возвращает простой ответ \"pong\" для проверки доступности сервера.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Системные"
                ],
                "summary": "Проверка состояния сервера",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Возвращает данные о продукте по его идентификатору. Пользователь должен быть авторизован.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Получение информации о продукте",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор продукта",
                        "name": "product_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о продукте успешно получена",
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    },
                    "204": {
                        "description": "Информация о продукте отсутствует"
                    },
                    "400": {
                        "description": "Ошибка запроса или получения информации о продукте"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Возвращает список продуктов с возможностью пагинации. Пользователь должен быть авторизован.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Получение списка продуктов",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Количество продуктов для выборки",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Курсор для пагинации",
                        "name": "cursor",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список продуктов успешно получен",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Products"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка запроса или получения списка продуктов"
                    }
                }
            },
            "put": {
                "description": "Обновляет информацию о продуктах. Пользователь должен быть авторизован.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Обновление продуктов",
                "parameters": [
                    {
                        "description": "Список продуктов для обновления",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Product"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Продукты успешно обновлены"
                    },
                    "400": {
                        "description": "Ошибка в запросе или при обновлении продуктов"
                    }
                }
            },
            "post": {
                "description": "Добавляет новые продукты в систему. Пользователь должен быть авторизован.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Добавление продуктов",
                "parameters": [
                    {
                        "description": "Список продуктов для добавления",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.AddingProduct"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Продукты успешно добавлены"
                    },
                    "400": {
                        "description": "Ошибка в запросе или при добавлении продуктов"
                    }
                }
            },
            "delete": {
                "description": "Удаляет продукты из системы по списку идентификаторов. Пользователь должен быть авторизован.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Удаление продуктов",
                "parameters": [
                    {
                        "description": "Список идентификаторов продуктов для удаления",
                        "name": "productIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Продукты успешно удалены"
                    },
                    "400": {
                        "description": "Ошибка в запросе или при удалении продуктов"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AddingProduct": {
            "type": "object",
            "properties": {
                "brandName": {
                    "description": "Бренд кроссовок",
                    "type": "string"
                },
                "description": {
                    "description": "Описание модели кроссовка",
                    "type": "string"
                },
                "factoryName": {
                    "description": "Завод изготовитель",
                    "type": "string"
                },
                "images": {
                    "description": "Картинки",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "items": {
                    "description": "Варианты кроссовок",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ProductItem"
                    }
                },
                "materials": {
                    "description": "Материалы изготовления",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Название модели кроссовка",
                    "type": "string"
                },
                "price": {
                    "description": "Цена продукта",
                    "type": "number"
                }
            }
        },
        "domain.Product": {
            "type": "object",
            "properties": {
                "brandName": {
                    "description": "Бренд кроссовок",
                    "type": "string"
                },
                "description": {
                    "description": "Описание модели кроссовка",
                    "type": "string"
                },
                "factoryName": {
                    "description": "Завод изготовитель",
                    "type": "string"
                },
                "id": {
                    "description": "Id продукта",
                    "type": "string"
                },
                "images": {
                    "description": "Картинки",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "items": {
                    "description": "Варианты кроссовок",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ProductItem"
                    }
                },
                "materials": {
                    "description": "Материалы изготовления",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Название модели кроссовка",
                    "type": "string"
                },
                "price": {
                    "description": "Цена продукта",
                    "type": "number"
                }
            }
        },
        "domain.ProductItem": {
            "type": "object",
            "properties": {
                "color": {
                    "description": "Цвет",
                    "type": "string"
                },
                "size": {
                    "description": "Размер",
                    "type": "integer"
                },
                "stockCount": {
                    "description": "Кол-во на складе",
                    "type": "integer"
                },
                "weight": {
                    "description": "Вес",
                    "type": "number"
                }
            }
        },
        "domain.Products": {
            "type": "object",
            "properties": {
                "cursor": {
                    "description": "Текущий курсор",
                    "type": "string"
                },
                "items": {
                    "description": "Список продуктов",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Product"
                    }
                },
                "limit": {
                    "description": "Кол-во элементов",
                    "type": "integer"
                },
                "nextCursor": {
                    "description": "Курсор для запроса след страницы",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "storage-service",
	Description:      "Сервис управления продукцией на складе для goat-logistics",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
