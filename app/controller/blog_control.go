package controller

import (
	"net/http"

	"blogPost/model"
	"blogPost/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogController struct {
	DB *gorm.DB
}

func NewBlogController(db *gorm.DB) *BlogController {
	return &BlogController{
		DB: db,
	}
}

func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	var blogs []model.Blog
	c.DB.Find(&blogs)
	ctx.JSON(http.StatusOK, view.BlogView{
		Blog: &model.Blog{},
	})
}

func (c *BlogController) GetBlogPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	var blog model.Blog

	c.DB.First(&blog, id)

	ctx.JSON(http.StatusOK, view.BlogView{
		Blog: &blog,
	})
}

func (c *BlogController) CreateBlogPost(ctx *gin.Context) {
	var newBlog model.Blog
	if err := ctx.BindJSON(&newBlog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.DB.Create(&newBlog)

	ctx.JSON(http.StatusCreated, view.BlogView{
		Blog: &newBlog,
	})
}

func (c *BlogController) UpdateBlogPost(ctx *gin.Context) {
	id := ctx.Param("id")
	var newBlog model.Blog

	if err := ctx.BindJSON(&newBlog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.DB.Model(&model.Blog{}).Where("id = ?", id).Updates(&newBlog)

	ctx.JSON(http.StatusOK, view.BlogView{
		Blog: &newBlog,
	})
}

func (c *BlogController) DeleteBlogPost(ctx *gin.Context) {
	id := ctx.Param("id")

	c.DB.Delete(&model.Blog{}, id)

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog post deleted"})
}
