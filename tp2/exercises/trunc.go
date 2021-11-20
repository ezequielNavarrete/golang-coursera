package exercises

import "fmt"

func Trunc() {
	fmt.Printf("Please, enter a floating point number\n")
	var fl float64
	fmt.Scan(&fl)
	fmt.Println(int(fl))
}
