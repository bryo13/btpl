package main

import (
	"fmt"
	"ink/repl"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Hello %s this is Ink\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
