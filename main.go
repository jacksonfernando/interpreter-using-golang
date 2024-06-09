package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/v2/golang-intrepeter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the monkey programming language!\n", user.Username)
	fmt.Printf("Feel feree to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
