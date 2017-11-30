package b_tree

import "fmt"

func RecoverTree(pre, in []int) *Tree {
    tree := NewTree()
    tree.root = &Node{val: pre[0]}
    if len(pre) <= 1 {
        return tree
    }
    recovery(tree.root, pre, in)
    return tree
}

func recovery(root *Node, pre, in []int) {
    if len(pre) <= 1 {
        return
    }
    index := search(pre[0], in)
    fmt.Println(pre, in)
    if index > 0 {
        root.left = &Node{val: pre[1]}
        recovery(root.left, pre[1:index + 1], in[:index])
    }
    if index < len(in) - 1 {
        root.right = &Node{val: pre[index+1]}
        recovery(root.right, pre[index+1:], in[index+1:])
    }
}

func search(src int, arr []int) int {
    for i, v := range arr {
        if v == src {
            return i
        }
    }
    return -1
}

func GetPostOrder(pre, in []int) []int {
    res := make([]int, len(pre))

    post(pre, in, res)
    return res
}

func post(pre, in, res []int) {
    if len(pre) == 0 {
        return
    }
    res[len(pre) - 1] = pre[0]
    split := search(pre[0], in)
    if split > 0 {
        post(pre[1:split+1], in[:split], res[:split])
    }
    if split < len(pre) - 1 {
        post(pre[split+1:], in[split+1:], res[split:len(pre)-1])
    }
}
