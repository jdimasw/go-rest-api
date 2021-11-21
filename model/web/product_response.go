package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductResponse struct {
	// Id     string
	Name   string
	Detail string
}

type ProductDetailResponse struct {
	Id     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	Detail string             `json:"detail"`
}
