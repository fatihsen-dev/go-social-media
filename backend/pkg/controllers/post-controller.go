package controllers

import (
	"fmt"
	"net/http"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	// config.GetDB().Select([]string{"id", "title", "subtitle", "description", "owner", "created_at"}).Find(&posts)
	config.GetDB().Find(&posts)
	c.JSON(http.StatusOK, posts)
}

type Post struct {
	Title             string     `validate:"required,min=10,max=40"`
	Subtitle          string     `validate:"required,min=15,max=80"`
	Description       string     `validate:"required,min=100,max=600"`
}

func CreatePost(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		return
	}
	var input Post
	if err := c.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
  	}

	validate := validator.New()
	err := validate.Struct(&input)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s validation failed on field %s", err.Tag(), err.Field()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
  	}

	y, ok := id.(float64)
	if ok {
		post := models.Post{Title: input.Title, Subtitle: input.Subtitle, Description: input.Description,Owner: uint(y)}
		config.GetDB().Create(&post)
		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post create error"})
	}
}