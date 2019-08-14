package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	channel <- "dam"
	channel <- "van dam"
	channel <- "cloud van dam"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
