package main

import (
	"fmt"
	"strings"
)

type hola []int

func main() {
	var arr = hola{1, 2, 3, 4, 5}
	fmt.Println(arr, len(arr))
	var fullname = []string{"yogendra", "saxena"}
	var res = strings.Join(fullname, " ")
	fmt.Println(res)
	var stringArr = strings.Split(res, " ")
	fmt.Println(len(stringArr))
	arr.logme("yogendra")
}

func (h hola) logme(name string) {
	fmt.Println("logged", h, name)
}
