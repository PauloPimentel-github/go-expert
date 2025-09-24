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
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Esporte"}
	db.Create(&category)

	// create product with category
	product := Product{
		Name:       "Chuteira Nike",
		Price:      250.00,
		CategoryID: category.ID,
	}
	db.Create(&product)

	// create serial number with product
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	}
	db.Create(&serialNumber)

	// relationship 1:N
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
