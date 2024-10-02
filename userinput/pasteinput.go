package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pasteinput() {

	reader := bufio.NewReader(os.Stdin)
	lines := []string{}

	fmt.Println("Enter multiline text (type 'exit' to finish):")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

		// Remove trailing newline character
		line = strings.TrimSpace(line)

		if line == "" && len(lines) > 0 {
			break
		}

		lines = append(lines, line)
	}

	fmt.Println("\nPasted text:")
	for _, l := range lines {
		fmt.Println(l)
	}
}
