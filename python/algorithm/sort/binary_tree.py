# coding: utf-8


class SimpleBinaryTree(object):
    """
    简单的二叉树,用字典生成,拥有添加,删除,搜索和遍历功能
    e.g.
        bt = SimpleBinaryTree()
        bt.add(20)
        bt.add(10)
        bt.add(30)
        bt.add(15)
        bt._root
        = >
            {
                'count': 1,
                'left': {
                    'count': 1,
                    'left': {},
                    'right': {
                        'count': 1,
                        'left': {},
                        'right': {},
                        'val': 15
                        },
                    'val': 10
                    },
                'right': {
                    'count': 1,
                    'left': {},
                    'right': {},
                    'val': 30
                    },
                'val': 20
            }

    """

    def __init__(self, unique=False):
        self._root = {}
        self._order = None
        self.unique = unique

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
            if self.unique:
                raise(ValueError, '{} exists'.format(val))
            node['count'] += 1

    def search(self, val):
        pass

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
                if self.unique:
                    node['count'] = 0
                else:
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


class BinaryTree(object):
    """
    e.g.
        root.val
        root.left - > node
        root.right - > node
        root.count
    """
    left = None
    right = None
    val = None
    count = 0

    def __init__(self, unique=False):
        self.unique = unique

    def add(self, val):
        if not isinstance(val, int) and not isinstance(val, float):
            raise (ValueError, 'int or float needed')
        self._add(val)

    def _add(self, val):
        if not self.val:
            self.val = val
        elif val > self.val:
            if not self.right:
                self.right = BinaryTree()
            self.right.add(val)
        elif val < self.val:
            if not self.left:
                self.left = BinaryTree()
            self.left.add(val)
        else:
            if self.unique:
                raise(ValueError, '{} exists'.format(val))
            self.count += 1

    def search(self, val):
        self._search(self, val)

    def _search(self, node, val):
        if not node:
            return False
        if node.val == val:
            return True
        elif node.val < val:
            return self._search(node.right, val)
        return self._search(node.left, val)

    def delete(self, val):
        pass

    def travel(self):
        order = []
        if self.val:
            if self.left:
                order.extend(self.left.travel())
            order.append(self.val)
            if self.right:
                order.extend(self.right.travel())
        return order
