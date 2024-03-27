package repositories

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/jmoiron/sqlx"

    "D:\avdonin9pies-21\internal\app\models"
)

type ProductRepository struct {
    db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
    query := "INSERT INTO products (name, price) VALUES (:name, :price) RETURNING id"
    result := pr.db.NamedQueryContext(ctx, query, product)
    if err := result.Get(product.ID); err != nil {
        return fmt.Errorf("failed to create product: %w", err)
    }
    return nil
}

func (pr *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
    query := "SELECT id, name, price, created_at, updated_at FROM products WHERE id = $1"
    product := &models.Product{}
    if err := pr.db.GetContext(ctx, product, query, id); err != nil {
        return nil, fmt.Errorf("failed to get product: %w", err)
    }
    return product, nil
}

// Other methods for working with products
