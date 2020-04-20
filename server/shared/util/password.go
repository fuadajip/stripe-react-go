package util

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(plain *string) (*string, error) {

	plainByte := []byte(*plain)

	hash, err := bcrypt.GenerateFromPassword(plainByte, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return aws.String(string(hash)), nil
}

func CompareHashPassword(plain *string, hashed *string) bool {

	byteHashed := []byte(*hashed)
	bytePlain := []byte(*plain)

	err := bcrypt.CompareHashAndPassword(byteHashed, bytePlain)
	if err != nil {
		return false
	}

	return true
}
