package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"storage-service/service"
	"storage-service/tools/customerror"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// GetProductHandler возвращает информацию о продукте по его идентификатору.
//
// @Summary Получение информации о продукте
// @Description Возвращает данные о продукте по его идентификатору. Пользователь должен быть авторизован.
// @Tags Продукты
// @Produce  json
// @Param product_id query string true "Идентификатор продукта"
// @Success 200 {object} domain.Product "Информация о продукте успешно получена"
// @Success 204 "Информация о продукте отсутствует"
// @Failure 400 "Ошибка запроса или получения информации о продукте"
// @Router /product [get]
func GetProductHandler(storageService service.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[get-product]")

		productId := parseProductId(r)

		product, err := storageService.GetProduct(storageCtx, productId)
		if err != nil {
			if errors.Is(err, customerror.NoDocuments) {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			storageCtx.Log().Error(fmt.Sprintf("не удалось получить продукт %s, ошибка: %v", productId, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = goathttp.WriteResponseJson(w, http.StatusOK, product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось сериализовать продукт %s, ошибка: %v", productId, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func parseProductId(r *http.Request) string {
	return r.URL.Query().Get("product_id")
}
