package main

import "fmt"

type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.Food)
}
func (animal *Animal) Move() {
	fmt.Println(animal.Locomotion)
}
func (animal *Animal) Speak() {
	fmt.Println(animal.Noise)
}
func (animal *Animal) InitMe(food, locomotion, noise string) {
	animal.Food = food
	animal.Locomotion = locomotion
	animal.Noise = noise
}
func main() {
	var cow, bird, snake Animal
	cow.InitMe("grass", "walk", "moo")
	bird.InitMe("worms", "fly", "peep")
	snake.InitMe("mice", "slither", "hsss")
	var animal, action string
	for {
		fmt.Print(">")
		fmt.Scan(&animal)
		fmt.Scan(&action)
		switch animal {
		case "cow":
			cow.actionCase(action)
		case "bird":
			bird.actionCase(action)
		case "snake":
			snake.actionCase(action)
		}
	}
}

func (animal *Animal) actionCase(action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}
