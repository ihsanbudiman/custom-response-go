package main

import (
	"kocag/handler"
	"kocag/usecase"
	httpresponse "kocag/utils/http-response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	httpresponse := httpresponse.NewResponseWriter()

	usecase := usecase.NewProductUsecase()

	handler := handler.NewHandler(httpresponse, usecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", handler.InsertProduct)
	http.ListenAndServe(":3000", r)
}
