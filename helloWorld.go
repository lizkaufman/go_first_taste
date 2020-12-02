package main

import "fmt"

var (
	myString     string
	secondString string
	myInt        int
)

const (
	constString string = "I'm a const!"
)

func init() {
	myString = "hello world"
	secondString = "hi again"
	myInt = 1
}

func sayMyName(name string) string {
	return name
}

func add(a int, b int) int {
	return a + b
}

func addMany(nums ...int) int {
	//variable number of integers
	total := 0
	for _, n := range nums {
		//like js forEach
		total += n
	}
	return total
}

func main() {
	fmt.Println(myString, secondString, constString, myInt)
	name := sayMyName("Liz")
	sum := add(myInt, 2)
	manySum := addMany(1, 2, 3, myInt)
	fmt.Println(name)
	fmt.Println(sum)
	fmt.Println(manySum)
	num := 28
	str := "age"
	together := fmt.Sprintf("property=%s value=%d", str, num) //formats values into placeholders
	fmt.Printf(together)
}
