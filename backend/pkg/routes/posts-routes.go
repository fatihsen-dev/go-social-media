package routes

import (
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/controllers"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func PostsRouter(postRouter *gin.RouterGroup) {
	postRouter.GET("/all", middlewares.AuthMiddleware(), controllers.GetPosts)
	postRouter.POST("/create", middlewares.AuthMiddleware(), controllers.CreatePost)
}