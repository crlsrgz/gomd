package main

import (
	"fmt"
	"gocli/userinput"
)

func main() {
	fmt.Println("User options")
	fmt.Println("do you want to use a file(f) do you want to paste text (p)?")
	userOption := ""
	// fmt.Scan(&userOption)
	fmt.Scanln(&userOption)

	if userOption == "f" {

		fmt.Printf("User chosed %s\n", userOption)
		userinput.Fileinput()

	} else if userOption == "p" {

		fmt.Printf("User chosed %s\n", userOption)
		userinput.Pasteinput()

	} else {
		fmt.Println("wrong input")
	}

}
