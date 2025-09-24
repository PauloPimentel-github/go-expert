package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println("Categoria:", category.Name)
		for _, product := range category.Products {
			println(" - Produto:", product.Name, " - Serial Number:", product.SerialNumber.Number)
		}
	}
}
