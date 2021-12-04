package main

import (
	"fmt"
)

/*
func main() {
	slice := make([]int, 0, 3)
	var number string
	for number != "X" {
		fmt.Printf("Please, enter a integer or X to finish the program\n")
		fmt.Scan(&number)
		if number != "X" {
			num, _ := strconv.Atoi(number)
			slice = append(slice, int(num))
		}
	}
	sort.Ints(slice)
	fmt.Println(slice)
}
*/
func main() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[0:4]
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(len(y), cap(y), len(z), cap(z))
}
