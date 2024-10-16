package handlers

import (
	"fmt"
	"net/http"
	"storage-service/domain"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// UpdateProductsHandler обновляет список продуктов.
//
// @Summary Обновление продуктов
// @Description Обновляет информацию о продуктах. Пользователь должен быть авторизован.
// @Tags Продукты
// @Accept  json
// @Param products body []domain.Product true "Список продуктов для обновления"
// @Success 200 "Продукты успешно обновлены"
// @Failure 400 "Ошибка в запросе или при обновлении продуктов"
// @Failure 401 "Пользователь не авторизован"
// @Router /products [put]
// @Security ApiKeyAuth
func UpdateProductsHandler(storageService service.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[update-product]")

		if !storageCtx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var products []domain.Product
		if err := goathttp.ReadRequestJson(r, &products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить обновленный продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := storageService.UpdateProducts(storageCtx, products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось обновить продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
