package manager

// New model for testing the app
func TestModel() model {
	return model{
		Passwords: []passwordEntry{
			{Service: "Twitter", Username: "user1", Password: "password1"},
			{Service: "Facebook", Username: "user2", Password: "password2"},
			{Service: "Instagram", Username: "user3", Password: "password3"},
		},
	}
}
