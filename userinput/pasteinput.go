package userinput

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Pasteinput() {

	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Store lines in array
	lines := []string{}

	fmt.Println("THe OS is", runtime.GOOS)

	fmt.Println("Enter text (Ctrl+D to end input on Linux/Mac, Ctrl+Z then Enter on Windows):")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

		// CLean up
		line = strings.TrimSpace(line)

		// try a command
		if line == "exit" {
			break
		}
		if line == "" && len(line) == 0 {
			fmt.Printf("Empty line do you want to stop or continue\n\n")
			continue
		}

		line = line + " -- " + strconv.Itoa(len(line))
		lines = append(lines, line)

	}

	fmt.Println("\nPasted text:")
	for _, l := range lines {
		fmt.Println(l)
	}
}
