package main

import (
	"fmt"
	"time"
)

func usage() {
	fmt.Println("[1] Start")
	fmt.Println("[2] Stop")
	fmt.Println("[3] Change timer")
	fmt.Println("[4] Exit")
}

func startTimer(ticker time.Ticker, timer [2]int, stop chan bool) {

	i := timer[0]
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
			// Swap timer
			timer = func() [2]int {
				return [2]int{timer[1], timer[0]}
			}()

			go startTimer(ticker, timer, stop)
			return
		}
	}
}

func main() {

	usage()

	var input int
	var ticker time.Ticker

	timer := [2]int{25, 5}

	stop := make(chan bool)

	for {
		_, err := fmt.Scanf("%d", &input)

		if err != nil {
			fmt.Println(err)
		}

		switch input {
		case 1:
			ticker = *time.NewTicker(time.Second)
			go startTimer(ticker, timer, stop)
		case 2:
			stop <- true
			ticker.Stop()
			fmt.Printf("\r\033[A\033[KTimer stopped\n")
		case 4:
			return
		}
	}
}
