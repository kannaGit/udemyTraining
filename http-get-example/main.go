package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	_     = iota
	KB    = 1 << (iota * 10)
	myVal = 1 << ('\t' + 2.0)
)

func main() {
	res, _ := http.Get("http://www.dealigg.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", page)
	fmt.Println()
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Println()
	fmt.Printf("%d\n", myVal)
}
