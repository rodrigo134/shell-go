package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	builtins := map[string]func([]string){}
	builtins = map[string]func([]string){

		"echo": func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
		"exit": func(args []string) {
			os.Exit(0)
		},
		"type": func(args []string) {
			
			if len(args)==0 {
				return
			}
			command := args[0]
			if _,ok :=builtins[command];ok{
				fmt.Printf("%s is a shell builtin\n",args[0])
			}else{
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
		cmd := parts[0]
		args := parts[1:]

		if fn, ok := builtins[cmd]; ok {
			fn(args)
		} else {
			fmt.Println(cmd, ":Not found")
		}

	}

}
