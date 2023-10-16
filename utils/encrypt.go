package utils

import "golang.org/x/crypto/bcrypt"

// EncriptPass hashes the given password using bcrypt
// with a given cost and returns the hashed string.
func EncriptPass(pass string) (string, error) {
	// set the cost for bcrypt hashing
	costo := 8

	// generate the hash for the given password with bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	// return the hashed password as a string and any errors that occurred
	return string(bytes), err
}
