package repositories

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/jmoiron/sqlx"

    "D:\avdonin9pies-21\internal\app\models"
)

type OrderRepository struct {
    db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
    query := "INSERT INTO orders (user_id, total_price) VALUES (:userId, :totalPrice) RETURNING id"
    result := or.db.NamedQueryContext(ctx, query, order)
    if err := result.Get(order.ID); err != nil {
        return fmt.Errorf("failed to create order: %w", err)
    }
    return nil
}

func (or *OrderRepository) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
    query := "SELECT id, user_id, total_price, created_at, updated_at FROM orders WHERE id = $1"
    order := &models.Order{}
    if err := or.db.GetContext(ctx, order, query, id); err != nil {
        return nil, fmt.Errorf("failed to get order: %w", err)
    }
    return order, nil
}

// Other methods for working with orders
