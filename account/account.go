package account

import (
	"encoding/json"
	"fmt"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) ToByte() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}
