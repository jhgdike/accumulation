

class SimpleFactory(object):

    def create_product(self, arg):
        if arg == 'a':
            return ConcreateProductA()
        elif arg == 'b':
            return ConcreateProductB()


class Product(object):

    def use(self):
        pass


class ConcreateProductA(Product):

    def use(self):
        print('I am ProductA')


class ConcreateProductB(Product):

    def use(self):
        print('I am ProductB')

if __name__ == '__main__':
    product = SimpleFactory().create_product('a')
    product.use()
    product = SimpleFactory().create_product('b')
    product.use()
