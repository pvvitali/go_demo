package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var chars []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890#$%")
var lenChars int = len(chars)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount() (*Account, error) {
	var login_data string
	var password_data string
	var url_data string
	fmt.Print("Enter login: ")
	fmt.Scanln(&login_data)
	if login_data == "" {
		err := errors.New("Empty login")
		return nil, err
	}
	fmt.Print("Enter password: ")
	fmt.Scanln(&password_data)
	if password_data == "" {
		password_data = generatePassword(12)
	}
	fmt.Print("Enter url: ")
	fmt.Scanln(&url_data)
	_, err := url.ParseRequestURI(url_data)
	if err != nil {
		return nil, err
	}

	return &Account{
		Login:     login_data,
		Password:  password_data,
		Url:       url_data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func generatePassword(countChars int) string {
	var mas = make([]rune, 0, 20)
	for i := 0; i < countChars; i++ {
		numRand := rand.IntN(lenChars)
		mas = append(mas, chars[numRand])
	}
	return string(mas)
}

func PromtAccount() string {
	var searchStr string
	fmt.Print("Enter search account (Example: http://user.com): ")
	fmt.Scanln(&searchStr)
	fmt.Print("Searching: ", searchStr, "\n\n")
	return searchStr

}

func PromtDeleteAccount() string {
	var urlString string
	fmt.Print("Enter Url for delete (Example: http://user.com): ")
	fmt.Scanln(&urlString)
	fmt.Print("Deleting: ", urlString, "\n\n")
	return urlString

}
