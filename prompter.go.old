package main

import (
    "fmt"
    "os"
	"strings"
	"bufio"
)

func main() {
	commands_bg := map[string]string{
		"\\blackbg" : "\\[\\e[40m\\]",
		"\\redbg" : "\\[\\e[41m\\]",
		"\\greenbg" : "\\[\\e[42m\\]",
		"\\brownbg" : "\\[\\e[43m\\]",
		"\\bluebg" : "\\[\\e[44m\\]",
		"\\purplebg" : "\\[\\e[45m\\]",
		"\\cyanbg" : "\\[\\e[46m\\]",
		"\\whitebg" : "\\[\\e[47m\\]",
	}

	commands := map[string]string{
		"\\black" : "\\[\\e[30m\\]",
		"\\red" : "\\[\\e[31m\\]",
		"\\green" : "\\[\\e[32m\\]",
		"\\brown" : "\\[\\e[33m\\]",
		"\\blue" : "\\[\\e[34m\\]",
		"\\purple" : "\\[\\e[35m\\]",
		"\\cyan" : "\\[\\e[36m\\]",
		"\\white" : "\\[\\e[37m\\]",
		"\\bold" : "\\[\\e[1m\\]",
		"\\dim" : "\\[\\e[2m\\]",
		"\\italic" : "\\[\\e[3m\\]",
		"\\underline" : "\\[\\e[4m\\]",
		"\\blink" : "\\[\\e[5m\\]",
		"\\reverse" : "\\[\\e[7m\\]",
		"\\overline" : "\\[\\e[53m\\]",
		"\\clear" : "\\[\\e[m\\]",
	}

	code := "" // the input string

	if (len(os.Args) > 1) {
		// read input string as the first command line argument
		code = os.Args[1]
	} else {
		// no command line arguments, read input from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			code = scanner.Text()
		}
	}

	// Replace background commands
	for key, value := range commands_bg { 
		code = strings.Replace(code, key, value, -1)
	}

	// Replace other commands
	for key, value := range commands { 
		code = strings.Replace(code, key, value, -1)
	}

	// Remove non-command spaces
	code = strings.Replace(code, " ", "", -1)

	// Replace space commands
	code = strings.Replace(code, "\\S", " ", -1)
	code = strings.Replace(code, "\\space", " ", -1)

	fmt.Println(code)
}