package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello!")

	user1 := account.Account{
		Login:     "Peter",
		Password:  "1234",
		Url:       "http://peter.com",
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
