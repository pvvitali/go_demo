package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"time"
)

func main() {

	var variant int = 0
menu:
	for {
		fmt.Println("Menu")
		fmt.Println("1 - Create account")
		fmt.Println("2 - Search account")
		fmt.Println("3 - Delete account")
		fmt.Println("4 - Exit")

		fmt.Scan(&variant)

		switch variant {
		case 1:
			createAccount()
		case 2:
			searchAccount()
		case 3:
			deleteAccount()
		case 4:
			break menu
		}
	}

}

func createAccount() {

	user1 := account.Account{
		Login:     "Vasia",
		Password:  "1234",
		Url:       "http://vasia.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	file, err := user1.ToByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	files.WriteFile(file, "log.txt")

}

func searchAccount() {

}

func deleteAccount() {

}
