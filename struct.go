package main

import (
	"fmt"

	"example.com/struct/user"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    var appUser *user.User
	
	appUser,err := user.NewUser(userFirstName,userLastName,userBirthdate)
  if(err !=nil){
	fmt.Print(err)
     return 
  }
	

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	fmt.Println("\n After clearing")
	appUser.OutputUserDetails()

	
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}