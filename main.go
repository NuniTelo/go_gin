package main

import (
	"github.com/gin-gonic/gin"
	"firstProject/controllers"

)

func main() {

	router := gin.Default() //declare router

	v1 := router.Group("/api")
	{
		v1.GET("/user",controllers.GetAllUsers) //decleare all routes

	}

	router.Run() // listen and serve on 0.0.0.0:8080
}





