package rest

import (
	"context"
	"github.com/DosyaKitarov/market-sniper/internal/app/product"
)

type productService interface {
	InsertProducts(ctx context.Context, asins []string, country, tld string) ([]product.Product, error)
	GetAllProducts(ctx context.Context) ([]product.Product, error)
	ProductToCsv(ctx context.Context, products []product.Product) ([][]string, error)
}

type Services struct {
	productService productService
}

func NewHandlerService(productService productService) *Services {
	return &Services{
		productService: productService,
	}
}
