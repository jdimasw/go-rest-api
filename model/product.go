package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Name   string
	Detail string
}

type ProductDetail struct {
	Id     primitive.ObjectID `bson:"_id",omitempty`
	Name   string
	Detail string
}
