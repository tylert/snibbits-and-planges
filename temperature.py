#!/usr/bin/env python


def degc(degf: float = None) -> float:
    return ((degf - 32) * 5) / 9


def degf(degc: float = None) -> float:
    return ((degc * 9) / 5) + 32


if __name__ == '__main__':
    print(degc(98.6))
    print(degc(37.0))
    print(degf(98.6))
    print(degf(37.0))
