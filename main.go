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

func startTimer(ticker time.Ticker, timer int, stop chan bool) {

	i := timer
	for range ticker.C {
		if i != 0 {
			select {
			case <-stop:
				return
			default:
				i--
				fmt.Printf("\r\033[A\033[K%d seconds left\n", i)
			}
		} else {
			usage()
			return
		}
	}
}

func main() {

	usage()

	var input int
	var ticker time.Ticker
	stop := make(chan bool)
	
	for {
		_, err := fmt.Scanf("%d", &input)

		if err != nil {
			fmt.Println(err)
		}
		
		switch input {
		case 1:
			ticker.Stop()
			ticker = *time.NewTicker(time.Second)
			go startTimer(ticker, 4, stop)
		case 2:
			stop <- true
			usage()
		}
	}
}
