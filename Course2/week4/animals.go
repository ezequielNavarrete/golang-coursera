package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type AnimalsQualities struct {
	Food       string
	Locomotion string
	Noise      string
}
type Cow struct {
	Qualities AnimalsQualities
}
type Bird struct {
	Qualities AnimalsQualities
}
type Snake struct {
	Qualities AnimalsQualities
}

func (cow Cow) Eat() {
	fmt.Println(cow.Qualities.Food)
}
func (cow Cow) Move() {
	fmt.Println(cow.Qualities.Locomotion)
}
func (cow Cow) Speak() {
	fmt.Println(cow.Qualities.Noise)
}
func (bird Bird) Eat() {
	fmt.Println(bird.Qualities.Food)
}
func (bird Bird) Move() {
	fmt.Println(bird.Qualities.Locomotion)
}
func (bird Bird) Speak() {
	fmt.Println(bird.Qualities.Noise)
}
func (snake Snake) Eat() {
	fmt.Println(snake.Qualities.Food)
}
func (snake Snake) Move() {
	fmt.Println(snake.Qualities.Locomotion)
}
func (snake Snake) Speak() {
	fmt.Println(snake.Qualities.Noise)
}

func (q *AnimalsQualities) InitMe(f, l, n string) {
	q.Food = f
	q.Locomotion = l
	q.Noise = n
}

func (cow *Cow) InitMe() {
	cow.Qualities.InitMe("grass", "walk", "moo")
}
func (bird *Bird) InitMe() {
	bird.Qualities.InitMe("worms", "fly", "peep")
}
func (snake *Snake) InitMe() {
	snake.Qualities.InitMe("mice", "slither", "hsss")
}
func main() {
	animals := make(map[string]interface{})
	var cow Cow
	var bird Bird
	var snake Snake
	cow.InitMe()
	bird.InitMe()
	snake.InitMe()
	var command, name, typeOrAction string
	for {
		fmt.Print(">")
		fmt.Scan(&command)
		fmt.Scan(&name)
		fmt.Scan(&typeOrAction)
		switch command {
		case "newanimal":
			switch typeOrAction {
			case "cow":
				animals[name] = cow
				fmt.Println("Created it!")
			case "bird":
				animals[name] = bird
				fmt.Println("Created it!")
			case "snake":
				animals[name] = snake
				fmt.Println("Created it!")
			}
		case "query":
			value, _ := animals[name]
			switch an := value.(type) {
			case Cow:
				actionCase(an, typeOrAction)
			case Bird:
				actionCase(an, typeOrAction)
			case Snake:
				actionCase(an, typeOrAction)
			}
		}
	}
}

func actionCase(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}
