package main

import (
	"bufio"
	"container/list"
	"io"
	"os"
	"os/exec"
	"strings"
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
		currentCommandArray := strings.Split(currentCommand.Value.(string), " ")

		cmd := exec.Command("pkill", currentCommandArray[0], currentCommandArray[1])
		cmdErr := cmd.Run()
		check(cmdErr)

		currentCommand = currentCommand.Next()
	}

	f.Close()
}
