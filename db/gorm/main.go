package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	Code  string
	Price uint
	Type string
}
// https://gorm.io/zh_CN/docs/
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Code: "D42", Price: 100, Type: "wew"})

	// Read
	var product Product
	db.First(&product, 1) // 根据整形主键查找
	fmt.Println(product)
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	fmt.Println(product)

	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
