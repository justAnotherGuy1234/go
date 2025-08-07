package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)

	if err != nil {
		fmt.Println("error hashing passwrord", err)
		return "", err
	}

	return string(res), nil

}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	if err != nil {
		fmt.Println("incorrect password")
		return false
	}
	return true
}
