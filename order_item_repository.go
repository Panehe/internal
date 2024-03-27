package repositories

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/jmoiron/sqlx"

    "github.com/your-package/models"
)

type OrderItemRepository struct {
    db *sqlx.DB
}

func NewOrderItemRepository(db *sqlx.DB) *OrderItemRepository {
    return &OrderItemRepository{db: db}
}

func (oir *OrderItemRepository) CreateOrderItem(ctx context.Context, orderItem *models.OrderItem) error {
    query := "INSERT INTO order_items (order_id, product_id, quantity) VALUES (:orderId, :productId, :quantity)"
    _, err := oir.db.NamedExecContext(ctx, query, orderItem)
    if err != nil {
        return fmt.Errorf("failed to create order item: %w", err)
    }
    return nil
}


// Other methods for working with order items
