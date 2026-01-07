package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("String was writing")
}

func ReadFile() ([]byte, error) {
	data, err := os.ReadFile("log.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}
