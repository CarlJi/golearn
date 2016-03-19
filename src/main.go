package main

import "fmt"

func main() {
	//strings.Map()

	//strconv.ParseBool()

	//regexp.Match()

	//fmt.Printf("Result: %v ", res)

	counterA := createCounter(2)
	counterB := createCounter(102)

	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A->%d, B->%d)", a, <-counterB)
	}

	fmt.Println()
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)

	return next
}
