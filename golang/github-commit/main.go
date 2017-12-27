package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	var commandsOne []string
	var commandsTwo []string
	var commandsThree []string
	commandsOne = append(commandsOne, "git")
	commandsTwo = append(commandsTwo, "clone")
	commandsThree = append(commandsThree, "https://github.com/julioc98/day-off-day.git")

	defer fmt.Println("\n\rDone All. \n\r")
	for i := 0; i < len(commandsOne); i++ {

		fmt.Println("\n\rStart\n\r")
		out, err := exec.Command(commandsOne[i], commandsTwo[i], commandsThree[i]).CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
	fmt.Println("\n\rDone. \n\r")

	arq, err := os.Open("./day-off-day/README.md")
	if err != nil {
		panic(err)
	}
	defer arq.Close()

	var lines []string
	scanner := bufio.NewScanner(arq)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(lines)
}
