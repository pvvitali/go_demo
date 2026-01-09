package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

var vaultNew = account.NewVault()

func main() {
	vaultNew.ReadVault()

	var variant int = 0
menu:
	for {
		fmt.Println("Menu")
		fmt.Println("1 - Create account")
		fmt.Println("2 - Search account")
		fmt.Println("3 - Delete account")
		fmt.Println("4 - Exit")

		fmt.Scanln(&variant)

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

	user1, err := account.NewAccount()
	if err != nil {
		fmt.Println(err)
		return
	}

	vaultNew.AddAccount(user1)

	file, err := vaultNew.ToByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	files.WriteFile(file, "log.json")

}

func searchAccount() {
	searchString, _ := account.PromtAccount()
	vaultNew.ReadVault()

	vaultNew.FindAccount(searchString)
	//vaultNew.PrintVault()
}

func deleteAccount() {

}
