package product

import (
	"context"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) insertProducts(ctx context.Context, products []Product) ([]Product, error) {
	collection := r.db.Database("AmazonScrapper").Collection("Products")

	// Convert products slice to a slice of interface{} for insertion
	var productsInterface []interface{} = make([]interface{}, len(products))
	for i, v := range products {
		productsInterface[i] = v
	}

	// Insert products into the collection
	_, err := collection.InsertMany(ctx, productsInterface)
	if err != nil {
		logger.Error(ctx, "insertProducts.InsertMany() error: ", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) getAllProducts() ([]Product, error) {
	collection := r.db.Database("AmazonScrapper").Collection("Products")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		logger.Error(context.Background(), "getAllProducts.Find() error: ", err)
		return nil, err
	}

	var products []Product
	if err = cursor.All(context.Background(), &products); err != nil {
		logger.Error(context.Background(), "getAllProducts.All() error: ", err)
		return nil, err
	}

	return products, nil
}
