package main

import (
	"../src/Methods"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)

// Om I find this particular thing a bit difficult to understand I understand the working but I am not getting any examples
func checkRoute() gin.HandlerFunc {
    return func (c *gin.Context) {
        fmt.Println("I am inside the middleware")
    }
}

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

	// Rendering HTML files (This is just to check how template are loaded, I will convert this into full working module as we go forward)
	router.LoadHTMLGlob("../resources/template/*.html")
	router.Static("/static", "../resources/static")  // Loading static files
	// router.LoadHTMLGlob("../resource1/template/*.tmpl")

	// This will load our index.html file
	router.GET("/" , func(c *gin.Context){
		c.HTML(200 , "index.html" , gin.H{})
	})
	
	// router.GET("/tmpl" , func(c *gin.Context){
	// 	c.HTML(200 , "index.tmpl" , gin.H{
	// 		"title" : "I am site title",
	// 		"body" : "I am the body of the page from the server side",
	// 	})
	// })


	// Calling different methods to hanlde user grouping this routes as version v1
	v1 := router.Group("/v1")
	v1.Use(checkRoute()) // Using middleware
	{
		v1.GET("/users" , Methods.GetAllUsers)
		v1.GET("/user/:name" , Methods.GetUserDetail) // Passing params
		v1.POST("/user/new" , Methods.CreateNewUser) // Passing params 
		v1.PUT("/user/:name" , Methods.UpdateUserDetail) // Passing params
		v1.DELETE("/user/:name" , Methods.DeleteUserDetail) // Passing params
	}

	// Running the server on port 3000
	router.Run(":3000")
}