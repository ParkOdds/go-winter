package main

import "fmt"

func main1() {
	fmt.Println("hello world")

	// arr := []int{2, 4, 4, 5}
	arr := make([]int, 4) // make[type, amount of array]
	arr[0] = 30           // assign specific array
	fmt.Println(arr)

	txt := "today is sunday"
	fmt.Println(txt[0:5]) // slide
	fmt.Println(len(txt)) // lenght
}
