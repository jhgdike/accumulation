package dynamic

import (
    "testing"
    "fmt"
)

func TestMakeChange(t *testing.T) {
    valueKinds := []int{25, 21, 10, 5, 1}
    money := 68
    println(MakeChange(money, valueKinds))
}

func TestZeroOnePack(t *testing.T) {
    SimpleBackPack()
}

func TestCompleteBackPack(t *testing.T) {
    CompleteBackPack()
}

func TestMultiBackPack(t *testing.T) {
    MultiBackPack()
}

func TestLCS(t *testing.T) {
    LCS([]byte("adsijlqwerqadfasfdaewrtqw"), []byte("aisdfjoalqejwrtqlwejrasdfafdill"))
}

func TestChangeStr(t *testing.T) {
    fmt.Println(ChangeStr([]byte("sailn"), []byte("failing")) == SChangeStr([]byte("sailn"), []byte("failing")))
}

func TestMaxSubArray(t *testing.T) {
    fmt.Println(MaxSubArray([]int{1, 4, -6, 1, 6, 2, -4, 3, 1, 6, 8, -9, -3, 2, 4, 6}))
}

func TestLongestIncreasingSubs(t *testing.T) {
    LongestIncreasingSubs([]byte{35, 36, 39, 3, 15, 27, 6})
}