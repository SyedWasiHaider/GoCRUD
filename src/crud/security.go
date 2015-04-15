package main

import "golang.org/x/crypto/bcrypt"

func getDigest(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func isCorrectPassword(password []byte, digest []byte) bool {
	if bcrypt.CompareHashAndPassword(digest, password) == nil {
		return true
	} else {
		return false
	}
}
