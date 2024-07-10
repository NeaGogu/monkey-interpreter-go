package main

import (
	"NeaGogu/monkey-interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)
	fmt.Printf("Hello %s! This is Monke!\n", user.Username)
	fmt.Printf("Type mf\n")
	repl.Start(os.Stdin, os.Stdout)
}
