package exercises

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Read() {
	type Name struct {
		Fname string
		Lname string
	}

	absPath, _ := filepath.Abs("../tp4/names.txt")
	file, _ := os.Open(absPath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var names []Name
	for scanner.Scan() {
		var name Name
		completeName := strings.Split(scanner.Text(), " ")
		name.Fname = completeName[0]
		name.Lname = completeName[1]
		names = append(names, name)
	}
	fmt.Println(names)
}
