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
	router.GET("/user/:name" , Methods.GetUserDetail)
	router.POST("/user/new" , Methods.CreateNewUser) 
	router.DELETE("/users" , Methods.DeleteAllUsers)
	router.DELETE("/user/:name" , Methods.DeleteUserDetail)
	// router.PUT("/" , Methods.PutBasicDetail)

	router.Run(":3000")
}