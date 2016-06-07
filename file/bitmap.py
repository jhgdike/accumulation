# coding: utf-8


class BitMap(object):
    """simple bitmap
    :param max_val: int
    this class can mark numbers between 0 and max_val(max_val > 0) with few
    space and return the iterator.
    """

    def __init__(self, max_val):
        self._bitmap = [0] * (max_val / 32 + 1)

    def mark(self, val):
        index = val / 32
        bit = val % 32
        self._bitmap[index] |= 2 ** bit

    def get(self):
        for index, item in enumerate(self._bitmap):
            for i in range(32):
                if 2 ** i & item:
                    yield index * 32 + i
