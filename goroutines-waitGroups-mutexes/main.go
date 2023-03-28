package main

import (
	"fmt"
	"sync"
)

/*
Coding Exercise #1

There is an error in the following Go Program. Even though the goroutine is correctly launched, it doesn't print any message.

 package main

import (
    "fmt"
)

func sayHello(n string) {
    fmt.Printf("Hello, %s!\n", n)
}

func main() {
    go sayHello("Mr. Wick")
}
Your task is to synchronize main and the goroutine using WaitGroups. The program should print the string received as argument by sayHello().


Coding Exercise #2

1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.

2. From main launch 3 goroutines that execute the function you have just created (sum)

3. Synchronize the goroutines and the main function using WaitGroups


Coding Exercise #3

1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.

2. Launch the function as a goroutine and synchronize it with main using WaitGroups

Note: You calculate the square root of a float named f using the Sqrt() function from math package like this:

x := math.Sqrt(f)
fmt.Println(x)


Coding Exercise #4

Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included).


Coding Exercise #5

You work at a Banking Application and have created 2 functions: one that deposits a value into an account and another that withdraws a value from the account.

You want to simulate many deposits and withdraws that take place simultaneously  and start some goroutines.

During testing you notice that a date race occurred.

Your task is to change the code in order to protect the account's balance from simultaneously writing using a mutex.

This is the initial program that has errors:

package main

import (
    "fmt"
    "sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup) {
    *b += n
    wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup) {
    *b -= n
    wg.Done()
}

func main() {
    var wg sync.WaitGroup

    wg.Add(200)

    balance := 100

    for i := 0; i < 100; i++ {
        go deposit(&balance, i, &wg)
        go withdraw(&balance, i, &wg)
    }
    wg.Wait()

    fmt.Println("Final balance value:", balance)
}
*/

func sayHello(wg *sync.WaitGroup, n string) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(a, b float64, wg *sync.WaitGroup) {

	s := a + b
	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", a, b, s)
	wg.Done()

}

func main() {

	// var wg sync.WaitGroup

	// wg.Add(1)

	// go sayHello(&wg, "Mr. Wick")

	// wg.Wait()

	//

	// var wg sync.WaitGroup

	// wg.Add(3)

	// sum(1, 2, &wg)
	// sum(2, 2, &wg)
	// sum(3, 2, &wg)

	// wg.Wait()

	//

	// var wg sync.WaitGroup

	// wg.Add(1)

	// go func(value float64, wg *sync.WaitGroup) {
	// 	x := math.Sqrt(value)
	// 	fmt.Println(x)
	// 	wg.Done()
	// }(8.12, &wg)

	// wg.Wait()

	//

	// var wg sync.WaitGroup

	// wg.Add(50)
	// for x := 100; x < 150; x++ {
	// 	go func(value float64, wg *sync.WaitGroup) {
	// 		x := math.Sqrt(value)
	// 		fmt.Printf("Square root of %.2f is %.6f\n", value, x)
	// 		wg.Done()
	// 	}(float64(x), &wg)
	// }

	// wg.Wait()

	//

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()

	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()

	wg.Done()
}
