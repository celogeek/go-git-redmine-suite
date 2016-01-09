package main

import (
    "fmt"
)

const VERSION = 0.01
const AUTHOR = "Celogeek"

func main() {
    fmt.Printf("Git Redmine Suite v%v by %v\n", VERSION, AUTHOR)
    fmt.Println("")
    fmt.Println("[task]")
    fmt.Println("  * git task info [TASK_ID]: get task information")
}
