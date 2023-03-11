package main

import (
	"log"
	"net/http"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/routes"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main(){
	//? create gin router
	router := gin.Default()

	//? Routers
	mainRouter := router.Group("/api")
	{
		routes.UsersRouter(mainRouter.Group("/user"))
		routes.PostsRouter(mainRouter.Group("/post"))
	}
	
	//? starting server
	log.Printf("Server started on port %s", utils.GetEnvVariable("PORT"))
	http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), router)
}