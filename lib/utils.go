package lib

import "fmt"

//GetMinMax returns the min and max of array of integers
func GetMinMax(data []int) (int, int, error) {
	if len(data) > 0 {
		min := data[0]
		max := data[0]
		for _, item := range data {
			if min > item {
				min = item
			}
			if max < item {
				max = item
			}
		}
		return min, max, nil
	}
	return 0, 0, fmt.Errorf("Can't find length of empty array")

}
