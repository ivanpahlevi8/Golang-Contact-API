package model

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *Account) SetUsername(username string) {
	m.Username = username
}

func (m *Account) SetPassword(password string) {
	m.Password = password
}

func (m *Account) GetUsername() string {
	return m.Username
}

func (m *Account) GetPassword() string {
	return m.Password
}
