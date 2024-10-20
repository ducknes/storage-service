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

// DeleteProductsHandler удаляет продукты по их идентификаторам.
//
// @Summary Удаление продуктов
// @Description Удаляет продукты из системы по списку идентификаторов. Пользователь должен быть авторизован.
// @Tags Продукты
// @Accept  json
// @Param productIds body []string true "Список идентификаторов продуктов для удаления"
// @Success 200 "Продукты успешно удалены"
// @Failure 400 "Ошибка в запросе или при удалении продуктов"
// @Router /products [delete]
func DeleteProductsHandler(storageService service.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[delete-product]")

		var productIds []string
		if err := goathttp.ReadRequestJson(r, &productIds); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить id продукта, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := storageService.RemoveProducts(storageCtx, productIds); err != nil {
			if errors.Is(err, customerror.NoDocumentAffected) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			storageCtx.Log().Error(fmt.Sprintf("не удалось удалить продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
