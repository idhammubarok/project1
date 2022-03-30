package main

import (
	"fmt"
	"log"
	"project1/book"
	"project1/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// mysql connection
	dsn := "root:@tcp(127.0.0.1:3306)/_pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	fmt.Println("DB Connected")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)       // di ambil dari repository
	bookService := book.NewService(bookRepository) // di ambil dari service
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1") // group versioning router

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.POST("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
