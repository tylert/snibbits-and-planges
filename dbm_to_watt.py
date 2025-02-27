#!/usr/bin/env python


from math import log10


def dbm(watt: float = None) -> float:
    return 10*log10(1000*watt)


def watt(dbm: float = None) -> float:
    return 10**(dbm/10)/1000


if __name__ == '__main__':
    print(dbm(20.0))
    print(watt(43.0))
    print(dbm(43.0))
    print(watt(20.0))
