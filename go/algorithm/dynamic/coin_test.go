package dynamic

import "testing"

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
