package manager

import (
	"encoding/json"
	"fmt"
	"os"

	m "github.com/playjeri/passutasku/src/manager/models"
)

func SaveFile(passwords []m.PasswordEntry) {
	data, err := json.Marshal(passwords)
	if err != nil {
		panic(err)
	}

	dataEncrypted := Encrypt(data)

	fmt.Println("Saving data to file")

	err = os.WriteFile(".data.txt", dataEncrypted, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadFile() []m.PasswordEntry {
	dataEncrypted, err := os.ReadFile(".data.txt")
	if err != nil {
		return []m.PasswordEntry{}
	}

	data := Decrypt(dataEncrypted)

	var passwords []m.PasswordEntry
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
