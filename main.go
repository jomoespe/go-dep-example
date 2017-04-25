package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("go-dep-example> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil || line == "exit" {
			break
		}
		fmt.Printf("received command: %s\n", line)
	}
}
