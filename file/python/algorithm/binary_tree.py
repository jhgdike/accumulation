# coding: utf-8


class SimpleBinaryTrue(object):

    def __init__(self):
        self._root = {}
        self._order = None

    def add(self, val):
        if not isinstance(val, int) and not isinstance(val, float):
            raise(ValueError, 'int or float needed')
        self._add(self._root, val)
        self._order = None

    def _add(self, node, val):
        if not node:
            node.update({'left': {}, 'right': {}, 'val': val, 'count': 1})
            assert node
        elif val < node['val']:
            self._add(node['left'], val)
        elif val > node['val']:
            self._add(node['right'], val)
        else:
            node['count'] += 1

    def delete(self, val):
        if not isinstance(val, int) or not isinstance(val, float):
            raise (ValueError, 'int or float needed')
        self._delete(self._root, val)
        self._order = None

    def _delete(self, node, val):
        if not node:
            raise(ValueError, 'no such value')
        if val == node['val']:
            if node['count']:
                node['count'] -= 1
            else:
                raise(ValueError, 'no such value')
        elif val < node['val']:
            self._delete(node['left'], val)
        elif val > node['val']:
            self._delete(node['right'], val)

    def travel(self):
        self._travel_sal(self._root)
        return self._order

    def _travel_sal(self, node):
        if node:
            self._travel_sal(node['left'])
            if not self._order:
                self._order = []
            if node['count']:
                self._order.append(node['val'])
            self._travel_sal(node['right'])
