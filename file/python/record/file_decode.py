# -*- coding: UTF-8 -*-

from __future__ import unicode_literals, absolute_import

BOM_CODE = {
    'BOM_UTF8': 'utf_8',
    'BOM_LE': 'utf_16_le',
    'BOM_BE': 'utf_16_be',
}

DEFAULT_CODES = ['utf8', 'gbk', 'utf16', 'big5']


def decode_file(f_read):
    """解码，使用待定"""
    for k in BOM_CODE:
        if k == f_read[:len(k)]:
            code = BOM_CODE[k]
            f_read = f_read[len(k):]
            text = f_read.decode(code)
            return text.splitlines()

    for encoding in DEFAULT_CODES:
        try:
            text = f_read.decode(encoding)
            return text.splitlines()
        except UnicodeDecodeError:
            continue

    raise Exception('解码失败')

