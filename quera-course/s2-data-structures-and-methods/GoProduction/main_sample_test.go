package main

import (
    "reflect"
    "testing"
)

func TestIsSquare_Sample(t *testing.T) {
    if got := IsSquare(25); got != true {
        t.Errorf("IsSquare(25) = %v; want true", got)
    }
}

func TestIsPalindrome_Sample(t *testing.T) {
    if got := IsPalindrome(12321); got != true {
        t.Errorf("IsPalindrome(12321) = %v; want true", got)
    }
}

func TestAbs_Sample(t *testing.T) {
    if got := Abs(-10); got != 10 {
        t.Errorf("Abs(-10) = %d; want 10", got)
    }
}

func TestCube_Sample(t *testing.T) {
    if got := Cube(3); got != 27 {
        t.Errorf("Cube(3) = %d; want 27", got)
    }
}

func TestFilter_Sample(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}
    isEven := func(x int) bool { return x%2 == 0 }
    expected := []int{2, 4, 6}

    got := Filter(input, isEven)
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("Filter() = %v; want %v", got, expected)
    }
}

func TestMap_Sample(t *testing.T) {
    input := []int{1, 2, 3}
    double := func(x int) int { return x * 2 }
    expected := []int{2, 4, 6}

    got := Map(input, double)
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("Map() = %v; want %v", got, expected)
    }
}
