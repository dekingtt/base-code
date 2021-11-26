package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOutput, err := dateCmd.Output()
	if nil != err {
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOutput))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep\n"))
	grepIn.Close()
	grepByte, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(">hello grep")
	fmt.Println(string(grepByte))

	lsCmd := exec.Command("ls", "-a", "-l", "-h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(lsOut))
}
