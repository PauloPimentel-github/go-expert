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
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{Name: "Notebook", Price: 5000})

	// create batch
	// products := []Product{
	// 	{Name: "Monitor", Price: 1500},
	// 	{Name: "Mouse", Price: 50},
	// 	{Name: "Teclado", Price: 100},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Monitor")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %.2f\n", product.Name, product.Price)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %.2f\n", product.Name, product.Price)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %.2f\n", product.Name, product.Price)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %.2f\n", product.Name, product.Price)
	// }

	var product Product
	db.First(&product, 1)
	product.Name = "Notebook Gamer PRO"
	db.Save(&product)

	var product2 Product
	db.First(&product2, 1)
	fmt.Println(product2)
	db.Delete(&product2)

}
