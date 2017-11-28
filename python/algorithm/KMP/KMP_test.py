from kmp import cal_next, KMP


def test_cal_next():
    next = cal_next('ababaca')
    assert next == [-1, -1, 0, 1, 2, -1, 0]


def test_kmp():
    assert 4 == KMP('安保处abcdefg', 'bcde')

if __name__ == '__main__':
    test_cal_next()
    test_kmp()
