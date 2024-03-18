package main

import (
	"fmt"
	"sync"
)

func printNumbers(num int, wg *sync.WaitGroup) {
	defer wg.Done()
		fmt.Println(num)
}

func printLetters(letter string, wg *sync.WaitGroup) {
	defer wg.Done()	
		fmt.Println(letter)
}
func main() {
	var wg sync.WaitGroup
	numbers := 10
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 1; i <= numbers; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}
	
	for i := 0; i < len(letter); i++ {
		wg.Add(1)
		go printLetters(letter[i], &wg)
	}
	
	wg.Wait()
}