package main

import (
	"fmt"
	"os"
)

func main() {
	// ReadFile reads the entire file and returns its contents
	// as a byte slice.
	data, err := os.ReadFile("/home/teja/Desktop/test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string to print it.
	fileContent := string(data)
	fmt.Println(fileContent)
}
