# coding: utf-8


class Singleton1(object):

    _singleton = None
    def __new__(cls, *args, **kwargs):
        if not cls._singleton:
            cls._singleton = object.__new__(*args, **kwargs)
        return cls._singleton


class Brog(object):

    _key = {}
    def __new__(cls, key, *args, **kwargs):
        if not hasattr(cls._key, key):
            obj = object.__new__(*args, **kwargs)
            cls._key[key] = obj
        return cls._key[key]


class Singleton2(type):

    def __init__(self, name, bases, class_dict):
        super(Singleton2, self).__init__(name, bases, class_dict)
        self._singleton = None

    def __call__(self, *args, **kwargs):
        if not self._singleton:
            self._singleton = super(Singleton2, self).__call__(*args, **kwargs)
        return self._singleton


class Object2(object):
    __metaclass__ = Singleton2


if __name__ == '__main__':
    """
    process are as follows:
        Object2 = Singleton2(name, bases, class_dict)
        Object2() = Singleton2(name, bases, class_dict)()
                  = Singleton2(name, bases, class_dict).__call__()
        as a result, all instances of Object2 point the same instance Singleton2._singleton
    """
    obj = Object2()
