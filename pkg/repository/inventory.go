package repository

import (
	"errors"
	"fmt"
	"strconv"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &inventoryRepository{DB: DB}

}
func (i *inventoryRepository) AddInventory(inventory models.AddInventories, url string) (models.InventoryResponse, error) {

	//var id uint
	query := `INSERT INTO inventories (category_id,product_name,size,stock,price,image)
	VALUES(?, ?, ?, ?, ?, ?);`

	i.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price, url)

	var InventoryResponse models.InventoryResponse

	// i.DB.Raw(`SELECT
	//     id as product_id,
	//     stock
	// FROM
	//     inventories
	// WHERE
	//     id=?`, id).Scan(&InventoryResponse)

	return InventoryResponse, nil

}
func (i *inventoryRepository) CheckInventory(pid uint) (bool, error) {
	var k int

	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error

	fmt.Println("i:", k)

	if err != nil {
		return false, err
	}
	if k == 0 {
		return false, err
	}
	return true, err

}
func (i *inventoryRepository) UpdateInventory(pid uint, stock int) (models.InventoryResponse, error) {
	fmt.Println("values:", pid, stock)

	//check the database connection
	if i.DB == nil {
		return models.InventoryResponse{}, errors.New("no datbase connection")

	}
	//upate the databse
	if err := i.DB.Exec("UPDATE inventories SET stock = stock + $1 WHERE id= $2", stock, pid).Error; err != nil {
		return models.InventoryResponse{}, err
	}
	//retrive the upate
	var newdetails models.InventoryResponse
	var newstock int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&newstock).Error; err != nil {
		fmt.Println("debug:1")
		return models.InventoryResponse{}, err
	}
	newdetails.ProductID = pid
	newdetails.Stock = newstock

	fmt.Println(newdetails)
	return newdetails, nil
}
func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)

	if err != nil {
		return errors.New("converting into integer not happend")
	}
	fmt.Println("this is the id :", id)
	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no record of that id exist")

	}
	return nil
}
func (i *inventoryRepository) ShowIndividualProducts(id string) (domain.Inventories, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return domain.Inventories{}, error
	}
	var product domain.Inventories
	err := i.DB.Raw(`
	SELECT *FROM inventories WHERE inventories.id =  ?`, uint(pid)).Scan(&product).Error
	if err != nil {
		return domain.Inventories{}, errors.New("error retrived record")

	}
	return product, nil
}
func (ad *inventoryRepository) ListProducts(page int, count int) ([]domain.Inventories, error) {
	//pagination pourose

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productDetails []domain.Inventories

	if err := ad.DB.Raw("SELECT id ,category_id,product_name,size ,stock,price FROM inventories limit ? offset ?", count, offset).Scan(&productDetails).Error; err != nil {
		return []domain.Inventories{}, err

	}
	return productDetails, nil
}
func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var l int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&l).Error; err != nil {
		return 0, err
	}
	return l, nil
}
