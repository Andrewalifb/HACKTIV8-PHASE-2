package main

import (
	"fmt"
	"sync"
)

func sendChannel(num int, chanNum chan int, wg *sync.WaitGroup) {
	wg.Done()
	chanNum <- num
}
func main() {
	var wg sync.WaitGroup
	var receivedNumbers []string
	evenNumber := make(chan int)
	oddNumber := make(chan int)

  for i := 1; i <= 20; i++ {
		wg.Add(1)
		if i % 2 == 0 {
			go sendChannel(i, evenNumber, &wg)
			receivedNumbers = append(receivedNumbers, fmt.Sprintf("Received an Even Number : %d", <-evenNumber))
		} else {
			go sendChannel(i, oddNumber, &wg)
			receivedNumbers = append(receivedNumbers, fmt.Sprintf("Received an Odd Number : %d", <-oddNumber))
		}
	}

  close(evenNumber) 
	close(oddNumber)

  for _, num := range receivedNumbers {
    fmt.Println(num)
  }
	wg.Wait()
}