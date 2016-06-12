# coding: utf-8


class Digraph(object):
    """
    Destiny: The beauty of programming.
    Title:   The longest of subsentence.
    Method Theory:  Digraph
    e.g.
        key, val
        I love
        love you
        love me
        you too
        you so
        so smart
        me too

        dg.add(key, val)
        dg.optimal
        = > 5
        dg.optimal_str
        = > 'I love you so smart'
    """

    def __init__(self):
        """
        self._dict: adjacency list 邻接表
        self._node: all the nodes
            type:
                0: end of the edge
                1: start of the edge
                2: both
            optimal:
                int: 以此单词开头的最长子句长度
        self._vertex: all the
        """
        self._dict = {}
        self.is_cal_node = False

    def add(self, key, val):
        if not self._dict.get(key):
            self._dict[key] = []
        self._dict[key].append(val)
        if self.is_cal_node:
            self.is_cal_node = False

    @property
    def vertex(self):
        """顶点(起始点)"""
        if not self.is_cal_node:
            self._cal_node()
        return self._vertex

    def _cal_node(self):
        self._node = {}
        for key, val in self._dict.items():
            if not self._node.get(key):
                self._node[key] = {"type": 1, "optimal": -1}
            if self._node[key]['type'] == 0:
                self._node[key]['type'] = 2
            for item in val:
                if not self._node.get(item):
                    self._node[item] = {"type": 0, "optimal": -1}
                elif self._node[item]['type'] == 1:
                    self._node[item]['type'] = 2
        self._vertex = []
        for name, desc in self._node.items():
            if desc['type'] == 1:
                self._vertex.append(name)

    def _cal_optimal(self):
        optimal = 0
        if not self.is_cal_node:
            self._cal_node()
            self.is_cal_node = True
            for item in self._vertex:
                opt = self._dynamic_cal_optimal(item)
                self._node[item]['optimal'] = opt
                if opt > optimal:
                    optimal = opt
                    self._start_node = item
        else:
            for item in self._vertex:
                opt = self._node[item]['optimal']
                if opt > optimal:
                    optimal = opt
        return optimal

    def _dynamic_cal_optimal(self, node):
        if not self._dict.get(node):
            return 1
        max_path_depth = 0
        for item in self._dict[node]:
            path_depth = self._node[item]['optimal']
            if path_depth < 0:
                path_depth = self._dynamic_cal_optimal(item)
                self._node[item]['optimal'] = path_depth
            if max_path_depth < path_depth:
                max_path_depth = path_depth
        return max_path_depth + 1

    @property
    def optimal(self):
        """最长长度"""
        return self._cal_optimal()

    @property
    def optimal_str(self):
        """最长子句"""
        if not self.is_cal_node:
            assert self._cal_optimal()
        optimal_list = [self._start_node]
        index = 0
        while self._node[optimal_list[index]]['type'] > 0:
            max_optimal = 0
            node_name = None
            for name in self._dict[optimal_list[index]]:
                optimal = self._node[name]['optimal']
                if optimal > max_optimal:
                    max_optimal = optimal
                    node_name = name
            optimal_list.append(node_name)
            index += 1
        return ' '.join(optimal_list)
