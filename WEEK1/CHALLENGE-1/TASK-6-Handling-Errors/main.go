package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

func sendChannel(num int, chanNum chan int, wg *sync.WaitGroup) {
	wg.Done()
	chanNum <- num
}
func checkGreaterThan20(num int) (bool, error) {
	if num > 20 {
		return true, errors.New("Error: Number " + strconv.Itoa(num) + " is greater than 20")
	}
	return false, nil
}
func main() {
	var wg sync.WaitGroup
	var receivedNumbers []string
	evenNumber := make(chan int)
	oddNumber := make(chan int)

  for i := 1; i <= 22; i++ {
		
		isGreater, errVal := checkGreaterThan20(i)
		if isGreater {
			receivedNumbers = append(receivedNumbers, errVal.Error())
		} else {
			wg.Add(1)
			if i % 2 == 0 {
				go sendChannel(i, evenNumber, &wg)
				receivedNumbers = append(receivedNumbers, fmt.Sprintf("Received an Even Number : %d", <-evenNumber))
			} else {
				go sendChannel(i, oddNumber, &wg)
				receivedNumbers = append(receivedNumbers, fmt.Sprintf("Received an Odd Number : %d", <-oddNumber))
			}
		}

	}

  close(evenNumber) 
	close(oddNumber)

  for _, num := range receivedNumbers {
    fmt.Println(num)
  }
	wg.Wait()
}