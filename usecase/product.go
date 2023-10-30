package usecase

import (
	"context"
	error_utils "kocag/utils/error"
)

type ProductUsecase struct {
}

type ProductUsecaseInterface interface {
	InsertProduct(ctx context.Context, name string, price int) error
}

func (usecase *ProductUsecase) InsertProduct(ctx context.Context, name string, price int) error {

	if name == "" {
		return &error_utils.CustomError{
			Message: "Name is empty",
			Status:  400,
		}
	}

	if price == 0 {
		return &error_utils.CustomError{
			Message: "Price is empty",
			Status:  500,
		}
	}

	return nil
}

func NewProductUsecase() ProductUsecaseInterface {
	return &ProductUsecase{}
}
