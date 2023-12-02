package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"HelloWorld/models"
	"HelloWorld/db"
)

// CreateProductHandler handles the creation of a product
func CreateProductHandler(c echo.Context) error {
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}

	err := db.CreateProduct(db.Instance, *p) // Assuming db.Instance is your initialized database instance
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, p)
}

// ReadAllProductsHandler handles the retrieval of all products
func ReadAllProductsHandler(c echo.Context) error {
	products, err := db.ReadAllProducts(db.Instance) // Assuming db.Instance is your initialized database instance
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}

// ReadProductByIDHandler handles the retrieval of a product by ID
func ReadProductByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := db.ReadProductByID(db.Instance, id) // Assuming db.Instance is your initialized database instance
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

// UpdateProductHandler handles the update of a product
func UpdateProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}

	p.ID = id
	err := db.UpdateProduct(db.Instance, *p) // Assuming db.Instance is your initialized database instance
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, p)
}

// DeleteProductHandler handles the deletion of a product
func DeleteProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := db.DeleteProduct(db.Instance, id) // Assuming db.Instance is your initialized database instance
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}