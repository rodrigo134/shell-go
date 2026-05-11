package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	
	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("$ ")

		_,err := reader.ReadString('\n')

		if err != nil{
			fmt.Print("Error input ",err)
		}

	}





}