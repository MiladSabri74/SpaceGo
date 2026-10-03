package main

import (
    "math"
    )
func AddElement(numbers *[]int, element int) {
	//TODO
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	//TODO
	if len(*numbers) == 0 {
		return 0
	}
	minimum := math.MaxInt
	for _, element := range *numbers {
		if element < minimum {
			minimum = element
		}
	}
	return minimum
}

func ReverseSlice(numbers *[]int) {
	//TODO
	sliceSize := len(*numbers)
	for i := 0; i < sliceSize/2; i++ {
		temp := (*numbers)[i]
		(*numbers)[i] = (*numbers)[sliceSize-i-1]
		(*numbers)[sliceSize-i-1] = temp
	}
}

func SwapElements(numbers *[]int, i, j int) {
	//TODO
	if i >= len(*numbers) || j >= len(*numbers) || i < 0 || j  < 0 {
	    return
	}
	temp := (*numbers)[i]
	(*numbers)[i] = (*numbers)[j]
	(*numbers)[j] = temp
}
