package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please Enter Your Command:")

	var cmd string = ""
	for true {
		fmt.Print("[Menu] ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			fmt.Println("Current avalid commmand: [help]|[version]|[quit]")
		case "version":
			fmt.Println("Menu Version: v0.01")
		case "quit":
			fmt.Println("Exit sucessfully")
			return
		default:
			fmt.Println("Wrong command! Enter \"help\" for command list")
		}
	}
}
