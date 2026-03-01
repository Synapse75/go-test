package main

import "fmt"

func main() {
	var name string
	name = "World"
	fmt.Println(name)

	var name1 = "Go"
	fmt.Println(name1)

	name2 := "Programming"
	fmt.Println(name2, "is fun!", name1)

	var a byte = 'a'
	fmt.Printf("The value of a is: %d\n", a)

	var b bool = true
	fmt.Printf("The value of b is: %t\n", b)

	var nameList [3]string = [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Name List:", nameList)

	var ageList = make([]int, 3, 5)
	fmt.Println("Age List:", ageList)
}
