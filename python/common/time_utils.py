# encoding: utf8

import math
import time
from datetime import datetime, timedelta

from adlab.common.utils.deprecated import deprecated

DAY_SECONDS = 24 * 60 * 60
HOUR_SECONDS = 60 * 60
MINUTE_SECONDS = 60
DATE_FORMAT = "%Y-%m-%d"
DATETIME_FORMAT = "%Y-%m-%d %H:%M:%S"
CH_MDHm_FORMAT = "%m月%d日%H时%M分"


def format_datetime(datetime_object, time_format=DATETIME_FORMAT):
    if not isinstance(datetime_object, datetime):
        return datetime_object
    return datetime_object.strftime(time_format)


def str_to_datetime(datetime_str, time_format=DATETIME_FORMAT):
    return datetime.strptime(datetime_str, time_format)


def str_to_date(date_str):
    return str_to_datetime(date_str, time_format=DATE_FORMAT)


def _format_chinese_seconds(seconds, current_level):
    if current_level < 1:
        return ""
    current_level -= 1
    if seconds < MINUTE_SECONDS:
        return "%d秒" % math.ceil(seconds)
    elif seconds > DAY_SECONDS:
        days = divmod(seconds, DAY_SECONDS)
        return "%d天%s" % (int(days[0]), format_seconds(days[1], current_level))
    elif seconds > HOUR_SECONDS:
        hours = divmod(seconds, HOUR_SECONDS)
        return '%d小时%s' % (int(hours[0]), format_seconds(hours[1], current_level))
    else:
        mins = divmod(seconds, MINUTE_SECONDS)
        return "%d分钟%s" % (int(mins[0]), format_seconds(mins[1], current_level))


def _format_seconds(seconds, current_level):
    if current_level < 1:
        return ""
    current_level -= 1
    if seconds < MINUTE_SECONDS:
        return "%ds" % math.ceil(seconds)
    elif seconds > DAY_SECONDS:
        days = divmod(seconds, DAY_SECONDS)
        return "%dd %s" % (int(days[0]), format_seconds(days[1], current_level))
    elif seconds > HOUR_SECONDS:
        hours = divmod(seconds, HOUR_SECONDS)
        return '%dh %s' % (int(hours[0]), format_seconds(hours[1], current_level))
    else:
        mins = divmod(seconds, MINUTE_SECONDS)
        return "%dm %s" % (int(mins[0]), format_seconds(mins[1], current_level))


def format_seconds(seconds, max_level=2):
    return _format_seconds(seconds, max_level)


@deprecated
def round_datetime_to_date(dt):
    return dt.replace(hour=0, minute=0, second=0, microsecond=0)


def ceil_datetime_to_hour(dt):
    floor_dt = floor_datetime_to_hour(dt)
    return floor_dt if dt == floor_dt else floor_dt + timedelta(hours=1)


def floor_datetime_to_hour(dt):
    return dt.replace(minute=0, second=0, microsecond=0)


def ceil_timestamp_to_hour(ts):
    floor_ts = floor_timestamp_to_hour(ts)
    return floor_ts if ts == floor_ts else floor_ts + 3600


def floor_timestamp_to_hour(ts):
    return ts - ts % 3600


def today_start_timestamp():
    cur_time = time.time()
    return int(cur_time - cur_time % 86400 - 3600 * 8)


def today_start_time():
    return ts2datetime(today_start_timestamp())


def format_date(date_object):
    if isinstance(date_object, basestring):
        return date_object
    return date_object.strftime(DATE_FORMAT)


def date_str_to_dt(date_str):
    return datetime.strptime(date_str, DATE_FORMAT)


def datetime_str_to_dt(datetime_str):
    return datetime.strptime(datetime_str, DATETIME_FORMAT)


# 如果callback()处理时间超过duration,则会一直循环调用;否则,每duratio时间间隔调用一次.
# callback返回true表示终止执行
# max_seconds表示该函数最多执行多长时间之后,会退出执行
def schedule(duration, callback, max_seconds):
    last_ts = time.time()
    start_ts = last_ts
    while True:
        if callback():
            return
        ts = time.time()
        sleep_ts = duration - (ts - last_ts)
        if sleep_ts > 0:
            time.sleep(sleep_ts)
        if 0 < max_seconds < ts - start_ts:
            return


def datetime_2_timestamp(dt):
    start = datetime(1970, 1, 1, 8)
    if dt < start:
        return 0
    return int((dt - datetime(1970, 1, 1, 8)).total_seconds())


def ts2datetime(ts):
    return datetime.fromtimestamp(ts)


def truncate_date_day(dt):
    return datetime(dt.year, dt.month, dt.day)


def truncate_date_hour(dt):
    return datetime(dt.year, dt.month, dt.day, dt.hour)


def truncate_date_minute(dt):
    return datetime(dt.year, dt.month, dt.day, dt.hour, dt.minute)


def truncate_date_second(dt):
    return datetime(dt.year, dt.month, dt.day, dt.hour, dt.minute, dt.second)


def iso_format(dt):
    if dt:
        return dt.isoformat()
