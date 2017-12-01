package dynamic

import "fmt"

func LCS(a, b []byte) {
    res := [100][100]int{}
    for i := 0; i< len(a); i ++ {
        for j := 0; j < len(b); j++ {
            if a[i] == b[j] {
                res[i+1][j+1] = res[i][j] + 1
            } else {
                res[i+1][j+1] = max(res[i+1][j], res[i][j+1])
            }
        }
    }
    lcsLen := res[len(a)][len(b)]
    str := make([]byte, lcsLen)
    for i, j := len(a), len(b); i > 0;{
        if res[i][j] > res[i-1][j] && res[i][j] > res[i][j-1] {
            str[lcsLen-1] = a[i-1]
            i --; j --; lcsLen --
        } else if res[i][j] == res[i-1][j] {
            i --
        } else {
            j --
        }
        if lcsLen == 0 {
            break
        }
    }
    fmt.Println(string(str))
}
