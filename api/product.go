package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"HelloWorld/models"
	"net/http"
)

type productManager interface {
	Add(product *models.Product) error
	ReadAll(c echo.Context) ([]models.Product, error)
}

type ProductRouter struct {
	productManager productManager
}

func NewProductRouter(productMgr productManager) *ProductRouter {
	return &ProductRouter{
		productManager: productMgr,
	}
}

// CreateProduct
func (p *ProductRouter) CreateProduct(c echo.Context) error {
	// note: tight coupling example
	//mgr := product2.NewManager()

	// Bind the request body to the Product struct
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		fmt.Printf("Error While Binding the product\n")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := p.productManager.Add(product); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Respond with a JSON message
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product created successfully",
		"name":    product.Name,
	})
}
func (p *ProductRouter) ReadAll(c echo.Context) error {
    products, err := p.productManager.ReadAll(c)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, products)
}