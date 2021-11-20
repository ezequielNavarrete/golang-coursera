package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Findian() {
	fmt.Printf("please, enter a string\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	lowerStr := strings.ToLower(line)
	if strings.HasPrefix(lowerStr, "i") && strings.HasSuffix(lowerStr, "n") && strings.Contains(lowerStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
