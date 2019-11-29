package Methods

import (
	"fmt"
	// "io/ioutil"
	// "bytes"
	"strings"
	"github.com/gin-gonic/gin"	
)

// Creating struct of type user
type User struct{
	Name 		string 	`json:name`
	Age			int		`json:age`		
	Company		string 	`json:company`
	Location 	string 	`json:location`
}

// Creating a global variable for users to be used across the module
var Users = []User{
	{
		Name : "Kushagra",
		Age : 25,
		Company : "RobusTest",
		Location : "Banglore",
	},
	{
		Name : "Yash",
		Age : 25,
		Company : "Commutatus",
		Location : "Chennai",
	},
	{
		Name : "Naveen",
		Age : 25,
		Company : "TCS",
		Location : "Banglore",
	},
}

func GetAllUsers(c *gin.Context){
	fmt.Println("Sending values to the user")
	c.JSON(200, Users)
}

func GetUserDetail(c *gin.Context){
	fmt.Println("Getting user details")
	name := c.Param("name")
	for _ , value := range Users{
		// This will return if the user is found in our DB
		if strings.ToLower(value.Name) == strings.ToLower(name){
			c.JSON(200 , value)
			return
		}
	}

	c.JSON(200, gin.H{
		"message" : "No user with name: " + name + " exist in our data.",
	})
}