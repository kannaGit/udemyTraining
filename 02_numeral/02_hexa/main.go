package main

import "fmt"

func main() {
	//	fmt.Printf("Print in binary: %d - %b - %x \n", 42, 42, 42)
	//	fmt.Printf("Print in binary: %d \n %b \n %#X \n", 42, 42, 42)
	for i := 1; i < 100; i++ {
		fmt.Printf("Print in binary: %d \t %b \t %x \t %s \n", i, i, i)
	}
}
