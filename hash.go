package githublibrary

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hashBytes, err :=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed tp generate hash password from string: %v",err)
	}
	return string(hashBytes), nil
}

func VerifyHashPassword(hash string, password string)(bool, error){
	err :=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, fmt.Errorf("failed to check password : %v", err)
	}
	return true, nil
}