package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("pkill", "tmux")
	log.Printf("Bye bye!")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
