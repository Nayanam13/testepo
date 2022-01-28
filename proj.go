package main

import "fmt"

func main() {
	name := "Nayanam"
	y := &name
	changeValue(y)
	fmt.Println(name)

}

func changeValue(str *string) int {
	*str = "Abhyam"
	return 1
}
