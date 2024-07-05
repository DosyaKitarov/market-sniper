package product

import (
	"context"
	"github.com/DosyaKitarov/market-sniper/api/scrapper"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/logger"
)

type repository interface {
	insertProducts(ctx context.Context, p []Product) ([]Product, error)
	getAllProducts() ([]Product, error)
	//getProducts(ctx context.Context, products []Product) ([]Product, error)
	//updateProduct(p Product) (Product, error)
}
type productService struct {
	repository repository
}

func NewProductService(repository repository) *productService {
	return &productService{
		repository: repository,
	}
}

func (s *productService) InsertProducts(ctx context.Context, asins []string, country, tld string) ([]Product, error) {

	var productArray []Product
	// Fetch product data
	for _, asin := range asins {
		var p Product

		json, err := scrapper.FetchProductData(ctx, asin, country, tld)
		if err != nil {
			logger.Error(ctx, "InsertProducts.FetchProductData() error: %v", err)
			return productArray, err
		}

		p.FormatProduct(asin, json)
		productArray = append(productArray, p)
	}

	s.repository.insertProducts(ctx, productArray)

	return productArray, nil
}

func (s *productService) GetAllProducts(ctx context.Context) ([]Product, error) {
	return s.repository.getAllProducts()
}

func (s *productService) ProductToCsv(ctx context.Context, products []Product) ([][]string, error) {
	// Initialize the CSV data slice with a header row
	csvData := [][]string{
		{"ASIN", "Name", "Brand", "Price", "PreviousPrice", "ChangeDate"}, // Adjust these headers based on the actual Product struct fields
	}

	// Iterate over the products to convert each to a CSV row
	for _, p := range products {
		row := []string{
			p.ASIN,
			p.Name,
			p.Brand,
			p.Price,
			p.PreviousPrice,
			p.ChangeDate,
		}
		csvData = append(csvData, row)
	}

	return csvData, nil
}
