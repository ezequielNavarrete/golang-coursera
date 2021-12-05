package main

import (
	"fmt"
	"sort"
	"strconv"
)

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
