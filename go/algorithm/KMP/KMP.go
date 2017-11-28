package KMP

/*
    字符串匹配。给你两个字符串，寻找其中一个字符串是否包含另一个字符串，如果包含，返回包含的起始位置。
    字符串长度分别为n,m (n > m).
    算法复杂度O(n + m).

    充分利用了目标字符串ptr的性质（比如里面部分字符串的重复性，即使不存在重复字段，在比较时，实现最大的移动量）。


    考察目标字符串ptr：
    ababaca
    这里我们要计算一个长度为m的转移函数next。

    next数组的含义就是一个固定字符串的最长前缀和最长后缀相同的长度。

    比如：abcjkdabc，那么这个数组的最长前缀和最长后缀相同必然是abc。
    cbcbc，最长前缀和最长后缀相同是cbc。
    abcbc，最长前缀和最长后缀相同是不存在的。

    **注意最长前缀：是说以第一个字符开始，但是不包含最后一个字符。比如aaaa相同的最长前缀和最长后缀是aaa。**

    对于目标字符串ptr = ababaca，长度是7，所以next[0]，next[1]，next[2]，next[3]，next[4]，next[5]，next[6]分别计算的是
    a，ab，aba，abab，ababa，ababac，ababaca的相同的最长前缀和最长后缀的长度。

    由于a，ab，aba，abab，ababa，ababac，ababaca的相同的最长前缀和最长后缀是“”，“”，“a”，“ab”，“aba”，“”，“a”,
    所以next数组的值是[-1,-1,0,1,2,-1,0]，这里-1表示不存在，0表示存在长度为1，2表示存在长度为3。这是为了和代码相对应。


    *** 求解next过程：
        1. 对于字符串ptr, 第一位是-1(p=0,k=-1, next[p]=k);
        2. p,k 从第二位开始（数组下标为1, p=1,k=0, 当第k位等于第p位时或者k==-1时，next[p]=k, p+=1,k+=1;继续步骤2;
        3. 否则回溯k=next[k]，继续步骤2，直到p = len(ptr) - 1;
 */

func calNext(str string) []int {
    slen := len(str)
    next := make([]int, slen)
    next[0] = -1
    k := -1
    for q := 1; q < slen; q ++ {
        for k > -1 && str[k+1] != str[q] {
            k = next[k]
        }
        if str[k+1] == str[q] {
            k += 1
        }
        next[q] = k
    }
    return next
}

func KMP(src, reg string) int {
    next := calNext(reg)
    k := -1
    for i := 0; i < len(src); i ++ {
        for k > -1 && reg[k+1] != src[i] {
            k = next[k]
        }
        if reg[k+1] == src[i] {
            k += 1
        }
        if k == len(reg) - 1 {
            return i - k
        }
    }
    return -1
}

func getNext(str string) []int {
    slen := len(str)
    next := make([]int, slen)
    k := -1
    for j := 0; j < slen; {
        if k == -1 || str[j] == str[k] {
            next[j] = k
            j, k = j+1, k+1
        } else {
            k = next[k]
        }
    }
    return next
}

func KMP_BASIC(src, reg string) int {
    next := getNext(reg)
    i, j := 0, 0
    for i < len(src) && j < len(reg) {
        if j == -1 || src[i] == reg[j] {
            i, j = i+1, j+1
        } else {
            j = next[j]
        }
    }
    if j == len(reg) {
        return i - j
    }
    return -1
}
