package main

import "fmt"

type arrType []string

func (arr arrType) printMe() {
	for i, item := range arr {
		fmt.Println(i, item)
	}
}
