package models

type InventoryResponse struct {
	ProductID uint
	Stock     int
}

type InventoryUpdate struct {
	Productid uint `json:"product_id"`
	Stock     int  `json:"stock"`
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}

type Inventories struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
type AddInventories struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
