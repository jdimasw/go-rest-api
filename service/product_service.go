package service

import (
	"context"
	"go-rest-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	FindById(ctx context.Context, productId string) web.ProductDetailResponse
	FindAll(ctx context.Context) []web.ProductDetailResponse
}
