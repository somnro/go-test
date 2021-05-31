package main

import "fmt"

func main() {
	arr1 := [4]int{1,2,6,8}
	fmt.Println(arr1)

	slice1 := []int{456,56}
	fmt.Println(len(slice1))
	fmt.Println("type: %T",slice1)
}
