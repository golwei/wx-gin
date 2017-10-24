package main

import (
	"fmt"

	"wx-gin/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:654654@tcp(db:3306)/nideshop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.SingularTable(true)
	// Migrate the schema
	//db.AutoMigrate(&NideshopUserLevel{})

	// Create
	//db.Create(&NideshopUserLevel{Name: "golwei"})

	// Read
	var goods []models.NideshopAttribute
	db.Where("name LIKE ?", "%材质%").Find(&goods)
	fmt.Println(goods)
	//	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	//	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//	db.Delete(&product)
}
