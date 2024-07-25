package main

import "fmt"

//to load run go mod tidy
import "rsc.io/quote"

func main() {

	//print statement in go
	fmt.Println("Hello world")
	//use of import statement
	fmt.Println(quote.Go())

}
