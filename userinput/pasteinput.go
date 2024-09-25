package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Pasteinput() {
	fmt.Println("User wants to write or paste something")
	fmt.Println("gimme some text: ")

	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println("the text is:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
