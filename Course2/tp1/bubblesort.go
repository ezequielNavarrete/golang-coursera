package main

import "fmt"

func main() {
	var number int
	var slice []int
	for i := 0; i < 10; i++ {
		fmt.Printf("Please, enter a integer\n")
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	BubbleSort(slice)
	fmt.Println(slice)
}

func BubbleSort(slice []int) {

	for range slice {
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
			}
		}
	}
}

func Swap(slice []int, pos int) {
	aux := slice[pos]
	slice[pos] = slice[pos+1]
	slice[pos+1] = aux
}
