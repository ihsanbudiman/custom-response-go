package rest_domain

type InsertProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
