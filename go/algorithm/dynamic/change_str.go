package dynamic

import "fmt"

func ChangeStr(src, des []byte) int {
    // 两个字符串间的编辑距离

    res := [100][100]int{}
    n, m := len(src), len(des)
    for i := 1; i <= n; i ++ {
        res[i][0] = i
    }
    for i := 1; i <= m; i ++ {
        res[0][i] = i
    }

    for i := 0; i < n; i ++ {
        for j := 0; j < m; j ++ {
            if src[i] == des[j] {
                res[i+1][j+1] = minArray(res[i][j], res[i][j+1]+1, res[i+1][j]+1)
            } else {
                res[i+1][j+1] = minArray(res[i][j] + 1, res[i][j+1]+1, res[i+1][j]+1)
            }
        }
    }
    fmt.Println(res[n][m])
    return res[n][m]
}

func SChangeStr(src, des []byte) int {
    n, m := len(src), len(des)
    res := make([]int, m+1)
    last := make([]int, m+1)
    for i := 1; i <= m; i ++ {
        res[i], last[i] = i, i
    }

    for i := 0; i < n; i ++ {
        res[0] = i+1
        for j := 0; j < m; j ++ {
            if src[i] == des[j] {
                res[j+1] = minArray(last[j], res[j+1]+1, res[j]+1)
            } else {
                res[j+1] = minArray(last[j] + 1, res[j+1]+1, res[j]+1)
            }
            last[j] = res[j]
        }
    }
    return res[m]
}
