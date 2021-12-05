package exercises

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func MakeJson() {
	var name string
	fmt.Printf("Please, enter a name\n")
	fmt.Scan(&name)
	fmt.Printf("Please, enter a adress\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	adress := scanner.Text()
	nameAndAdress := make(map[string]string, 2)
	nameAndAdress["name"] = name
	nameAndAdress["adress"] = adress
	barr, _ := json.Marshal(nameAndAdress)
	fmt.Println(string(barr))
}
