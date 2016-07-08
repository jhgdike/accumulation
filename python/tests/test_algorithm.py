# -*- coding: utf-8 -*-

import unittest

from algorithm.sort.binary_tree import BinaryTree, SimpleBinaryTree
from algorithm.sort.heap import Heap
from tests import BaseTestCase as TestCase


class SimpleBinaryTreeTest(TestCase):

    def test_binary_tree(self):
        bt = SimpleBinaryTree()
        bt.add(20)
        bt.add(10)
        bt.add(30)
        bt.add(15)
        bt.add(25)
        assert bt.travel() == [10, 15, 20, 25, 30]


class BinaryTreeTest(TestCase):

    def test_binary_tree(self):
        bt = BinaryTree()
        bt.add(20)
        bt.add(10)
        bt.add(30)
        bt.add(15)
        bt.add(25)
        assert bt.travel() == [10, 15, 20, 25, 30]


class HeapTest(TestCase):

    def test_heap(self):
        a = [9, 12, 17, 30, 50, 20, 60, 65, 4, 49]
        heap = Heap(a)
        heap._heappify(a)
        heap.sort()
        assert a == [65, 60, 50, 49, 30, 20, 17, 12, 9, 4]


if __name__ == '__main__':
    unittest.main()
