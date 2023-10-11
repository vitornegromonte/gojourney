package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("An empty array looks like: ", a)

	a[4] = 100
	fmt.Println("Changing array's index 4, from 0 to 100: ", a)
	fmt.Println("Getting value using array[index]: ", a[4])

	fmt.Println("Array's length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Two-dimensional array: ", twoD)
}
