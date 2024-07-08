package main

import (
	"fmt"
)

func main() {
	var account string
	var token string

	fmt.Println("Enter your github account")
	fmt.Scanln(&account)
	fmt.Println("Enter your github token")
	fmt.Scanln(&token)

	fmt.Println(account)
	fmt.Println(token)

}
