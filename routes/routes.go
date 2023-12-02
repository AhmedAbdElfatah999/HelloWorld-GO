package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"HelloWorld/handlers"
)

// InitRoutes initializes all routes
func InitRoutes(e *echo.Echo, instance *sql.DB) {
	// Product routes
	e.POST("/products", handlers.CreateProductHandler)
	e.GET("/products", handlers.ReadAllProductsHandler)
	e.GET("/products/:id", handlers.ReadProductByIDHandler)
	e.PUT("/products/:id", handlers.UpdateProductHandler)
	e.DELETE("/products/:id", handlers.DeleteProductHandler)
}