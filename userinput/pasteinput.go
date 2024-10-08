package userinput

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Pasteinput() {

	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Store lines in array
	lines := []string{}
	// Operating system
	operatingSystem := runtime.GOOS

	if operatingSystem == "windows" {

		color.Cyan("::: Usage :::\n")
		color.Blue("1. Enter or Paste text\n2. Press Ctrl+Z\n3. Press Enter:\n\n")
	} else if operatingSystem == "linux" || operatingSystem == "darwin" {

		fmt.Print("1. Enter or Paste text\n2. Press Ctrl+D\n3. Press Enter:\n\n")
	} else {

		fmt.Print("1. Enter or Paste text\n2. Press Ctrl+D on Linux and Mac, Press Ctrl+Z on Windows\n3. Press Enter:\n\n")
	}

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

		// CLean up
		line = strings.TrimSpace(line)

		// try a command
		if line == "exit -n" {
			break
		}
		if line == "" && len(line) == 0 {
			color.Red("Empty line detected:")
			fmt.Printf("If you want to stop type \"exit -n\" \n\n")
			continue
		}

		line = line + " -- " + strconv.Itoa(len(line))
		lines = append(lines, line)

	}

	// fmt.Println("\nPasted text:")
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
	var myOutput string
	for _, line := range lines {
		myOutput = myOutput + line + "\n"
	}
	fmt.Println("\nPasted text in Block:")
	fmt.Println(myOutput)

}
