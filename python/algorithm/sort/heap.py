# coding: utf-8


class Heap(object):

    heap = None

    def __init__(self, heap=None):
        self.heap = heap or []
        self.length = len(heap)
        # self._heap_list(0)
        self._heap_list_loop()

    def insert(self, val):
        self.heap.append(val)
        self.length += 1
        while i > 0:
            j = (i - 1) / 2
            if self.heap[i] < self.heap[j]:
                self.heap[i], self.heap[j] = self.heap[j], self.heap[i]
            i = j

    def _heap_list_recursion(self, index):
        left = 2 * index + 1
        if 2 * (index + 1) < self.length:
            self._heap_list_recursion(left)
            self._heap_list_recursion(left + 1)
        self._cmp_and_swp(index, left, self.length)

    def swap(self, a, b):
        self.heap[a], self.heap[b] = self.heap[b], self.heap[a]

    def _heap_list_loop(self):
        for i in xrange(self.length - 1, 0, -1):
            j = (i - 1) / 2
            if self.heap[i] < self.heap[j]:
                self.swap(i, j)

    def _cmp_and_swp(self, father, child, n):
        if child + 1 < n:
            if self.heap[child] < self.heap[child + 1] and \
                    self.heap[child] < self.heap[father]:
                self.swap(father, child)
                return child
            elif self.heap[child + 1] < self.heap[child] and \
                    self.heap[child + 1] < self.heap[father]:
                self.swap(child + 1, father)
                return child + 1
        elif child < n:
            if self.heap[child] < self.heap[father]:
                self.swap(father, child)
                return child
        return father

    def _delete(self, n):
        self.heap[0], self.heap[n] = self.heap[n], self.heap[0]
        i = 0
        while True:
            if 2 * i + 1 > n - 1:
                return
            if 2 * i + 1 == n - 1:
                self._cmp_and_swp(i, n - 1, n)
                return
            if 2 * i + 2 <= n - 1:
                index = self._cmp_and_swp(i, 2 * i + 1, n)
                if i == index:
                    return
                i = index

    def sort(self):
        for i in range(self.length - 1, 0, -1):
            self._delete(i)
