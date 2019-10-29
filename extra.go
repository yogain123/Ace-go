package main

import "fmt"

func main() {
	for _, item := range []int{1, 2, 3, 4, 5} {
		fmt.Println(item)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("i", i)
	}

	var j = 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	m := map[string]string{
		"name": "yogendra",
	}

	fmt.Println(m)
}
