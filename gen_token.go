package aauth

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenToken(username, password string) string {
	hash := sha256.Sum256([]byte(username+password))
	return hex.EncodeToString(hash[:])
}