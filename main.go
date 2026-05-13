package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	builtins := map[string]func(io.Writer, []string){}
	builtins = map[string]func(io.Writer, []string){

		"echo": func(w io.Writer, args []string) {
			fmt.Fprintln(w, strings.Join(args, " "))
		},
		"cd": func(w io.Writer, args []string) {
			if len(args) == 0 {
				fmt.Fprintln(w, "Usage: cd [Directory]")
				return
			}
			path := args[0]
			if path == "~" {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintln(w, "error home directory")
					return
				}
				path = home

			}
			err := os.Chdir(path)
			if err != nil {
				fmt.Fprintf(w, "cd: %s: No such file or directory\n", args[0])
			}

		},

		"pwd": func(w io.Writer, args []string) {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(w, "Error getting current directory")
				return
			}
			fmt.Fprintln(w, dir)

		},
		"exit": func(w io.Writer, args []string) {
			os.Exit(0)
		},
		"type": func(w io.Writer, args []string) {

			if len(args) == 0 {
				return
			}
			command := args[0]
			if _, ok := builtins[command]; ok {
				fmt.Fprintf(w, "%s is a shell builtin\n", args[0])
				return
			}

			path, err := exec.LookPath(command)
			if err == nil {
				fmt.Fprintf(w, "%s is %s\n", command, path)
				return
			}
			fmt.Fprintf(w, "%s: not found\n", command)

		},
		"clear": func(w io.Writer, args []string) {
			cmd := exec.Command("clear")
			cmd.Stdout = w
			cmd.Run()
		},
	}

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Print("Error input ", err)
		}

		prompt := strings.TrimSpace(input)

		if prompt == "" {
			continue
		}

		parts := parseCommand(prompt)
		commands := parts[0]
		args := parts[1:]

		// Check for redirection and extract command/args without redirect operators
		command, outputFileName, hasRedirect := parseRedirection(parts)
		if hasRedirect && len(command) > 0 {
			commands = command[0]
			args = command[1:]
		}

		var outputFile *os.File
		if hasRedirect && outputFileName != "" {
			var err error
			outputFile, err = os.Create(outputFileName)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
			// Close file immediately after command execution
		}

		if fn, ok := builtins[commands]; ok {
			if outputFile != nil {
				fn(outputFile, args)
			} else {
				fn(os.Stdout, args)
			}
		} else {
			_, err := exec.LookPath(commands) //path = /usr/bin/ls
			if err != nil {
				fmt.Println(commands + ": not found")
				continue

			} else {
				externalCmd := exec.Command(commands, args...) ///usr/bin/ls -la
				externalCmd.Stdin = os.Stdin
				if outputFile != nil {
					externalCmd.Stdout = outputFile
				} else {
					externalCmd.Stdout = os.Stdout
				}
				externalCmd.Stderr = os.Stderr
				if err := externalCmd.Run(); err != nil {
					fmt.Println("Error executing", externalCmd.Args[0]+":", err)
				}

			}

		}

		// Close output file if it was opened
		if outputFile != nil {
			outputFile.Close()
		}
	}
}

func parseCommand(input string) []string {
	inQuotes := false
	var quoteChar byte

	commandFinal := []string{}
	current := ""

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch == '\\' && !inQuotes {
			if i+1 < len(input) {
				current += string(input[i+1])
				i++
			} else {
				current += string(ch)
			}
			continue
		}
		if ch == '\\' && inQuotes && quoteChar == '"' {
			if i+1 < len(input) {
				next := input[i+1]

				if next == '"' || next == '\\' {
					current += string(next)
					i++
					continue
				}
			}

			current += string(ch)
			continue
		}

		if ch == '"' || ch == '\'' {
			if !inQuotes {
				inQuotes = true
				quoteChar = ch
				continue
			}

			if ch == quoteChar {
				inQuotes = false
				quoteChar = 0
				continue
			}

			current += string(ch)
			continue
		}

		if ch == ' ' && !inQuotes {
			if current != "" {
				commandFinal = append(commandFinal, current)
				current = ""
			}
			continue
		}

		current += string(ch)
	}

	if current != "" {
		commandFinal = append(commandFinal, current)
	}

	return commandFinal
}

// parseRedirection checks for input and output redirection in the command.
// [echo hello > output.txt] => command: "echo hello", outputFile: "output.txt", hasRedirect: true
func parseRedirection(input []string) ([]string, string, bool) {
	for i := 0; i < len(input); i++ {
		token := input[i]
		if token == ">" || token == "1>" {
			if i+1 >= len(input) {
				break
			}
			return input[:i], input[i+1], true
		}
	}

	return input, "", false
}
