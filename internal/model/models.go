package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Role     string `json:"role" gorm:"default:'cashier'"` // admin, cashier
}

type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	gorm.Model
	Name       string  `json:"name" gorm:"not null"`
	Price      float64 `json:"price" gorm:"not null"`
	Stock      int     `json:"stock" gorm:"default:0"`
	CategoryID uint    `json:"category_id" gorm:"not null"`
	Category   Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}

type Order struct {
	gorm.Model
	UserID      uint        `json:"user_id" gorm:"not null"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Status      string      `json:"status" gorm:"default:'pending'"` // pending, completed, cancelled
	OrderItems  []OrderItem `json:"order_items,omitempty"`
	User        User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	UnitPrice float64 `json:"unit_price" gorm:"not null"`
	Product   Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
