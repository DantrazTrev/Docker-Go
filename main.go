package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker run <image> <cmd> <params>
// go run main.go <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("idl")
	}
}
func run() {
	fmt.Printf("Running %v ", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
