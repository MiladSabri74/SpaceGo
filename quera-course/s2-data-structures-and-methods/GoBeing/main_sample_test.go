package main

import (
    "testing"
)

func TestFormatTimeLimit_Sample(t *testing.T) {
    expected := "02:30:15"
    result := FormatTimeLimit(2, 30, 15)

    if result != expected {
        t.Errorf("Expected %q, but got %q", expected, result)
    }
}

func TestProcessScreenTime_Sample(t *testing.T) {
    wantH, wantM, wantS := 1, 5, 30
    gotH, gotM, gotS := ProcessScreenTime(3930)

    if gotH != wantH || gotM != wantM || gotS != wantS {
        t.Errorf("Expected (%d, %d, %d), but got (%d, %d, %d)", wantH, wantM, wantS, gotH, gotM, gotS)
    }
}
