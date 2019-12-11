# coding: utf-8
"""
用法：
pool = ThreadPool(job_number)
pool.add_task(function, *args, **kwargs)
pool.wait_complete()
"""
import logging
from queue import Queue
from threading import Thread


class Consumer(Thread):
    """调度消费者"""

    def __init__(self, queue, err_que):
        Thread.__init__(self)
        self._q = queue
        self.err_que = err_que
        self.daemon = True
        self.start()

    def run(self):
        while True:
            f, args, kwargs = self._q.get()
            try:
                if not f(*args, **kwargs):
                    self.err_que.put(Exception("task failed"))
            except Exception as e:
                logging.exception("params: function:{}, args:{}, kwargs:{}. exception `{}`".format(f, args, kwargs, e))
                self.err_que.put(e)
            self._q.task_done()


class ThreadPool(object):
    def __init__(self, num_t):
        self._q = Queue(num_t)
        self._err_que = Queue()
        for _ in range(num_t):
            Consumer(self._q, self._err_que)

    def add_task(self, f, *args, **kwargs):
        self._q.put((f, args, kwargs))

    def wait_complete(self):
        self._q.join()

    @property
    def err_count(self):
        return self._err_que.qsize()

    def clear(self):
        self._err_que = Queue()
