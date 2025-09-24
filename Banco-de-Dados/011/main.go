package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{Name: "Notebook", Price: 5000})

	var product Product
	db.First(&product, 1)
	product.Name = "Notebook Dell"
	db.Save(&product)

	var product2 Product
	db.First(&product2, 1)
	fmt.Println(product2)
	// soft delete
	db.Delete(&product2)

}
