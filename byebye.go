package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("/home/multiojuice/.byebyerc")
	check(err)

	configReader := bufio.NewReader(f)

	line, isPrefix, readErr := configReader.ReadLine()
	check(readErr)

	for line != nil {
		fmt.Printf("%s\n", string(line))
		line, isPrefix, readErr = configReader.ReadLine()

		if isPrefix || readErr == io.EOF {
			break
		}
	}

	// cmd := exec.Command("pkill", "-SIGINT", "tmux")
	// log.Printf("Bye bye!")
	// cmdErr := cmd.Run()
	// log.Printf("Command finished with error: %v", cmdErr)

	f.Close()
}
