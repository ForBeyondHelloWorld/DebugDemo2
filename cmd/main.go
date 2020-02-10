package main

import (
	"debugdemo2/lib"
	"fmt"
)

func main() {
	min, max, err := lib.GetMinMax([]int{5, 2, 4, 3, 6})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}
}
