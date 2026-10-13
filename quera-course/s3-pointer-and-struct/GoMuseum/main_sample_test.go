package main

import (
    "reflect"
    "testing"
)

func TestAddElement_Sample(t *testing.T) {
    nums := []int{1, 2}
    AddElement(&nums, 3)

    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Expected %v, got %v", expected, nums)
    }
}

func TestFindMin_Sample(t *testing.T) {
    nums := []int{8, 3, 5, 2, 9}
    minimum := FindMin(&nums)

    if minimum != 2 {
        t.Errorf("Expected minimum to be 2, got %d", minimum)
    }
}

func TestReverseSlice_Sample(t *testing.T) {
    nums := []int{1, 2, 3, 4}
    ReverseSlice(&nums)

    expected := []int{4, 3, 2, 1}
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Expected %v, got %v", expected, nums)
    }
}

func TestSwapElements_Sample(t *testing.T) {
    nums := []int{10, 20, 30}
    SwapElements(&nums, 0, 2)

    expected := []int{30, 20, 10}
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Expected %v, got %v", expected, nums)
    }
}

func TestRemoveDuplicates_Sample(t *testing.T) {
    nums := []int{1, 2, 2, 3, 1}
    RemoveDuplicates(&nums)

    expected := []int{1, 2, 3}
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Expected %v, got %v", expected, nums)
    }
}

func TestTrimCapacity_Sample(t *testing.T) {
    // Create a slice with length 2 and capacity 10
    nums := make([]int, 2, 10)
    nums[0], nums[1] = 1, 2

    TrimCapacity(&nums)

    if cap(nums) != 2 {
        t.Errorf("Expected capacity to be reduced to 2, got %d", cap(nums))
    }
}
