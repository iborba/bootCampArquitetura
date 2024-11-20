package models

// Product represents a product entity in the database.
// @Description A product model
type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"` // Product ID
	Name        string  `json:"name"`                 // Product Name
	Description string  `json:"description"`          // Product Description
	Price       float64 `json:"price"`                // Product Price
	Stock       int     `json:"stock"`                // Product Stock
}
