

class FactoryBase(object):

    def create_product(self):
        pass


class FactoryA(object):

    def create_product(self):
        return ConcreateProductA()


class FactoryB(object):

    def create_product(self):
        return ConcreateProductB()


class ProductBase(object):

    def use(self):
        pass


class ConcreateProductA(ProductBase):

    def use(self):
        print('I am ProductA')


class ConcreateProductB(ProductBase):

    def use(self):
        print('I am ProductB')

if __name__ == '__main__':
    product = FactoryA().create_product()
    product.use()
    product = FactoryB().create_product()
    product.use()
