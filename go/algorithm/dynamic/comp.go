package dynamic

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxByte(a, b byte) byte {
    if a > b {
        return a
    }
    return b
}

func minByte(a, b byte) byte {
    if a < b {
        return a
    }
    return b
}

func minArry(a ...int) int {
    tmp := a[0]
    for i := 1; i < len(a); i ++ {
        tmp = min(tmp, a[i])
    }
    return tmp
}
