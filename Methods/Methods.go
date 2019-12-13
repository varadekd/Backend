package Methods

import (
	"strings"
	"github.com/gin-gonic/gin"	
)

// Creating struct of type user
type User struct{
	Name 		string 	`json:"name" binding:"required"`
	Age 		int 	`json:"age" binding:"required"`	
	Company		string 	`json:"company" binding:"required"`
	Location 	string 	`json:"location"`
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
	c.JSON(200, Users)
}

// Fetches a specific user detail
func GetUserDetail(c *gin.Context){
	name := c.Param("name")
	for _ , value := range Users{
		// This will return if the user is found in our DB
		if strings.ToLower(value.Name) == strings.ToLower(name){
			c.JSON(200 , value)
			return
		}
	}

	// If user is not found we will return 404 with message 
	c.JSON(404, gin.H{
		"message" : "No user with name: " + name + " doesn't exist in our data.",
	})
}

// Creates a new user
func CreateNewUser(c *gin.Context){
	var data User

	// Using ShouldBindJSON to check the functionality and handle error manually
	err := c.ShouldBindJSON(&data)

	if err != nil {
		// In case of error we will send proper message to user
		c.JSON(400 , gin.H{
			"message" : err.Error(),
		})
	} else {
		for _ , value := range Users{
			// Checks for the duplicate entry in the users list if found return it
			if strings.ToLower(value.Name) == strings.ToLower(data.Name){
				c.JSON(409 , gin.H{
					"message" : "User already exist cannot create copy of the same user",
				})
				return
			}
		}
		Users = append(Users , data)
		c.JSON(200 , data)
	}
}

// Updation of new user
func UpdateUserDetail(c *gin.Context){
	var data User
	name := c.Param("name")
	// Trying the BindJSON to test how error are handled when something is missing or empty
	err := c.BindJSON(&data)
	if err != nil{
		c.JSON(400, gin.H{
			"message" : err.Error(),
		})
	} else{

		for index , value := range Users{
			
			if strings.ToLower(value.Name) == strings.ToLower(name){
				Users[index].Age = data.Age
				Users[index].Company = data.Company
				Users[index].Location = data.Location
				c.JSON(200 , value)
				return
			}
		}

		c.JSON(404 , gin.H{
			"message" : "Unale to find " + name + " in our DB for upation",
		})	
	}
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
	c.JSON(404, gin.H{
		"message" : "Unable to delete user with name: " + name + " because doesn't exist in our data.",
	})
}

