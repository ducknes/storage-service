package handlers

import (
	"fmt"
	"net/http"
	"storage-service/domain"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// AddProductsHandler добавляет список продуктов в хранилище.
//
// @Summary Добавление продуктов
// @Description Добавляет новые продукты в систему. Пользователь должен быть авторизован.
// @Tags Продукты
// @Accept  json
// @Param products body []domain.AddingProduct true "Список продуктов для добавления"
// @Success 200 "Продукты успешно добавлены"
// @Failure 400 "Ошибка в запросе или при добавлении продуктов"
// @Router /products [post]
func AddProductsHandler(storageService service.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[add-product]")

		var products []domain.AddingProduct
		if err := goathttp.ReadRequestJson(r, &products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить новый продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := storageService.SaveProducts(storageCtx, products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось добавить новый продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
