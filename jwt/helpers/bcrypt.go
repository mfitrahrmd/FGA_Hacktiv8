package helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const SALT = 8

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), SALT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(hashedPassword)
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err.Error())

		return false
	}

	return true
}
