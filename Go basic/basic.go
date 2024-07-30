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

	//map in Go with val
	

	 map_2 := map[int]string{
     
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
 
	fmt.Println("Map-2: ", map_2)

	var My_map = make(map[float64]string)
    fmt.Println(My_map)
 
    // As we already know that make() function
    // always returns a map which is initialized
    // So, we can add values in it
    My_map[1.3] = "Rohit"
    My_map[1.5] = "Sumit"
    fmt.Println(My_map)

	// Iterating map using for range loop
    for id, pet := range map_2 {
 
        fmt.Println(id, pet)
    }

	value_1 := map_2[90]
    fmt.Println("Value of key[90]: ", value_1)

	// Checking the key is available
    // or not in the m_a_p map
    pet_name0, ok0 := map_2[90]
    fmt.Println("\nKey present or not:", ok0)
    fmt.Println("Value:", pet_name0)
 
    // Using blank identifier
    _, ok11 := map_2[90]
    fmt.Println("\nKey present or not:", ok11)

	// Checking the key is available
    // or not in the m_a_p map
    pet_name, ok := map_2[99]
    fmt.Println("\nKey present or not:", ok)
    fmt.Println("Value:", pet_name)
 
    // Using blank identifier
    _, ok1 := map_2[99]
    fmt.Println("\nKey present or not:", ok1)

	// Deleting keys
    // Using delete function
	fmt.Println("Before Deleteing : ",map_2)
    delete(map_2, 93)
	fmt.Println("After Deleteing : ",map_2)
}
