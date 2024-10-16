package handlers

import (
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
	_ "storage-service/docs"
)

func SwaggerHandler() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"))
}
