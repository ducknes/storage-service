package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"storage-service/api/handlers"
	"storage-service/service"
	"storage-service/settings"
	"storage-service/tools/goathttp"

	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
)

func NewServer(ctx context.Context, logger goatlogger.Logger, cfg settings.Config, storageService service.Storage) *http.Server {
	router := mux.NewRouter()
	router.Use(goathttp.CommonJsonMiddleware, goathttp.CORSMiddleware, goathttp.PanicRecoveryMiddleware(logger))

	router.HandleFunc("/ping", handlers.PingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/health", handlers.HealthHandler()).Methods(http.MethodGet)

	addProductHandlers(router, storageService)
	addSwaggerHandler(router)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}

func addProductHandlers(router *mux.Router, storageService service.Storage) {
	router.HandleFunc("/product", handlers.GetProductHandler(storageService)).Methods(http.MethodGet)
	router.HandleFunc("/products", handlers.GetProductsHandler(storageService)).Methods(http.MethodGet)
	router.HandleFunc("/products", handlers.AddProductsHandler(storageService)).Methods(http.MethodPost)
	router.HandleFunc("/products", handlers.UpdateProductsHandler(storageService)).Methods(http.MethodPut)
	router.HandleFunc("/products", handlers.DeleteProductsHandler(storageService)).Methods(http.MethodDelete)
}

func addSwaggerHandler(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(handlers.SwaggerHandler())
}
