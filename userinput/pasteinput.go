package userinput

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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

		color.Blue("::: Usage :::\n")
		color.Blue(MessageInstructionsWindows)

	} else if operatingSystem == "linux" || operatingSystem == "darwin" {

		fmt.Print(MessageInstructionLinux)
	} else {

		fmt.Print(MessageInstructionOther)
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
			fmt.Printf("Carry on pasting or typing; to stop type \"exit -n\" \n\n")
			continue
		}

		// line = line + " -- " + strconv.Itoa(len(line))
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

	fmt.Printf("\nIndex ready: \n")
	fmt.Println(myOutput)

}
