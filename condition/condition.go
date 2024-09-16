package main

import "fmt"

func main() {

	//if else statement in Go

	var a int = 10
	if a < 20 {
		fmt.Println("a is less than 20")
	} else {
		fmt.Println("a is not less than 20")
	}

	// switch statement in Go

	var plan string = "retail"
	switch plan {
	case "retail":
		fmt.Println("Retail Plan")
	case "wholesale":
		fmt.Println("Wholesale Plan")
	case "dealer":
		fmt.Println("Dealer Plan")
	default:
		fmt.Println("Default Plan")
	}

	//in Go break state is is implicit if you want pass that case to next then use fallthrough

	var plan2 string = "dealer"
	switch plan2 {
	case "retail":
		fmt.Println("Retail Plan")
	case "wholesale":
		fmt.Println("Wholesale Plan")
	case "dealer":
		fallthrough
	case "test":
		fmt.Println("Test Plan")
	default:
		fmt.Println("Default Plan")
	}
}
