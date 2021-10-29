package repository

import (
	"context"
	"go-rest-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Save(ctx context.Context, collection *mongo.Collection, product model.Product) model.Product
	// 	Update(ctx context.Context)
	// 	Delete(ctx context.Context)
	// 	FindById(ctx context.Context)
	FindAll(ctx context.Context, collection *mongo.Collection) []model.ProductDetail
}
