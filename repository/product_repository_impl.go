package repository

import (
	"context"
	// "fmt"
	"go-rest-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, collection *mongo.Collection, product model.Product) model.Product {
	res, err := collection.InsertOne(ctx, product)
	if err != nil {
		panic(err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		product.Id = oid.Hex()
	}

	// product.Id = res.InsertedID.Hex()

	return product
}
