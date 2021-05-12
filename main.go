package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//docker run <container> cmd args
//go run main.go run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Only run command is supported! :( ")
	}
}

func run() {
	//fork and exec a new child process for running the command
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//New namespaces for the process: NEWUTS(hostname and the NIS domain name), PID(process id)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	errorCatcher(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//Need to mount the linux fs before calling the below statements
	errorCatcher(syscall.Chroot("/mnt/containerFS"))
	errorCatcher(os.Chdir("/"))
	errorCatcher(syscall.Mount("proc", "proc", "proc", 0, ""))
	errorCatcher(cmd.Run())
	errorCatcher(syscall.Unmount("proc", 0))

}

func errorCatcher(err error) {
	if err != nil {
		panic(err)
	}
}
