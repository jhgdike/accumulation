package KMP

import (
    "testing"
    "fmt"
)

func TestCalNext(t *testing.T) {
    s := "ababaca"
    fmt.Println(calNext(s))
}

func TestGetNext(t *testing.T) {
    s := "ababaca"
    fmt.Println(getNext(s))
}

func TestKMP(t *testing.T) {
    src, desc := "测abcdefg", "bcde"
    startIndex := KMP(src, desc)
    startIndexB := KMP_BASIC(src, desc)
    if 4 != startIndex || startIndex != startIndexB {
        t.Error("KMP error: result: ", startIndex, startIndexB)
    }
}
