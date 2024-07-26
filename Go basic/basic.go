package main

import "fmt"

//to load run go mod tidy
import "rsc.io/quote"

func main() {

	//print statement in go
	fmt.Println("Hello world")
	//use of import statement
	fmt.Println(quote.Go())

	//declare in go
	a := 2
	fmt.Println(a)

	count := 10

	//for loop in Go 
	for i := 0; i < count; i++ {
		
		// If statement in Go
		if i == 5 {
			fmt.Println("Goooooo")
			continue
		}
		fmt.Println(i)
	}
}
