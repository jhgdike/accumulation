package b_tree

import (
    // "fmt"
    "testing"
    "fmt"
)

func TestTree(t *testing.T) {
    tree := &Tree{}
    arr := [9]int{5, 2, 7, 8, 1, 4, 6, 9, 10}
    for _, v := range arr {
        tree.Insert(v)
    }
    tree.PreOrder()  // [5 2 1 4 7 6 8 9 10]
    tree.InOrder()   // [1 2 4 5 6 7 8 9 10]
    tree.PostOrder() // [1 4 2 6 10 9 8 7 5]
}

func TestAlth(t *testing.T) {
    pre := []int{5, 2, 1, 4, 7, 6, 8, 9, 10}
    in := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
    tree := RecoverTree(pre, in)
    tree.PostOrder()  // [1 4 2 6 10 9 8 7 5]
}

func TestPost(t *testing.T) {
    pre := []int{5, 2, 1, 4, 7, 6, 8, 9, 10}
    in := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
    post := GetPostOrder(pre, in)
    fmt.Println(post) // [1 4 2 6 10 9 8 7 5]
}
