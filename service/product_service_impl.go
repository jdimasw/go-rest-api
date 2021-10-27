package service

import (
	"context"
	"go-rest-api/model"
	"go-rest-api/model/web"
	"go-rest-api/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Client            *mongo.Client
}

func NewProductService(productRepository repository.ProductRepository, client *mongo.Client) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Client:            client,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	collection := service.Client.Database(os.Getenv("DB_NAME")).Collection("products")

	product := model.Product{
		Name:   request.Name,
		Detail: request.Detail,
	}

	product = service.ProductRepository.Save(ctx, collection, product)

	return web.ProductResponse{
		Id:     product.Id,
		Name:   product.Name,
		Detail: product.Detail,
	}
}