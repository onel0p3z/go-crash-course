package main

import "fmt"

func main() {

	//long form
	// var name string = "juan"

	// short form
	// var name = "juan"

	// short form
	name := "juan"

	// short form
	name, age := "juan", 30

	fmt.Printf("name is a '%T'", name)
	fmt.Println(name, age)

	for i, letter := range name {
		fmt.Printf("%v ---> %s \n", i, string(letter))
	}

	fmt.Printf("'%s' len %v\n", name, len(name))

	for i := len(name) - 1; i >= 0; i-- {
		fmt.Printf("%s\n", string(name[i]))
	}

}
