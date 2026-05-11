package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	
	reader := bufio.NewReader(os.Stdin)

	var builtins map[string]func([]string)

	builtins = map[string]func([]string){

		"echo": func(s []string) {		},
		"exit":func(s []string) {
			os.Exit(0)
		},
		"type":func(s []string) {},

	}

	for{
		fmt.Print("$ ")

		input,err := reader.ReadString('\n')

		if err != nil{
			fmt.Print("Error input ",err)
		}

		prompt := strings.TrimSpace(input)

		if prompt == "" {
			continue
		}
		
		parts :=strings.Fields(prompt)
		cmd := parts[0]
		args :=parts[1:]

		if fn,ok := builtins[cmd];ok{
			fn(args)
		}


	}





}