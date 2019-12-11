from threading import _Timer


class RepeatingTimer(_Timer):
    def run(self):
        while not self.finished.is_set():
            self.function(*self.args, **self.kwargs)
            self.finished.wait(self.interval)


def test():
    def hello():
        print("hello, world")

    t = RepeatingTimer(10.0, hello)
    t.start()
