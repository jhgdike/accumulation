# coding: utf-8

import os
import sys
from unittest import TestCase

sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


class BaseTestCase(TestCase):

    def setUp(self):
        pass
