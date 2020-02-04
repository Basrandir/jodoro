package main

import (
	"fmt"
	"time"
)

func usage() {
	fmt.Println("[1] Start")
	fmt.Println("[2] Stop")
	fmt.Println("[3] Change timer")
}

func startTimer() {
	ticker := time.NewTicker(time.Second)

	i:= 15
	
	for range ticker.C {
		i--
		fmt.Printf("\r\033[K%d seconds left", i)
	}
}

func main() {
	usage()

	var input int
	_, err := fmt.Scanf("%d", &input)

	if err != nil {
		fmt.Println(err)
	}

	switch input {
	case 1:
		go startTimer()
		time.Sleep(time.Second * 15)
		fmt.Println("")
	}
}
