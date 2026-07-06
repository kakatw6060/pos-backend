package sqlite

import (
	"pos-backend/internal/model"
	"gorm.io/gorm"
)

// --- Generic Base Implementation ---
type baseRepo[T any] struct {
	db *gorm.DB
}

func (r *baseRepo[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepo[T]) GetByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepo[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepo[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *baseRepo[T]) List() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// --- User Repository ---
type userRepo struct {
	baseRepo[model.User]
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{baseRepo[model.User]{db: db}}
}

func (r *userRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// --- Category Repository ---
type categoryRepo struct {
	baseRepo[model.Category]
}

func NewCategoryRepository(db *gorm.DB) *categoryRepo {
	return &categoryRepo{baseRepo[model.Category]{db: db}}
}

// --- Product Repository ---
type productRepo struct {
	baseRepo[model.Product]
}

func NewProductRepository(db *gorm.DB) *productRepo {
	return &productRepo{baseRepo[model.Product]{db: db}}
}

func (r *productRepo) GetByCategoryID(categoryID uint) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// --- Order Repository ---
type orderRepo struct {
	baseRepo[model.Order]
}

func NewOrderRepository(db *gorm.DB) *orderRepo {
	return &orderRepo{baseRepo[model.Order]{db: db}}
}

func (r *orderRepo) CreateOrder(order *model.Order, items []model.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for i := range items {
			items[i].OrderID = order.ID
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}
		return nil
	})
}
