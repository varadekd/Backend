package main

import (
	"../src/Methods"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)


func main(){
	// Writing logs to file 
	// If you want to print log only to console then please skip the below code from here 

	gin.DisableConsoleColor() // This will disbale log color (You can skip this if you want to logs with console colors)
	// gin.ForceConsoleColor() // Use this line when you want to print logs using console color

	file , err := os.Create("server.log")

	if err != nil{
		fmt.Println(err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(file , os.Stdout) // This line helps us to write the logs in both file and console.
	//gin.DefaultWriter = io.MultiWriter(file) // This will help you to write log only to file 
	// Till here 

	router := gin.Default()
	fmt.Println("Starting JSON server at 3000")

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