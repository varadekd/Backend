package main

import (
	"../src/Methods"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	fmt.Println("Starting JSON server at 3000")

	// Calling different methods to hanlde user
	router.GET("/users" , Methods.GetAllUsers)
	router.GET("/user/:name" , Methods.GetUserDetail) // Passing params
	router.POST("/user/new" , Methods.CreateNewUser) // Passing params 
	router.DELETE("/users" , Methods.DeleteAllUsers)
	router.DELETE("/user/:name" , Methods.DeleteUserDetail) // Passing params

	router.Run(":3000")
}