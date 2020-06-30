package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("whoami")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(out)
		os.Exit(-1)
	}
	fmt.Println(string(out))
}
