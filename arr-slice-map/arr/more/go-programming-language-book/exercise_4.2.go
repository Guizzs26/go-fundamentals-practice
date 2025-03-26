package goprogramminglanguagebook

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func GenerateHash(input io.Reader, hsa string) []byte {
	data, err := io.ReadAll(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	switch hsa {
	case "sha384":
		hash := sha512.Sum384(data)
		return hash[:]
	case "sha512":
		hash := sha512.Sum512(data)
		return hash[:]
	default:
		hash := sha256.Sum256(data)
		return hash[:]
	}
}
