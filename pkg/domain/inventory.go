package domain

type Inventories struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	//Image       string   `json:"image"`
	Size  string  `json:"size" gorm:"size:5;default:'M';check:size IN ('S', 'M', 'L', 'XL', 'XXL')"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}
type Images struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	InventoryID uint        `json:"inventories_id"`
	Inventory   Inventories `json:"product" gorm:"foreignKey:InventoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ImageUrl    string      `json:"image_url"`
}
