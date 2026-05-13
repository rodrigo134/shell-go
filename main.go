package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	builtins := map[string]func([]string){}
	builtins = map[string]func([]string){

		"echo": func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
		"cd": func(args []string) {
			if len(args) == 0 {
				fmt.Println("Usage: cd [Directory]")
				return
			}
			path := args[0]
			if path == "~" {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Println("error home directory")
					return
				}
				path = home
				

				}
				err := os.Chdir(path)
				if err != nil {
					fmt.Printf("cd: %s: No such file or directory\n", args[0])
			}
		
		},

		"pwd": func(args []string) {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory")
				return
			}
			fmt.Println(dir)

		},
		"exit": func(args []string) {
			os.Exit(0)
		},
		"type": func(args []string) {

			if len(args) == 0 {
				return
			}
			command := args[0]
			if _, ok := builtins[command]; ok {
				fmt.Printf("%s is a shell builtin\n", args[0])
			} else {
				fmt.Printf("%s: not found\n", command)
			}

		},
		"clear": func(args []string) {
			cmd :=exec.Command("clear")
			cmd.Stdout = os.Stdout
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

		parts := strings.Fields(prompt)
		commands := parts[0]
		args := parts[1:]

		if fn, ok := builtins[commands]; ok {
			fn(args)
		} else {
			_, err := exec.LookPath(commands) //path = /usr/bin/ls
			if err != nil {
				fmt.Println(commands + ": not found")
				continue

			} else {
				externalCmd := exec.Command(commands, args...) ///usr/bin/ls -la
				externalCmd.Stdin = os.Stdin
				externalCmd.Stdout = os.Stdout
				externalCmd.Stderr = os.Stderr
				externalCmd.Run()
				if err != nil {
					fmt.Println("Error executing", externalCmd.Args[0]+":", err)
				}

			}

		}
	}
}
