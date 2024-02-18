package manager

import (
	"encoding/json"
	"log"
	"os"
)

func SaveFile(passwords []PasswordEntry) {
	data, err := json.Marshal(passwords)
	if err != nil {
		panic(err)
	}

	dataEncrypted := Encrypt(data)

	log.Printf("Saving %d passwords", len(passwords))

	err = os.WriteFile(".data.txt", dataEncrypted, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadFile() []PasswordEntry {
	dataEncrypted, err := os.ReadFile(".data.txt")
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

func OpenLogFile() *os.File {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
