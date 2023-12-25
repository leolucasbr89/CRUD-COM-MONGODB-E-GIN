package utils

import "golang.org/x/crypto/bcrypt"

func CriptografaSenha(senha string) (string, error) {
	senhaBytes := []byte(senha)
	hash, err := bcrypt.GenerateFromPassword(senhaBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}