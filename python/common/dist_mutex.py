# -*- coding: utf-8 -

import logging
import signal
from contextlib import contextmanager
from datetime import datetime, timedelta
from redis import Redis

from .utils import retry

redis_cli = Redis()


class DistMutex(object):
    """一个简单的分布式锁"""

    def __init__(self, key):
        self._key = 'mutex' + key
        self._expire_time = None

    @property
    def is_locked(self):
        if self._expire_time:
            return self._expire_time > datetime.now()
        return False

    @retry()
    def lock(self, timeout):
        if not self.is_locked:
            set_ok = redis_cli.set(self._key, '', ex=timeout, nx=True)
            if set_ok:
                self._expire_time = datetime.now() + timedelta(seconds=timeout)
                return True

    @retry()
    def unlock(self):
        if self.is_locked:
            redis_cli.delete(self._key)
        self._expire_time = None


def _handler(signum, frame):
    raise Exception("dist_mutex timeout")


@contextmanager
def dist_mutex_context(key, timeout):
    if timeout is None:
        yield False
    mutex = DistMutex(key)
    signal.signal(signal.SIGALRM, _handler)  # 超时强杀
    try:
        mutex.lock(timeout)
        signal.alarm(timeout)  # 设置超时闹钟
        yield mutex.is_locked
    except Exception as e:
        logging.exception(e)
        raise e
    finally:
        signal.alarm(0)  # 取消闹钟
        try:
            mutex.unlock()
        except Exception as e:
            logging.error('dist_mutex_context 解锁异常: %s' % e)


@contextmanager
def frequency_control(key, timeout):
    if timeout is None:
        yield False
    mutex = DistMutex(key)
    try:
        mutex.lock(timeout)
        yield mutex.is_locked
    except Exception as e:
        logging.exception(e)
        yield False
