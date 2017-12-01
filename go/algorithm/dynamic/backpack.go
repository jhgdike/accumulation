package dynamic

import (
    "fmt"
)

func SimpleBackPack() {
    /* N件物品， V容量，损耗Ci, 收益Wi；求最大收益
    状态变化，转移矩阵
    第I件,
    d[i, v] = {
        max(d[i-1, v], [d[i-1, v-ci] + wi])
    */
    V := 100
    C := []int{10, 15, 20, 13, 12, 19, 28, 31}
    W := []int{10, 10, 23, 42, 12, 32, 20, 31}

    F := make([]int, V + 1)
    for i := 0; i < len(C); i ++ {
        ZeroOnePack(F, C[i], W[i])
    }
    fmt.Println(F[V])  // 140
}

func ZeroOnePack(F []int, c, w int) {
    // 从后往前的原因是防止前面影响后面的结果。第k个v，前面应该是未用过的状态
    for i := len(F) - 1; i >= c; i -- {
        F[i] = max(F[i], F[i-c] + w)
    }
}

func CompleteBackPack() {
    V := 100
    C := []int{10, 15, 20, 13, 12, 19, 28, 31}
    W := []int{10, 10, 23, 42, 12, 32, 20, 31}

    F := make([]int, V + 1)
    for i := 0; i < len(C); i ++ {
        CompletePack(F, C[i], W[i])
    }
    fmt.Println(F[V])  // 140
}

func CompletePack(F []int, c, w int) {
    // 从前往后而且只循环一次的原因是，前面已经加过的v可以无限次数的加入到后面，及时已经加过也不影响后面的结果
    for i := c; i < len(F); i ++ {
        F[i] = max(F[i], F[i - c] + w)
    }
}

func MultiBackPack() {
    V := 100
    C := []int{10, 15, 20, 13, 12, 19, 28, 31}
    M := []int{2, 3, 6, 7, 5, 3, 2, 5, 1, 6}
    W := []int{10, 10, 23, 42, 12, 32, 20, 31}

    F := make([]int, V + 1)
    for i := 0; i < len(C); i ++ {
        MultiPack(F, C[i], W[i], M[i])
    }
    fmt.Println(F[V])  // 140
}

func MultiPack( F []int, c, w, m int) {
    if m * c >= len(F) - 1 {
        CompletePack(F, c, w)
    } else {
        for p := 1; p < m; p <<= 1 {
            ZeroOnePack(F, p*c, p*w)
        }
        ZeroOnePack(F, m*c, m*w)
    }
}
