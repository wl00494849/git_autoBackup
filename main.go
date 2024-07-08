package main

import (
	"fmt"
)

var _account string
var _token string

func init() {
	fmt.Printf("Enter your github account:")
	fmt.Scanln(&_account)
	fmt.Printf("Enter your github token:")
	fmt.Scanln(&_token)
}

func main() {
	for true {

	}
	fmt.Println(_account)
	fmt.Println(_token)
}

func cmd() {

}
