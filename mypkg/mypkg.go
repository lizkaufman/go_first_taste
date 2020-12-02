package mypkg

import "fmt"

const (
	PublicConstant  = "public!"
	PrivateConstant = "private"
)

func init() {
	fmt.Println("public: " + PublicConstant)
	fmt.Println("private: " + PrivateConstant)
}
