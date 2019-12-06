package main

import (
	"../src/Methods"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	fmt.Println("Starting JSON server at 3000")

	// Handling static routes
	router.LoadHTMLGlob("../resource/index.html")

	// Calling different methods to hanlde user grouping this routes as version v1
	v1 := router.Group("/v1")
	
	{
		v1.GET("/users" , Methods.GetAllUsers)
		v1.GET("/user/:name" , Methods.GetUserDetail) // Passing params
		v1.POST("/user/new" , Methods.CreateNewUser) // Passing params 
		v1.PUT("/user/:name" , Methods.UpdateUserDetail) // Passing params
		v1.DELETE("/users" , Methods.DeleteAllUsers)
		v1.DELETE("/user/:name" , Methods.DeleteUserDetail) // Passing params
	}

	// Running the server on port 3000
	router.Run(":3000")
}