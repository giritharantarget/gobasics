package main

import "fmt"


func Struct(){

	type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u  user;
	u.ID =42
	u.FirstName = "Sam"
	fmt.Println(u)


}