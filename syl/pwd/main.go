package main

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
	"log"
	"os"
)

func main() {
	// Generate a password that is 30 characters long with 8 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	pwd, err := password.Generate(30, 8, 10, false, false)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(pwd)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(dir)
}
