package main

import (
	"fmt"
	"sync"
)

func printNumbers(num int, nums chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	nums <- num
}


func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)
	numbers := 10

	for i := 1; i <= numbers; i++ {
		wg.Add(1)
		go printNumbers(i, numberChan,&wg)
	}
	
  var receivedNumbers []int
  for i := 0; i < numbers; i++ {
    receivedNumbers = append(receivedNumbers, <-numberChan)
  }

  close(numberChan) 

  for _, num := range receivedNumbers {
    fmt.Println(num)
  }
	wg.Wait()

}