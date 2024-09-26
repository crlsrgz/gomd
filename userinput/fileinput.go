package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func Fileinput() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	userOs := runtime.GOOS
	fmt.Println("the OS is ", userOs)
	fmt.Println("the OS directory is ", dir)

	fmt.Println("give me the file path: ")
	// var filesPath string
	// fmt.Scan(&userOption)
	filesPath := ""
	_, errPath := fmt.Scanln(&filesPath)
	if errPath != nil {
		log.Fatal(errPath)
	}
	filesPath = dir + "/" + filesPath
	fmt.Println("filesPath is ", filesPath)
	// open file
	f, err := os.Open(filesPath)

	if err != nil {
		log.Fatal(err)
	}

	//close the file at the end
	defer f.Close()

	// read line by line
	scanner := bufio.NewScanner(f)

	outResult := ""

	for scanner.Scan() {
		// fmt.Printf("line Nr. : %s\n", scanner.Text())
		lineText := scanner.Text()
		if strings.Contains(lineText, "## ") && !strings.Contains(lineText, "### ") {

			outResult = fmt.Sprintf("%s%s\n", outResult, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}
	fmt.Print(outResult)
}
