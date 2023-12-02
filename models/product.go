package models

// Product represents the structure of the "Products" table
type Product struct {
    ID       int
    Name     string
    Type     string
    Category string
    Articles int
}