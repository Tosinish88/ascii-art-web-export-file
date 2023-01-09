package ascii

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Ascii() {
	input := os.Args
	if len(input) <= 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]" + "\n" + "EX: go run . something standard")
		return
	} else if len(input) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]" + "\n" + "EX: go run . something standard")
		return
	}
	//checks input that it is legal to the program
	SanitizeInput(input[1])

	//split input to get filename
	//split := strings.Split(input[3], "=")
	//filename := split[1]

	//checks input that it is a valid banner
	banner := banner_name(input[2])

	// send input to asci function
	PrintAscii(input[1], banner, "1")

	//prints the asci art
	//fmt.Println(myart)

	prg := "cat"
	arg1 := "output.txt"

	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

}

// function to print asci
func PrintAscii(input string, banner string, counter string) string {
	//content of output used for unit testing
	content := ""
	//read file

	lines := readfile(banner)

	//convert input to runes
	runes := ConvertToRune(input)

	// nested loop to print line by line depending on input.
	splittwo := "\r\n"
	words := strings.Split(string(runes), splittwo)
	for _, word := range words {
		for h := 1; h < 9; h++ {
			if word == "" {
				//fmt.Println("")
				content = content + "\n"
				break
			}
			for _, l := range []byte(word) {
				for i, line := range lines {
					if i == (int(l)-32)*9+h {
						content = content + line
						//fmt.Print(line)
					}
				}
			}
			content = content + "\n"
			//fmt.Println()
		}
	}
	// writes the content to output.txt and also returns the content to the main function
	writeToFile([]string{content}, "files/output"+counter+".txt")
	return string(content)

}

// function to read file
func readfile(name string) []string {
	//read full file
	file, err := os.Open(name)
	//error check
	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	}
	//close file
	defer file.Close()
	//reads full file content to data
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Could not read file!")
		os.Exit(2)
	}
	//split string on newline
	lines := strings.Split(string(data), "\n")

	return lines
}

// function to write content to file
func writeToFile(text []string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for _, line := range text {
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// function to convert string to rune
func ConvertToRune(word string) []rune {
	s := word
	r := []rune(s)
	return r
}

func SanitizeInput(input string) string {
	if input == "" {
		os.Exit(0)
	}

	for _, value := range input {
		if value > 128 || value < 0 {
			fmt.Println("You can not use non ascii characters!")
			os.Exit(0)
		}
	}

	//convert input to runes
	if input == "\\n" {
		fmt.Println("")
		os.Exit(0)
	}

	return input
}

func banner_name(banner string) string {
	if banner == "" {
		fmt.Println("Usage: go run . [STRING] [BANNER]" + "\n" + "EX: go run . something standard")
		os.Exit(0)
	} else if banner == "standard" {
		return "standard.txt"
	} else if banner == "shadow" {
		return "shadow.txt"
	} else if banner == "thinkertoy" {
		return "thinkertoy.txt"
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]" + "\n" + "EX: go run . something standard")
		os.Exit(0)
	}
	return ""
}
