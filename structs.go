package main

import (
	"fmt"
)

type myStruct struct {
	//struct = class
	FieldA string
	FieldB int
}

//have to make your own constructor - convention is start it w/ New and then the struct name
func NewMyStruct(a string, b int) myStruct {
	return myStruct{
		FieldA: a,
		FieldB: b,
	}
}

type Person struct {
	Name         string
	Age          int
	EyeColor     string
	HairColor    string
	IsProgrammer bool
}

func NewPerson(name string, age int, eyeColor string, hairColor string, isProgrammer bool) Person {
	return Person{
		Name:         name,
		Age:          age,
		EyeColor:     eyeColor,
		HairColor:    hairColor,
		IsProgrammer: isProgrammer,
	}
}

//method on Person w/ reciever that ties it to the type Person
func (this Person) Details() string {
	//receiver before function name w/ var name and type (using this for similarity to JS); convention is first letter of type (would be p here normally)
	return this.Name + " has " + this.EyeColor + " eyes and " + this.HairColor + " hair."
}

func (p Person) areTheyProgrammer() bool {
	return p.IsProgrammer
}

func main() {
	liz := NewPerson("Liz", 28, "blue", "blonde", true)
	fmt.Println(liz)
	fmt.Printf("name: %s, age: %d, eyes: %s, hair: %s, programmer: %t", liz.Name, liz.Age, liz.EyeColor, liz.HairColor, liz.IsProgrammer)
	fmt.Println("\ndetails: " + liz.Details())
}
