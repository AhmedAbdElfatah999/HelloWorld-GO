package main

import (
	"github.com/labstack/echo/v4"
	"HelloWorld/routes"
	"HelloWorld/db"
	"fmt"
)

func main() {

	// Initialize the database
	dbInstance, err := db.InitDB()
	if err != nil {
		// Handle the error, log, and exit if necessary
		panic(err)
	}
	defer dbInstance.Close()

	// Set the global database instance
	db.SetDBInstance(dbInstance)

	// Create an instance of Echo
	e := echo.New()

	// Initialize routes
	routes.InitRoutes(e,dbInstance)

	// Start the Echo server
	e.Start(":8080")
}