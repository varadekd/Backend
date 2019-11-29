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

// Fetches all the user details
func GetAllUsers(c *gin.Context){
	fmt.Println("Sending values to the user")
	c.JSON(200, Users)
}

// Fetches a specific user detail
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
	// This will return the response in JSON format
	c.JSON(200, gin.H{
		"message" : "No user with name: " + name + " doesn't exist in our data.",
	})
}

// Creates a new user
// func CreateNewUser(c *gin.Context){

// }

// Delete all the user details
func DeleteAllUsers(c *gin.Context){
	fmt.Println("Deleting users")
	Users = nil
	// This will return the response in string format
	c.String(200, "All users are deleted")
}

// Deleting a particular user detail
func DeleteUserDetail(c *gin.Context){
	name := c.Param("name")

	for index , value := range Users{
		// This will return if the user is found in our DB
		if strings.ToLower(value.Name) == strings.ToLower(name){
			Users = append(Users[:index] , Users[index+1:]...)
			c.JSON(200 , value)
			return
		}
	}
	
	// This will return the response as string type as not found
	c.String(200, "Unable to delete user with name: " + name + " because doesn't exist in our data.")
}

