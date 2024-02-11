package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

var commands = map[string]string{
	"\\blackbg":   "\\[\\e[40m\\]",
	"\\redbg":     "\\[\\e[41m\\]",
	"\\greenbg":   "\\[\\e[42m\\]",
	"\\brownbg":   "\\[\\e[43m\\]",
	"\\bluebg":    "\\[\\e[44m\\]",
	"\\purplebg":  "\\[\\e[45m\\]",
	"\\cyanbg":    "\\[\\e[46m\\]",
	"\\whitebg":   "\\[\\e[47m\\]",
	"\\black":     "\\[\\e[30m\\]",
	"\\red":       "\\[\\e[31m\\]",
	"\\green":     "\\[\\e[32m\\]",
	"\\brown":     "\\[\\e[33m\\]",
	"\\blue":      "\\[\\e[34m\\]",
	"\\purple":    "\\[\\e[35m\\]",
	"\\cyan":      "\\[\\e[36m\\]",
	"\\white":     "\\[\\e[37m\\]",
	"\\bold":      "\\[\\e[1m\\]",
	"\\dim":       "\\[\\e[2m\\]",
	"\\italic":    "\\[\\e[3m\\]",
	"\\underline": "\\[\\e[4m\\]",
	"\\blink":     "\\[\\e[5m\\]",
	"\\reverse":   "\\[\\e[7m\\]",
	"\\overline":  "\\[\\e[53m\\]",
}

var special_attrs = map[string]string{
	"\\S":     " ",
	"\\space": " ",
}

// TODO add \D{format}
var default_attrs = []string{
	"\\u",
	"\\h",
	"\\a",
	"\\d",
	"\\H",
	"\\j",
	"\\l",
	"\\n",
	"\\r",
	"\\s",
	"\\t",
	"\\T",
	"\\@",
	"\\A",
	"\\v",
	"\\V",
	"\\w",
	"\\W",
	"\\!",
	"\\#",
	"\\$",
	"\\nnn",
}

const (
	SCAN int16 = iota
	WAIT_CMD_CHAR
	READ_CMD
	WAIT_OPENING_BRACE
)

type Stack []string

var termination_chars = []string{"\\", "{", "}", " "}

var state = SCAN

func get_keys(my_map map[string]string) []string {
	keys := make([]string, len(my_map))
	i := 0
	for key := range my_map {
		keys[i] = key
		i++
	}
	return keys
}

func contains(my_slice []string, search string) bool {
	for i := 0; i < len(my_slice); i++ {
		if search == my_slice[i] {
			return true
		}
	}
	return false
}

func add_clear_to_output(stack Stack, output string) string {
	output += "\\[\\e[0m\\]"
	for _, attr := range stack {
		output += commands[attr]
	}
	return output
}

func push(my_stack Stack, my_attr string) Stack {
	my_stack = append(my_stack, my_attr)
	return my_stack
}

func pop(my_stack Stack) Stack {
	length := len(my_stack)
	if length < 1 {
		fmt.Println("Invalid syntax")
		os.Exit(1)
	}
	return append(my_stack[:length-1])
}

func is_empty(my_stack Stack) bool {
	return (len(my_stack) == 0)
}

func get_input_code() string {
	code := ""
	if (len(os.Args) > 1) {
		// Read input string as the first command line argument
		code = strings.Join(os.Args[1:], " ")
	} else {
		// No command line arguments, read input from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			code = scanner.Text()
		}
	}
	return code
}

func main() {
	var stack Stack
	
	code := get_input_code()
	
	cmd_keys := get_keys(commands)
	special_attrs_keys := get_keys(special_attrs)
	cmd_buffer := "\\" // must always be reset to "\\" and not an empty string
	output := ""

	for i := 0; i < len(code); i++ {
		cur_char := string(code[i])

		//fmt.Printf("Index: %d, State: %d, Char: %s, Cmd_Buffer: %s\n", i, state, cur_char, cmd_buffer)
		//fmt.Println(stack)

		switch state {
		case SCAN:
			// Transition condition
			if cur_char == "\\" {
				state = WAIT_CMD_CHAR
				break
			}

			// Cancel an attribute
			if cur_char == "}" {
				stack = pop(stack)
				output = add_clear_to_output(stack, output)
			}

			// State operations
			if cur_char != " " && cur_char != "}" {
				output += cur_char
			}

		case WAIT_CMD_CHAR: // just scan until a letter
			if cur_char != " " && cur_char != "\\" {
				cmd_buffer += cur_char
				state = READ_CMD
			}

			if cur_char == "\\" {
				// This means that the user wants to add the actual character \ into the variable
				output += "\\\\"
				cmd_buffer = "\\"
				state = SCAN
			}

		case READ_CMD: // as long as a termination char is not reached, save cmd
			if contains(termination_chars, cur_char) {
				// This state is over, decide which path to take
				if cur_char == "}" {
					// This path is taken when a default attribute is being read
					// and immediately afterwards a command needs to be cancelled
					if contains(default_attrs, cmd_buffer) {
						output += cmd_buffer
						stack = pop(stack)
						output = add_clear_to_output(stack, output)
						cmd_buffer = "\\"
						state = SCAN
						break

					} else if contains(special_attrs_keys, cmd_buffer) {
						output += special_attrs[cmd_buffer]
						stack = pop(stack)
						output = add_clear_to_output(stack, output)
						cmd_buffer = "\\"
						state = SCAN
						break

					} else {
						fmt.Println("Invalid syntax")
						os.Exit(1)
					}

				} else if contains(cmd_keys, cmd_buffer) {
					// This path is taken when a valid command is detected
					output += commands[cmd_buffer]
					stack = push(stack, cmd_buffer)
					cmd_buffer = "\\"

					if (cur_char == "{") {
						state = SCAN
					} else {
						state = WAIT_OPENING_BRACE
					}
					break

				} else if contains(default_attrs, cmd_buffer) {
					// This path is taken when the buffered item is not a command but a default attribute
					output += cmd_buffer
					cmd_buffer = "\\"
					state = SCAN
					break

				} else if contains(special_attrs_keys, cmd_buffer) {
					// This path is taken when the buffered item is a special command (e.g. a space)
					output += special_attrs[cmd_buffer]
					cmd_buffer = "\\"
					state = SCAN
					break
					
				} else {
					fmt.Printf("Invalid command: %s\n", cmd_buffer)
					os.Exit(1)
				}
			}
			cmd_buffer += cur_char // if the current char is not a termination command

		// If the opening brace is not immediately after the command
		case WAIT_OPENING_BRACE:
			if (cur_char == "{") {
				state = SCAN
			} else if (cur_char != " ") {
				fmt.Println("Invalid syntax")
				os.Exit(1)
			}

		default:
			fmt.Println("Impossible state!")
			os.Exit(1)
		}
	}

	if is_empty(stack) == false {
		fmt.Println("WARNING: Unterminated attribute")
	}
	fmt.Println(output)
}
