package main

import "fmt"

const milesToKMConversion float32 = 1.60934

func main() {
	var miles float32

	fmt.Print("Enter the Miles:")
	fmt.Scan(&miles)

	KM := miles * milesToKMConversion

	fmt.Println("Miles to KM Converted: ", miles, KM)
}
