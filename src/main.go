package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./exemples/00.lang")
	source := string(bytes)

	fmt.Printf("Code: %s\n", source)
}
