

class FactoryBase(object):

    def create_productA(self):
        pass

    def create_productB(self):
        pass


class FactoryA(FactoryBase):

    def create_productA(self):
        return ConcreateProductA1()

    def create_productB(self):
        return ConcreateProductB1()


class FactoryB(FactoryBase):

    def create_productA(self):
        return ConcreateProductA2()

    def create_productB(self):
        return ConcreateProductB2()


class ProductBaseA(object):

    def use(self):
        pass


class ProductBaseB(object):

    def eat(self):
        pass


class ConcreateProductA1(ProductBaseA):

    def use(self):
        print('I use ProductA1')


class ConcreateProductA2(ProductBaseA):

    def use(self):
        print('I use ProductA2')


class ConcreateProductB1(ProductBaseB):
    def eat(self):
        print('I eat ProductB1')


class ConcreateProductB2(ProductBaseB):
    def eat(self):
        print('I eat ProductB2')

if __name__ == '__main__':
    factory = FactoryA()
    pa = factory.create_productA()
    pb = factory.create_productB()
    pa.use()
    pb.eat()
    factory = FactoryB()
    pa = factory.create_productA()
    pb = factory.create_productB()
    pa.use()
    pb.eat()
