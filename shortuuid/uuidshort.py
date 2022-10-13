#!/usr/bin/env python

import uuid as u

import click
import shortuuid


def genv3(name, namespace, alphabet):
    match namespace.upper():
        case 'DNS':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid3(namespace=u.NAMESPACE_DNS, name=name)
            )
        case 'OID':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid3(namespace=u.NAMESPACE_OID, name=name)
            )
        case 'URL':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid3(namespace=u.NAMESPACE_URL, name=name)
            )
        case 'X500':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid3(namespace=u.NAMESPACE_X500, name=name)
            )
        case _:
            raise ValueError


def genv4(alphabet):
    return shortuuid.ShortUUID(alphabet=alphabet).encode(uuid=u.uuid4())


def genv5(name, namespace, alphabet):
    match namespace.upper():
        case 'DNS':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid5(namespace=u.NAMESPACE_DNS, name=name)
            )
        case 'OID':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid5(namespace=u.NAMESPACE_OID, name=name)
            )
        case 'URL':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid5(namespace=u.NAMESPACE_URL, name=name)
            )
        case 'X500':
            return shortuuid.ShortUUID(alphabet=alphabet).encode(
                uuid=u.uuid5(namespace=u.NAMESPACE_X500, name=name)
            )
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
    '--decode',
    '-d',
    help='Convert (decode) a short UUID to a long UUID',
    default=None,
)
@click.option(
    '--name',
    '-n',
    help='Generate a UUIDv5 using a specified name instead of a UUIDv4',
    default=None,
)
@click.option(
    '--namespace',
    '-ns',
    help='Namespace to use for the UUIDv5 name (default "DNS")',
    default='DNS',
)
@click.option(
    '--typeuuid',
    '-t',
    help='Generate a new UUID of version (type) v5/v4/v3/v2/v1 (default "4")',
    default='4',
)
@click.option(
    '--uuid',
    '-u',
    help='Existing UUID to shorten or lengthen',
    default=None,
)
@click.help_option('--help', '-h')
def main(alphabet, decode, name, namespace, typeuuid, uuid):
    '''Generate a short UUIDv4 if no parameters specified'''

    # default alphabet '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base57)
    # desired alphabet '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base58)
    # 122 bits of entropy for UUIDs

    if decode:
        print(shortuuid.ShortUUID(alphabet=alphabet).decode(string=decode))
    elif uuid:
        print(shortuuid.ShortUUID(alphabet=alphabet).encode(uuid=u.UUID(uuid)))
    elif name:
        match typeuuid:
            case '3':
                print(genv3(name=name, namespace=namespace, alphabet=alphabet))
            case '5':
                print(genv5(name=name, namespace=namespace, alphabet=alphabet))
            case _:
                print(genv5(name=name, namespace=namespace, alphabet=alphabet))
    else:
        print(genv4(alphabet=alphabet))


if __name__ == '__main__':
    main()


# https://docs.python.org/3/library/uuid.html  python uuid standard library
# https://github.com/skorokithakis/shortuuid  python implementation
# https://pypi.org/project/shortuuid/  python implementation
