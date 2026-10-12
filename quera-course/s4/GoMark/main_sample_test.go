package main

import (
    "reflect"
    "testing"
)

func TestProcessImages_Simple(t *testing.T) {
    input := []string{"photo1", "photo2", "photo3"}

    expected := []string{
        "photo1_watermarked",
        "photo2_watermarked",
        "photo3_watermarked",
    }

    result := ProcessImages(input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("ProcessImages failed.\nExpected: %v\nGot: %v", expected, result)
    }
}
