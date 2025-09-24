package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create category
	// category := Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// create product with category
	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      5000,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	// relationship 1:N
	var products []Product
	db.Preload("Category").Find(&products)

	for _, product := range products {
		println(product.Name, product.Category.Name)
	}
}
