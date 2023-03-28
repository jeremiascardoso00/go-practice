package main

import (
	"fmt"
)

func main() {
	// var c1 chan float64

	// // Declaring and initilizing a RECEIVE-ONLY channel
	// c2 := make(<-chan rune)

	// // Declaring and initilizing a SEND-ONLY channel
	// c3 := make(chan<- rune)

	// c4 := make(chan int, 10)

	// fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)

	//

	// ch := make(chan string)

	// go func(n string) {
	// 	ch <- n
	// }("Gophers!")

	// value := <-ch
	// fmt.Println("Value received from channel:", value)

	//

	// c := make(chan int)

	// go func(n int) {
	// 	c <- n
	// }(100)

	// fmt.Println(<-c)

	//

	c := make(chan int, 50)

	for i := 1; i <= 50; i++ {
		go func(n int, c chan int) {
			c <- n * n
		}(i, c)
		result := <-c
		fmt.Println(result)
	}

}

// func power(n int, c chan int) {
// 	c <- n * n
// }
