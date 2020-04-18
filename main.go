package main

import (
	"bufio"
	"counter-queue/log"
	"counter-queue/queue"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of counter: ")
	text, _ := reader.ReadString('\n')
	txt := strings.TrimSpace(text)

	num, err := strconv.Atoi(txt)
	if err != nil {
		log.Stdout("failed read number of counter: ", err)
		os.Exit(1)
	}

	queue.ProcessQueue(num)
}
