package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(channel chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter message:")
	text, _ := reader.ReadString('\n')
	channel <- text
}

func output(text string) {
	fmt.Printf("The output: %s", text)
}

func main() {
	var done chan struct{} = make(chan struct{})
	var channel chan string = make(chan string)
	go input(channel)
	go func() {
		for {
			select {
			case text := <-channel:
				output(text)
				done <- struct{}{}
			}
		}
	}()
	<-done
}
