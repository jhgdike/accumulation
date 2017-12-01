package dynamic

import "fmt"

func ChangeStr(src, des []byte) {
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
                res[i+1][j+1] = minArry(res[i][j], res[i][j+1]+1, res[i+1][j]+1)
            } else {
                res[i+1][j+1] = minArry(res[i][j] + 1, res[i][j+1]+1, res[i+1][j]+1)
            }
        }
    }
    fmt.Println(res[n][m])
}
