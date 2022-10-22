package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // 继承 gorm 的一个结构体
	Code       string
	Price      uint
}

var db *gorm.DB

func initDB() {
	dsn := "root:@/go_db?charset=utf8mb4"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("初始化失败 err:%v\n", err)
		return
	}
	fmt.Printf("初始化成功\n")
}

func createTable() {
	// 迁移 schema，会自动创建表
	db.AutoMigrate(&Product{})
}

func Create() {
	product := Product{
		Code:  "42",
		Price: 100,
	}
	db.Create(&product)
}

func find() {
	var product Product
	// 根据整型数据查找
	db.First(&product, 1)
	fmt.Printf("%v\n", product)

	var product2 Product
	db.First(&product2, "code = ?", "42")
	fmt.Printf("%v\n", product2)
}

func update() {

	var product Product

	// 更新一个字段
	db.Model(&product).Update("Price", 200)

	// 更新多个字段(只更新非零值字段)
	db.Model(&product).Updates(Product{Price: 200, Code: "43"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "43"})
}

func delete() {
	var product Product
	db.Delete(&product, 1)
}

func main() {

	initDB()

	//createTable()

	// 创建数据
	Create()

	// 读取
	find()

	// 更新
	//update()

	// 删除
	//delete()
}
