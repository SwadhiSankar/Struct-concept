package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time // nested struct 
}


func (u *User) clearUserName(){
	u.firstName =""
	u.lastName =""

}

func newUser(firstName, lastName, birthDate string) (*User,error){
	if(firstName == "" || lastName =="" || birthDate == ""){
		return nil, errors.New("Missing Data")
	}
	return &User{
		firstName: firstName,
		lastName : lastName,
		birthdate : birthDate,
		createdAt: time.Now(),
	},nil
}
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    var appUser *User 
	// appUser = User{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	appUser,err := newUser(userFirstName,userLastName,userBirthdate)
  if(err !=nil){
	fmt.Print(err)
     return 
  }
	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetails(&appUser)

	appUser.outputUserDetails()
	appUser.clearUserName()
	fmt.Println("\n After clearing")
	appUser.outputUserDetails()

	
}
//another way 
func (u *User) outputUserDetails(){
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