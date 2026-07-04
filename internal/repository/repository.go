package repository

import "pos-backend/internal/model"

// BaseRepository defines generic CRUD operations
type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
	Update(entity *T) error
	Delete(id uint) error
	List() ([]T, error)
}

type UserRepository interface {
	BaseRepository[model.User]
	GetByUsername(username string) (*model.User, error)
}

type CategoryRepository interface {
	BaseRepository[model.Category]
}

type ProductRepository interface {
	BaseRepository[model.Product]
	GetByCategoryID(categoryID uint) ([]model.Product, error)
}

type OrderRepository interface {
	BaseRepository[model.Order]
	CreateOrder(order *model.Order, items []model.OrderItem) error
}
