# coding: utf-8


def tosearch(arr, val):
    left, right = 0, len(arr)
    mid = (left + right) / 2
    while True:
        if arr[mid] > arr:
            right = mid
        elif arr[mid] < arr:
            left = mid
        else:
            return mid
        mid = (left + right) / 2
    return -1
