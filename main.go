package main

import (
	"log"
	"todolistgo/modules/item/transport/ginitem"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := "root:tuandat1601@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected:", db)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		item := v1.Group("/items")
		{
			item.POST("", ginitem.CreateItem(db))
			item.GET("", ginitem.ListItem(db))
			item.GET("/:id", ginitem.GetItem(db))
			item.PATCH("/:id", ginitem.UpdateItem(db))
			item.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}
	router.Run(":8080")
}






