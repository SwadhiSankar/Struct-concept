package user

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

type Admin struct {
	Email string
	Password string
	User
}


func (u *User) ClearUserName(){
	u.firstName =""
	u.lastName =""

}

func NewUser(firstName, lastName, birthDate string) (*User,error){
	if(firstName == "" || lastName =="" || birthDate == ""){
		return nil, errors.New("missing Data")
	}
	return &User{
		firstName: firstName,
		lastName : lastName,
		birthdate : birthDate,
		createdAt: time.Now(),
	},nil
}
func (u *User) OutputUserDetails(){
	fmt.Print(u.firstName,u.lastName,u.birthdate,u.createdAt)
   }

   func NewAdmin(email , password string)Admin{
	return Admin{
		Email: email,
		Password: password,
		User: User{
			firstName: "Swathi",
			lastName: "S",
			birthdate: "14/08/2002",
			createdAt: time.Now(),
		},
	}
   }