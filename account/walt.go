package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts []Account `json:"accouns"`
	UpdateAt time.Time `json:"updateAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts: []Account{},
		UpdateAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	vault.UpdateAt = time.Now()
}

func (vault *Vault) ToByte() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}

func (vault *Vault) ReadVault() error {
	file, err := files.ReadFile()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(file, vault)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (vault *Vault) PrintVault() {
	fmt.Println(*vault)
}

func (vault *Vault) FindAccount(strFind string) {
	var count uint
	strFind = strings.ToLower(strFind)
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, strFind) {
			fmt.Println(value.Url, "|", value.Login, "|", value.Password)
			count++
		}
	}
	fmt.Print("Find: ", count, " accounts\n\n")
}
