package main

import "fmt"

func main() {
	var name string = "yogendra saxena"
	age := 23
	pincode := getPincode()
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(pincode)
	cards := arrType{"608058", getPincode()}
	cards = append(cards, "560065")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println("====")
	cards.printMe()
}

func getPincode() string {
	return "208022"
}
