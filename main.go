package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func generatePassword(length int, includeUpper, includeNumbers, includeSpecial bool) (string, error) {
	charSet := lowerChars

	if includeUpper {
		charSet += upperChars
	}
	if includeNumbers {
		charSet += numberChars
	}
	if includeSpecial {
		charSet += specialChars
	}

	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password[i] = charSet[index.Int64()]
	}

	return string(password), nil
}

func main() {
	var length int
	var includeUpper, includeNumbers, includeSpecial bool

	fmt.Print("Enter password length: ")
	fmt.Scan(&length)

	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scan(&includeUpper)

	fmt.Print("Include numbers? (true/false): ")
	fmt.Scan(&includeNumbers)

	fmt.Print("Include special characters? (true/false): ")
	fmt.Scan(&includeSpecial)

	password, err := generatePassword(length, includeUpper, includeNumbers, includeSpecial)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}
