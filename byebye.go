package main

import (
	"bufio"
	"container/list"
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

	configList := list.New()

	configReader := bufio.NewReader(f)

	line, isPrefix, readErr := configReader.ReadLine()
	check(readErr)

	for line != nil {
		configList.PushBack(string(line))
		line, isPrefix, readErr = configReader.ReadLine()

		if isPrefix || readErr == io.EOF {
			break
		}
	}

	currentCommand := configList.Front()

	for currentCommand != nil {
		fmt.Println(currentCommand.Value)
		currentCommand = currentCommand.Next()
	}

	// cmd := exec.Command("pkill", "-SIGINT", "tmux")
	// log.Printf("Bye bye!")
	// cmdErr := cmd.Run()
	// log.Printf("Command finished with error: %v", cmdErr)

	f.Close()
}
