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
			path, err := exec.LookPath(commands) //path = /usr/bin/ls
			if err != nil {
				fmt.Println(commands + ": not found")
				continue

			}
			cmd := exec.Command(path, args...) ///usr/bin/ls -la
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

		}
	}
}
