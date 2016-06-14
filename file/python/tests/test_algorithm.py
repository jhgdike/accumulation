# -*- coding: utf-8 -*-

import os
import sys
import unittest

sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from algorithm.sort.binary_tree import BinaryTree, SimpleBinaryTree
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

if __name__ == '__main__':
    unittest.main()
