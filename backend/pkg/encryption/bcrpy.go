package encryption

import "golang.org/x/crypto/bcrypt"

func PasswordBcrypt(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
