package models

type PasswordEntry struct {
	Service  string
	Username string
	Password string
}

type ShowPasswordsModel struct {
	Passwords []PasswordEntry
}

func (m ShowPasswordsModel) View() string {

	s := "Your passwords:\n\n"

	for _, p := range m.Passwords {
		s += p.Service + "\n"
		s += "  Username: " + p.Username + "\n"
		s += "  Password: " + p.Password + "\n\n"
	}

	s += "Press q to quit."
	return s
}
