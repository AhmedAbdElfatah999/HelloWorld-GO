package routes

import (
	"github.com/labstack/echo/v4"
	"HelloWorld/handlers"
	"HelloWorld/database"
	"HelloWorld/db"
	"HelloWorld/api"

)



// InitRoutes initializes all routes
func InitRoutes(e *echo.Echo) {
	sqlBuilder := database.NewSqlBuilder()
// Initialize the database
dbInstance, err := db.InitDB()
if err != nil {
	// Handle the error, log, and exit if necessary
	panic(err)
}
//defer dbInstance.Close()
productHandler := handlers.NewHandler(dbInstance, sqlBuilder)
productRouter := api.NewProductRouter(productHandler)
	// Product routes
	e.POST("/products", productRouter.Create)
	e.GET("/products", productRouter.ReadAll)
	// e.GET("/products/:id", handlers.ReadProductByIDHandler)
	// e.PUT("/products/:id", handlers.UpdateProductHandler)
	// e.DELETE("/products/:id", handlers.DeleteProductHandler)
}