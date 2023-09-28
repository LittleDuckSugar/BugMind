package models

import "golang.org/x/crypto/bcrypt"

type Password struct {
	PasswordStr string `json:"password"`
}

// EncryptPassword do what is means, encrypt
func (password *Password) Encrypt() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password.PasswordStr), 14)
	password.PasswordStr = string(bytes)
}

func (password *Password) CheckPassword(passwordStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(password.PasswordStr), []byte(passwordStr))
}
