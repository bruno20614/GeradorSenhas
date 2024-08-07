package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	digits  = "0123456789"
	symbols = "!@#$%^&*()-_=+[]{}<>?/|\\"
)

func main() {
	var length int
	fmt.Println("Write the password length")
	fmt.Scan(&length)

	password := generatePassword(length)
	fmt.Println("Password Created", password)
}

func generatePassword(length int) string {
	charset := letters + digits + symbols
	password := make([]byte, length)

	for i := range password {
		randomChar, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomChar.Int64()]
	}
	return string(password)
}
