package main

import (
	"fmt"
	"os"

	"github.com/certyclick_verify/core"
)

func main() {
	args := os.Args
	file, err := core.OpenFile(args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	hashed, err := core.CalculateHash(file)
	if err != nil {
		fmt.Println("Error calculating file hash:", err)
		return
	}

	if fmt.Sprintf("%x", hashed) == args[2] {
		fmt.Println("The file hash matches the provided hash.")
	} else {
		fmt.Println("The file hash does not match the provided hash.")
		fmt.Println("Provided hash:", args[2])
		fmt.Println("Calculated hash:", fmt.Sprintf("%x", hashed))
	}
}
