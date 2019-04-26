package controllers

import (
	"firstProject/Models"
	"github.com/gin-gonic/gin"
)

var listOfUsers []Models.User   //this is a list of users we will declare here

func createAListOfUsers() {

	//add some fake data to the list of users
	listOfUsers = append(listOfUsers,Models.User{FirstName:"Nuni",LastName:"telo",Address:&Models.Address{Street: "5 Maji",ZipCode:"1001"}})
	listOfUsers = append(listOfUsers,Models.User{FirstName:"Nuni1",LastName:"telo",Address:&Models.Address{Street: "5 Maji",ZipCode:"1001"}})
	listOfUsers = append(listOfUsers,Models.User{FirstName:"Nuni2",LastName:"telo",Address:&Models.Address{Street: "5 Maji",ZipCode:"1001"}})

}

func GetAllUsers(c *gin.Context) {
	createAListOfUsers()
	c.JSON(200, gin.H{
		"status":true,
		"users":listOfUsers,
	})
}
