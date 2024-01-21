package main

import (
	"github.com/labstack/echo/v4"
	"HelloWorld/routes"
)

func main() {
	// Create an instance of Echo
	e := echo.New()

	// Initialize routes
	routes.InitRoutes(e)

	// Start the Echo server
	e.Start(":8080")
}