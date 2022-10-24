package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}
	if err := handleCmd(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func handleCmd(args []string) error {
	// ...

	return nil
}
