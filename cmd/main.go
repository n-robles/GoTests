package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Behavior interface {
	sleep()
	eat()
}

type Cat struct {
	Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

func main() {
	doggy := &Dog{
		Animal{
			"Rover", // name
			3,       // age
		},
		5,
	}

	catsy := &Cat{
		Animal{
			"Garfield", // name
			5,          // age
		},
		2,
	}

	doggy.DoSomething()
	catsy.DoSomething()

	catsy.DoSomething()
	doggy.DoSomething()
}

func (a *Animal) DoSomething() {
	fmt.Printf(" %s is doing some stuff", a.name)
}

func (cat *Cat) sleep() {
	fmt.Printf("The CAT %s is sleeping", cat.name)
}

func (cat *Cat) eat() {
	fmt.Printf("The CAT %s is sleeping, it won't eat", cat.name)
}

func (dog *Dog) sleep() {
	fmt.Printf("The DOG %s don't want to sleep", dog.name)
}

func (dog *Dog) eat() {
	fmt.Printf("The DOG %s is eating", dog.name)
}
