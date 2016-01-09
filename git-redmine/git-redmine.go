package main

import (
	"fmt"
	"os"
)

const VERSION = 0.01
const AUTHOR = "Celogeek"

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	command := os.Args[1]
	var params []string
	if len(os.Args) > 2 {
		params = os.Args[2:len(os.Args)]
	}

	switch command {
	case "task":
		task(params)
	default:
		help()
	}
}

func info() {
	fmt.Printf("Git Redmine Suite v%v by %v\n", VERSION, AUTHOR)
	fmt.Println("")
}

func help() {
	info()
	task_help()
}
