package models

// New model for testing the app
func TestModel() ShowPasswordsModel {
	return ShowPasswordsModel{
		Passwords: []PasswordEntry{
			{Service: "Twitter", Username: "user1", Password: "password1"},
			{Service: "Facebook", Username: "user2", Password: "password2"},
			{Service: "Instagram", Username: "user3", Password: "password3"},
		},
	}
}
