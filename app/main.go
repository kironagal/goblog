package main

import (
	"blogPost/controller"
	"blogPost/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blogPost.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}
	model.Migrate(db)

	router := gin.Default()

	// Create a new blog controller
	blogController := controller.BlogController{
		DB: db,
	}

	// Register the controller with the router
	router.GET("/posts", blogController.GetAllBlogs)
	router.GET("/posts/:id", blogController.GetBlogPostById)
	router.POST("/posts", blogController.CreateBlogPost)
	router.PUT("/posts/:id", blogController.UpdateBlogPost)
	router.DELETE("/posts/:id", blogController.DeleteBlogPost)

	// Start the server
	router.Run("localhost:8080")
}
