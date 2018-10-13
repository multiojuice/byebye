package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSignalFromString(configSignal string) string {
	switch configSignal {
	case "Hangup":
		fallthrough
	case "hangup":
		return "-SIGHUP"
	case "Interupt":
		fallthrough
	case "interrupt":
		return "-SIGINT"
	case "Terminate":
		fallthrough
	case "terminate":
		return "-SIGTERM"
	case "Kill":
		fallthrough
	case "kill":
		return "-SIGKILL"
	}
	return "-SIGTERM"
}

func byebye(isAll bool) {
	currentUser, userErr := user.Current()
	check(userErr)

	f, err := os.Open(currentUser.HomeDir + "/.byebyerc")
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

	if isAll {
		for currentCommand != nil {
			currentCommandArray := strings.Split(currentCommand.Value.(string), " ")

			cmd := exec.Command("pkill", getSignalFromString(currentCommandArray[0]), currentCommandArray[1])
			cmdErr := cmd.Run()
			check(cmdErr)

			currentCommand = currentCommand.Next()
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for currentCommand != nil {
			currentCommandArray := strings.Split(currentCommand.Value.(string), " ")

			fmt.Println("Enter 'y' if you would like to end: ", currentCommandArray[1])
			scanner.Scan()
			if scanner.Text() == "y" || scanner.Text() == "yes" {
				cmd := exec.Command("pkill", getSignalFromString(currentCommandArray[0]), currentCommandArray[1])
				cmdErr := cmd.Run()
				check(cmdErr)
			}
			currentCommand = currentCommand.Next()
		}
	}
	f.Close()
}

func main() {
	if len(os.Args) < 2 {
		byebye(true)
	} else {
		switch os.Args[1] {
		case "help":
			fmt.Println("\tbyebye help")
			fmt.Println("\t\tSummary: byebye is a tool that will kill, interupt, or shutdown your processes.")
			fmt.Println("\t\tSubcommands:")
			fmt.Println("\t\t\t'help' -> well you are doing it...")
			fmt.Println("\t\t\t'some' -> this will go through each process you have listed and ask if you would like to byebye it")
			fmt.Println("\t\t\tno subcommands will byebye each process in your .byebyerc")
			fmt.Println("\t\tDirections:")
			fmt.Println("\t\t\tYou must first configure your .byebyerc into your homedir:")
			fmt.Println("\t\t\tEach line in your .byebyerc represents a process and the way you want to kill it. It will take two arguments that are space seperated, first the signal you want to send it, second the process name you want to kill. ")
			fmt.Println("\t\t\tSignal types are:")
			fmt.Println("\t\t\t\t'hangup' -> sends a SIGHUP or 1 to the program, this probably won't do anything.")
			fmt.Println("\t\t\t\t'interrupt' -> sends a SIGINT or 2 to the program, this will 'gracefully' shut down most apps")
			fmt.Println("\t\t\t\t'terminate' -> sends a SIGTERM or 15 signal, this will be a bit more forceful than an interupt. the same signal send by the kill command")
			fmt.Println("\t\t\t\t'kill' -> sends a SIGKILL or a 9 to the process, this will completely and always shut a process down, very forceful.")
			fmt.Println("\t\tExample config:")
			fmt.Println("\t\t\tinterrupt code")
			fmt.Println("\t\t\tinterrupt chrome")
			fmt.Println("\t\t\tterminate tmux")
			fmt.Println("\t\tTips:")
			fmt.Println("\t\t\t- The order of your process config matters! If you kill your tmux session or your terminal instance first, nothing else will be executed, do those last")
			fmt.Println("\t\t\t- If byebye does't end a process how you want it to, try changing the configuration and signal that you send to that process, try to adjust the force of the signal (described above) to fit your use case.")

		case "some":
			byebye(false)

		default:
			fmt.Println("Unknown subcommand, try `byebye help`")
		}
	}
}
