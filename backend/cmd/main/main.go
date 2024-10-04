package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/routes"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	mainRouter := router.Group("/api")
	{
		routes.UsersRouter(mainRouter.Group("/user"))
		routes.PostsRouter(mainRouter.Group("/post"))
	}

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Api Ready"})
	})

	log.Printf("Server started on port %s", utils.GetEnvVariable("PORT"))
	http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), router)
}
