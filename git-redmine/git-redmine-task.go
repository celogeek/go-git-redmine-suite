package main

import (
	"fmt"
)

func task(params []string) {
	if len(params) < 2 {
		info()
		task_help()
		return
	}

	command := params[0]
	params = params[1:len(params)]

	switch command {
	case "info":
		task_info(params)
	default:
		info()
		task_help()
	}
}

func task_help() {
	fmt.Println("[task]")
	fmt.Println("  * git task info [TASK_ID]: get task information")
}

func task_info(params []string) {
	info()
	fmt.Printf("Gathering information on task %v\n", params[0])
}
