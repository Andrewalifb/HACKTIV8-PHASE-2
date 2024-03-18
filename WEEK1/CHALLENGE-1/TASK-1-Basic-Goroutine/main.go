package main

import (
	"fmt"
	"time"
)

func printNumbers(num int) {
		fmt.Println(num)
}

func printLetters(letter string) {
		fmt.Println(letter)
}
func main() {

	numbers := 10
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 1; i <= numbers; i++ {

		go printNumbers(i)
	}
	
	for i := 0; i < len(letter); i++ {

		go printLetters(letter[i])
	}
	
	time.Sleep(1 * time.Second)

}