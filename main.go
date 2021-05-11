package main

import (
	"fmt"
	"os"
	"os/exec"
)

//docker run <container> cmd args
//go run main.go run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Only run command is supported! :( ")
	}
}

func run() {
	fmt.Printf("running %v \n", os.Args[2:])
	//Run commands, no isolation
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
