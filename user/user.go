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