package handlers

import (
	"fmt"
	"net/http"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
	"strconv"
)

// GetProductsHandler возвращает список продуктов.
//
// @Summary Получение списка продуктов
// @Description Возвращает список продуктов с возможностью пагинации. Пользователь должен быть авторизован.
// @Tags Продукты
// @Produce  json
// @Param limit query int true "Количество продуктов для выборки" default(10)
// @Param cursor query string false "Курсор для пагинации"
// @Success 200 {array} domain.Product "Список продуктов успешно получен"
// @Failure 400 "Ошибка запроса или получения списка продуктов"
// @Router /products [get]
func GetProductsHandler(storageService service.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[get-products]")

		limit, cursor := parseQuery(r)

		products, err := storageService.GetProducts(storageCtx, limit, cursor)
		if err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось получить список продуктов, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = goathttp.WriteResponseJson(w, http.StatusOK, products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось сериализовать список продуктов, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func parseQuery(r *http.Request) (int64, string) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	cursor := r.URL.Query().Get("cursor")

	return int64(limit), cursor
}
