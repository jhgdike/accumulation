def cal_next(pat):
    next = [-1] * len(pat)
    k = -1
    i = 0
    while i < len(pat):
        if k == -1 or pat[i] == pat[k]:
            next[i] = k
            i, k = i + 1, k + 1
        else:
            k = next[k]
    return next


def KMP(long_str, s_str):
    next = cal_next(s_str)
    i, k = 0, -1
    while i < len(long_str) and k < len(s_str):
        if k == -1 or long_str[i] == s_str[k]:
            i, k = i + 1, k + 1
        else:
            k = next[k]
    if k == len(s_str):
        return i - k
    return -1
