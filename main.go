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

func startTimer(ticker time.Ticker, timer [4]int, stop chan bool) {

	i := timer[0]
	
	timerType := "Work"

	if timer[3] == 7 {
		timerType = "Long Break"
	} else if timer[3] % 2 == 1 {
		timerType = "Break"
	}
	
	for range ticker.C {
		if i != 0 {
			select {
			case <-stop:
				return
			default:
				i--
				fmt.Printf("\r\033[A\033[K%s: %d seconds left\n", timerType, i)
			}
		} else {
			timer[3]++

			if timer[3] < 7 {
				// Swap work/break timer
				timer = func() [4]int {
					return [4]int{timer[1], timer[0], timer[2], timer[3]}
				}()
			} else {
				// Swap work/long break timer
				timer = func() [4]int {
					return [4]int{timer[2], timer[1], timer[0], timer[3]}
				}()

				// Reset countdown after long break
				if timer[3] == 8 {
					timer[3] = 0
				}
			}

			go startTimer(ticker, timer, stop)
			return
		}
	}
}

func main() {

	usage()

	var input int
	var ticker time.Ticker

	timer := [4]int{25, 5, 30, 0}

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
