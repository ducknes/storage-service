package handlers

import "net/http"

// PingHandler проверяет доступность сервера.
//
// @Summary Проверка состояния сервера
// @Description Возвращает простой ответ "pong" для проверки доступности сервера.
// @Tags Системные
// @Produce plain
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}
