package main

import (
	"fmt"

	"./mypkg"
)

func main() {
	fmt.Println("main.public: " + mypkg.PublicConstant)
}
