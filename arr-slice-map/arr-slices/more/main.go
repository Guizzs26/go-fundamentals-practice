package main

import (
	goprogramminglanguagebook "arr-slice-map/arr/more/go-programming-language-book"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	// -------------------------- //

	/*
	   %d → Numeric representation of the byte array.

	   %x → Hexadecimal representation of the hash (common way to display digests).

	   %t → Compares c1 and c2 (true if they are equal, false if they are different).

	   %T → Prints the type of c1 ([32]byte).
	*/

	// Exercise 4.1
	hash1 := sha256.Sum256([]byte("x")) // Calculate SHA256 for input x
	hash2 := sha256.Sum256([]byte("X")) // Calculate SHA256 for input X
	fmt.Printf("%x\n%x\n%t\n%T\n", hash1, hash2, hash1 == hash2, hash1)
	fmt.Printf("%d\n%d\n%t\n%T\n", hash1, hash2, hash1 == hash2, hash1)
	fmt.Println(goprogramminglanguagebook.BitDifference(hash1, hash2))

	// Exercise 4.2
	hsa := flag.String("hsa", "sha256", "Choose between sha256, sha384 or sha512 hash algorithms")
	flag.Parse()
	hash := goprogramminglanguagebook.GenerateHash(os.Stdin, *hsa)
	fmt.Printf("%x\n", hash)
}
