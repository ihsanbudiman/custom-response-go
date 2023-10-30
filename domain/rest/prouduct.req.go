package rest_domain

import (
	"encoding/json"
	error_utils "kocag/utils/error"
	"net/http"
)

type InsertProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (i *InsertProductRequest) Validate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return &error_utils.CustomError{
			Message: "Error",
			Status:  http.StatusInternalServerError,
		}
	}

	if i.Name == "" {
		return &error_utils.CustomError{
			Message: "Name is empty",
			Status:  http.StatusBadRequest,
		}
	}

	if i.Price == 0 {
		return &error_utils.CustomError{
			Message: "Price is empty",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
