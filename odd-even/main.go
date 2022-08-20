package main

import "fmt"

func main() {
	// just created a simple int slice from 0 to 10.
	s := []int{0,1,2,3,4,5,6,7,8,9,10}

	// this loops through the slice range and
	// compare using module to check if the number
	// is even or odd.
	for i := range s {
		if i % 2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		} 
	}
}