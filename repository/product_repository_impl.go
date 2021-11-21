package repository

import (
	"context"
	// "fmt"
	"go-rest-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, collection *mongo.Collection, product model.Product) model.Product {
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		panic(err)
	}

	return product

	// var productId string

	// if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
	// 	productId = oid.Hex()
	// }
	// objID, err := primitive.ObjectIDFromHex(res.InsertedID)
	// if err != nil {
	// 	panic(err)
	// }
	// productId := objID

	// return model.ProductDetail{
	// 	Id:     productId,
	// 	Name:   product.Name,
	// 	Detail: product.Detail,
	// }
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, collection *mongo.Collection, productId string) model.ProductDetail {
	var product model.ProductDetail
	objID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		panic(err)
	}
	err = collection.FindOne(ctx, bson.D{{"_id", objID}}).Decode(&product)
	if err != nil {
		panic(err)
	}

	return product
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, collection *mongo.Collection) []model.ProductDetail {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var products []model.ProductDetail

	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var result model.ProductDetail
		err := cur.Decode(&result)
		if err != nil {
			panic(err)
		}
		// do something with result...
		products = append(products, result)
	}

	return products
}
