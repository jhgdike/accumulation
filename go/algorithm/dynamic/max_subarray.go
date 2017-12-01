package dynamic

import (
    "fmt"
)

func MaxSubArray(a []int) int {
    mx, cur := a[0], a[0]

    for i := 1; i < len(a); i ++ {
        if mx < 0 {
            if a[i] < 0 {
                mx = max(mx, a[i])
            } else {
                mx, cur = a[i], a[i]
            }
        } else {
            cur += a[i]
            if cur < 0 {
                cur = 0
            } else {
                mx = max(mx, cur)
            }
        }
    }
    return mx
}

func LongestIncreasingSub(str []byte) {
    lenSub := make([]int, len(str))
    indexArr := make([]int, len(str))
    lenSub[0] = 1
    for i := 0; i < len(str); i ++ {
        indexArr[i] = -1
    }

    for i :=1; i < len(str); i++ {
        for j := i-1; j >= 0; j -- {
            if str[i] > str[j] && lenSub[i] <= lenSub[j] {
                indexArr[i] = j
                lenSub[i] = lenSub[j] + 1
            } else {
                lenSub[i] = max(1, lenSub[i])
            }
        }
    }
    lenth := lenSub[len(str)-1]
    subStr := make([]byte, lenth)
    for index, j := indexArr[len(str)-1], lenth-1; index >= 0; j--{
        subStr[0] = 1
    }
}

func LongestIncreasingSubs(str []byte) {
    res := [100]int{}
    res[0] = 1

    for i :=1; i < len(str); i++ {
        for j := 0; j < i; j ++ {
            if str[i] > str[j] {
                res[i] = max(res[i], res[j] + 1)
            }
        }
    }
    maxLen := 1
    for _, v := range res {
        maxLen = max(maxLen, v)
    }
    fmt.Println(maxLen)
}
