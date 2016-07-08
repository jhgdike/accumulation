# coding: utf-8


def swap(heap, i, j):
    heap[i], heap[j] = heap[j], heap[i]


class Heap(object):

    heap = None

    def __init__(self, heap=None):
        self.heap = heap or []
        self.length = len(heap)
        # self._heap_list(0)

    def push(self, val):
        self.heap.append(val)
        i = self.length
        self.length += 1
        while i > 0:
            j = (i - 1) / 2
            if self.heap[i] < self.heap[j]:
                swap(self.heap, i, j)
            i = j

    def _heap_list_recursion(self, index):
        left = 2 * index + 1
        if 2 * (index + 1) < self.length:
            self._heap_list_recursion(left)
            self._heap_list_recursion(left + 1)
        self._cmp_and_swp(index, left, self.length)

    def _heappify(self, heap):
        """init the heap or called heapify"""
        for i in xrange(self.length - 1, 0, -1):
            j = (i - 1) / 2
            if heap[i] < heap[j]:
                swap(heap, i, j)

    def _cmp_and_swp(self, father, child, n):
        if child + 1 < n:
            if self.heap[child] < self.heap[child + 1] and \
                    self.heap[child] < self.heap[father]:
                swap(self.heap, father, child)
                return child
            elif self.heap[child + 1] < self.heap[child] and \
                    self.heap[child + 1] < self.heap[father]:
                swap(self.heap, child + 1, father)
                return child + 1
        elif child < n:
            if self.heap[child] < self.heap[father]:
                swap(self.heap, father, child)
                return child
        return father

    def _delete(self, n=None):
        if not n:
            n = len(self.heap)
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

    def pop(self):
        self._delete()
        self.length -= 1
        return self.heap.pop()

    def sort(self):
        for i in range(self.length - 1, 0, -1):
            self._delete(i)
