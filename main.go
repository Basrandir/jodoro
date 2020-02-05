package main

import (
	"fmt"
	"time"
)

func usage() {
	fmt.Println("[1] Start")
	fmt.Println("[2] Stop")
	fmt.Println("[3] Pause")
	fmt.Println("[4] Change timer")
}

func startTimer(ticker *time.Ticker, timer int, stop chan bool) {

	i := timer
	for range ticker.C {
		select {
		case <-stop:
			return
		default:
			i--
			fmt.Printf("\r\033[K%d seconds left", i)
		}
	}
}

func main() {
	usage()

	var input int
	_, err := fmt.Scanf("%d", &input)

	if err != nil {
		fmt.Println(err)
	}

	for {
		ticker := time.NewTicker(time.Second)
		stop := make(chan bool)
		
		switch input {
		case 1:
			go startTimer(ticker, 25, stop)
			time.Sleep(time.Second * 25)
			ticker.Stop()
			usage()
		case 2:
			
		}
	}
}
