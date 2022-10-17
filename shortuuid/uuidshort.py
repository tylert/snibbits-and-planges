#!/usr/bin/env python

import uuid as u

import click
import shortuuid


def genv1(node=None, clock_seq=None):
    return u.uuid1(node=node, clock_seq=clock_seq)


def genv3(name=None, namespace=None):
    match namespace.upper():
        case 'DNS':
            return u.uuid3(namespace=u.NAMESPACE_DNS, name=name)
        case 'OID':
            return u.uuid3(namespace=u.NAMESPACE_OID, name=name)
        case 'URL':
            return u.uuid3(namespace=u.NAMESPACE_URL, name=name)
        case 'X500':
            return u.uuid3(namespace=u.NAMESPACE_X500, name=name)
        case _:
            raise ValueError


def genv4():
    return u.uuid4()


def genv5(name=None, namespace=None):
    match namespace.upper():
        case 'DNS':
            return u.uuid5(namespace=u.NAMESPACE_DNS, name=name)
        case 'OID':
            return u.uuid5(namespace=u.NAMESPACE_OID, name=name)
        case 'URL':
            return u.uuid5(namespace=u.NAMESPACE_URL, name=name)
        case 'X500':
            return u.uuid5(namespace=u.NAMESPACE_X500, name=name)
        case _:
            raise ValueError


@click.command()
@click.option(
    '--alphabet',
    '-a',
    help='Alphabet to use to decode/encode a short UUID (default base58)',
    default='123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz',
)
@click.option(
    '--long',
    '-l',
    help='Show the long UUID instead of the short one (default false)',
    default=False,
    is_flag=True,
    show_default=True,  # click insists that the default remains hidden if the default value is false
)
@click.option(
    '--name',
    '-n',
    help='Name to use for the UUIDv5 or v3 hash',
    default='',
)
@click.option(
    '--namespace',
    '-ns',
    help='Namespace to use for the UUIDv5 or v3 hash (DNS, OID, URL, X500)',
    default='DNS',
    show_default=True,
)
@click.option(
    '--typeuuid',
    '-t',
    help='Generate a new UUID of version (type) v5/v4/v3/v2/v1',
    default='4',
    show_default=True,
)
@click.option(
    '--uuid',
    '-u',
    help='Existing UUID to shorten or lengthen',
    default=None,
)
@click.help_option('--help', '-h')
def main(alphabet, long, name, namespace, typeuuid, uuid):
    '''Generate a short UUIDv4 if no parameters specified'''

    if uuid:
        try:
            luu = u.UUID(uuid)
        except ValueError:
            # It might be a short UUID already
            luu = shortuuid.ShortUUID(alphabet=alphabet).decode(string=uuid)
    else:
        # A non-empty name but default type means we probably want UUIDv5
        if name and typeuuid == '4':
            typeuuid = '5'

        match typeuuid:
            # case '1':
            #     luu = genv1(node=node, clock_seq=clock_seq)
            case '3':
                luu = genv3(name=name, namespace=namespace)
            case '4':
                luu = genv4()
            case '5':
                luu = genv5(name=name, namespace=namespace)
            case _:
                raise ValueError

    suu = shortuuid.ShortUUID(alphabet=alphabet).encode(luu)
    if long:
        print(luu)
    else:
        print(suu)


if __name__ == '__main__':
    main()


# https://docs.python.org/3/library/uuid.html
# https://github.com/skorokithakis/shortuuid
# https://pypi.org/project/shortuuid/

# default alphabet '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base57)
# desired alphabet '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base58)
# 122 bits of entropy for UUIDs
