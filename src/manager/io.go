package manager

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveFile(passwords []PasswordEntry) {
	data, err := json.Marshal(passwords)
	if err != nil {
		panic(err)
	}

	dataEncrypted := Encrypt(data)

	fmt.Println("Saving data to file")

	err = os.WriteFile("data.txt", dataEncrypted, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadFile() []PasswordEntry {
	dataEncrypted, err := os.ReadFile("data.txt")
	if err != nil {
		return []PasswordEntry{}
	}

	data := Decrypt(dataEncrypted)

	var passwords []PasswordEntry
	err = json.Unmarshal(data, &passwords)
	if err != nil {
		panic(err)
	}

	return passwords
}
