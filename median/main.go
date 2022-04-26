package main

import (
	"fmt"
	"sort"
)

/*
median of values:
	- sort values
	- if odd number of values: return middle
	- if even number of values: return average of middles
*/
func main() {
	nums := []int{3,4,5}
	fmt.Println("len:", len(nums))
	// Slice elements are mutable, but length is not, cannot append in place
	
	nums[1] = 200
	fmt.Println("1:", nums[1])

	for i := range nums {
		fmt.Println(i)
	}

	for i, v := range nums{
		fmt.Println("index:", i)
		fmt.Println("value:", v)
	}
	// Append has to reallocate to a new variable
	nums = append(nums, 40)
	fmt.Println("New slice after append:", nums)
	// Appending slices together
	s := []int{7,8,9}
	nums = append(nums, s...)
	fmt.Println("Nums after appending slices:", nums)
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}
	// Use our own copy of values in order not to change it
	vs := make([]float64, len(values))
	copy(vs, values)

	sort.Float64s(vs)
	i := len(vs) / 2
	if len(vs)%2 == 1 {
		return vs[i], nil
	}
	v := (vs[i-1] + vs[i]) / 2
	return v, nil
}