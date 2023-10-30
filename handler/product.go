package handler

import (
	rest_domain "kocag/domain/rest"
	"kocag/usecase"
	"net/http"

	httpresponse "kocag/utils/http-response"
)

type Handler struct {
	response httpresponse.IResponseWriter
	usecase  usecase.ProductUsecaseInterface
}

func (handler *Handler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	req := rest_domain.InsertProductRequest{}
	err := req.Validate(r)

	if err != nil {
		handler.response.HttpError(w, err)
		return
	}

	err = handler.usecase.InsertProduct(r.Context(), req.Name, req.Price)
	if err != nil {
		handler.response.HttpError(w, err)
		return
	}

	handler.response.JSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func NewHandler(response httpresponse.IResponseWriter, usecase usecase.ProductUsecaseInterface) *Handler {
	return &Handler{
		response: response,
		usecase:  usecase,
	}
}
