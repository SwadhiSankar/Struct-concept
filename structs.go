package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time // nested struct 
}


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    var appUser User 
	appUser = User{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetails(&appUser)

	appUser.outputUserDetails()

	
}
//another way 
func (u User) outputUserDetails(){
 fmt.Print(u.firstName,u.lastName,u.birthdate,u.createdAt)
}
// func outputUserDetails(u *User){
// 	fmt.Print((*u).lastName,u.firstName,u.birthdate,u.createdAt)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}