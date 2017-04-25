// A shell-style sample application
package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

const (
	prompt   = "go-dep-example> "
	exit_cmd = "exit"
)

func main() {
	rl, err := readline.New(prompt)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil || line == exit_cmd {
			break
		}
		fmt.Printf("received command: %s\n", line)
	}
}
