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
	fmt.Println("committed")
	return 1
}
