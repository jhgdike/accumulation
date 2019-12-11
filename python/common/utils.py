# -*- coding: utf-8 -

from functools import wraps
import logging
import math


def retry(tries=3):
    def retry_wrapper(func):
        @wraps(func)
        def retry_func(*args, **kwargs):
            _tries = tries
            while _tries > 0:
                try:
                    ret = func(*args, **kwargs)
                except Exception as err:
                    logging.exception(err)
                    _tries -= 1
                    if not _tries:
                        raise err
                else:
                    return ret

        return retry_func

    return retry_wrapper


def calc_total_page(total_number, page_size):
    return int(math.ceil(1.0 * total_number / max(page_size, 1)))


def calc_page_info(total_number, page, page_size):
    return {
        'page': page,
        'page_size': page_size,
        'total_number': total_number,
        'total_page': calc_total_page(total_number, page_size),
    }
