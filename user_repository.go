package repositories

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/jmoiron/sqlx"

    "D:\avdonin9pies-21\internal\app\models"
)

type UserRepository struct {
    db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
    query := "INSERT INTO users (name, email) VALUES (:name, :email) RETURNING id"
    result := ur.db.NamedQueryContext(ctx, query, user)
    if err := result.Get(user.ID); err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
    query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
    user := &models.User{}
    if err := ur.db.GetContext(ctx, user, query, id); err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}

// Other methods for working with users
