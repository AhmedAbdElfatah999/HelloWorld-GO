package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"HelloWorld/models"
)
const (
    host     = "postgresContainer2"
    port     = 5432
    user     = "postgres"
    password = "pass123"
    dbname   = "marketplace"
)
var (
	Instance *sql.DB
	once       sync.Once
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
    var err error
    once.Do(func() {
        connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)

        Instance, err = sql.Open("postgres", connStr)
        if err != nil {
            return
        }

        if err := Instance.Ping(); err != nil {
			
            return
        }
    })

    return Instance, err
}

// SetDBInstance sets the global database instance
func SetDBInstance(db *sql.DB) {
	Instance = db
}

// CreateProduct adds a new product to the "Products" table
func CreateProduct(db *sql.DB, p models.Product) error {
	_, err := db.Exec("INSERT INTO Products(ID, NAME, TYPE, CATEGORY, ATICLES) VALUES($1, $2, $3, $4, $5)",
		p.ID, p.Name, p.Type, p.Category, p.Articles)
	return err
}

// ReadAllProducts retrieves all products from the "Products" table
func ReadAllProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Category, &p.Articles)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// ReadProductByID retrieves a product by its ID from the "Products" table
func ReadProductByID(db *sql.DB, id int) (models.Product, error) {
	var p models.Product
	err := db.QueryRow("SELECT * FROM Products WHERE ID = $1", id).
		Scan(&p.ID, &p.Name, &p.Type, &p.Category, &p.Articles)
	return p, err
}

// UpdateProduct updates a product in the "Products" table
func UpdateProduct(db *sql.DB, p models.Product) error {
	_, err := db.Exec("UPDATE Products SET NAME = $2, TYPE = $3, CATEGORY = $4, ATICLES = $5 WHERE ID = $1",
		p.ID, p.Name, p.Type, p.Category, p.Articles)
	return err
}

// DeleteProduct removes a product from the "Products" table by ID
func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM Products WHERE ID = $1", id)
	return err
}