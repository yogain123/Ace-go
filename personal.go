package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

type hola []int

type Person struct {
	name string
	age  int
}

func main() {
	var arr = hola{1, 2, 3, 4, 5, 9, 2}
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Reverse(sort.IntSlice(s))
	fmt.Println(s)
	fmt.Println(arr, len(arr))
	var fullname = []string{"yogendra", "saxena"}
	var res = strings.Join(fullname, " ")
	fmt.Println(res)
	var stringArr = strings.Split(res, " ")
	fmt.Println(len(stringArr))
	arr.logme("yogendra")

	var p Person
	p.name = "yogendra"
	p.age = 25
	fmt.Println((p))
	updateName(&p)
	fmt.Println(p)
	p.updateAge()
	fmt.Println(p)

	c := make(chan int) // making channel

	for _, item := range arr {
		go delay(item, c) // go routin
		//fmt.Println(<-c)  // receiving data from a channel is a blocking operation
		// channelValue := <-c
		// fmt.Println(channelValue)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(<-c)
	}
}

func delay(item int, c chan int) {
	_, error := http.Get("https://www.google.com")
	if error != nil {
		fmt.Println("Error")
		c <- 404 // putting data in channel
		os.Exit(1)
	}
	//fmt.Println("success")
	c <- 200 // putting data in channel

}

func updateName(p *Person) {
	(*p).name = "saxena"
}

func (p *Person) updateAge() {
	(*p).age = 35
}

func (h hola) logme(name string) {
	fmt.Println("logged", h, name)
}
