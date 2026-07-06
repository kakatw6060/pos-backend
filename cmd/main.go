package main

import (
	"log"
	"pos-backend/internal/handler"
	"pos-backend/internal/model"
	"pos-backend/internal/repository/gorm"
	"pos-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init Database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 2. Init Repositories
	userRepo := gorm.NewUserRepository(db)
	catRepo := gorm.NewCategoryRepository(db)
	prodRepo := gorm.NewProductRepository(db)
	orderRepo := gorm.NewOrderRepository(db)

	// 3. Init Handlers
	userH := &handler.BaseHandler[model.User]{Repo: userRepo}
	catH := &handler.BaseHandler[model.Category]{Repo: catRepo}
	prodH := &handler.BaseHandler[model.Product]{Repo: prodRepo}
	orderH := &handler.OrderHandler{
		BaseHandler: handler.BaseHandler[model.Order]{Repo: orderRepo},
		OrderRepo:   orderRepo,
	}

	// 4. Setup Router
	r := gin.Default()

	// User API
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userH.Create)
		userGroup.GET("/", userH.List)
		userGroup.GET("/:id", userH.Get)
		userGroup.PUT("/:id", userH.Update)
		userGroup.DELETE("/:id", userH.Delete)
	}

	// Category API
	catGroup := r.Group("/categories")
	{
		catGroup.POST("/", catH.Create)
		catGroup.GET("/", catH.List)
		catGroup.GET("/:id", catH.Get)
		catGroup.PUT("/:id", catH.Update)
		catGroup.DELETE("/:id", catH.Delete)
	}

	// Product API
	prodGroup := r.Group("/products")
	{
		prodGroup.POST("/", prodH.Create)
		prodGroup.GET("/", prodH.List)
		prodGroup.GET("/:id", prodH.Get)
		prodGroup.PUT("/:id", prodH.Update)
		prodGroup.DELETE("/:id", prodH.Delete)
	}

	// Order API
	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("/", orderH.CreateWithItems)
		orderGroup.GET("/", orderH.List)
		orderGroup.GET("/:id", orderH.Get)
		orderGroup.PUT("/:id", orderH.Update)
		orderGroup.DELETE("/:id", orderH.Delete)
	}

	// Note: go-admin integration typically happens by adding their middleware 
	// and routing to their admin engine. Given the complexity of their scaffolding,
	// this backend is now fully compatible as it uses GORM models.
	// To add go-admin:
	// 1. install go-admin CLI: go install github.com/go-admin-team/go-admin/v2/bin/go-admin@latest
	// 2. run: go-admin init
	// 3. integrate the generated Admin engine into this Gin router.

	log.Println("POS Backend running on :8080")
	r.Run(":8080")
}
