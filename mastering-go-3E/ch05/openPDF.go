package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("evince",os.Getenv("TEST"))
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}