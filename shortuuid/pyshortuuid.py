#!/usr/bin/env python


import uuid as u

import click
import shortuuid
import zbase32


def gen_uuidv1(node: str = None, clock_seq: str = None) -> str:
    ''' '''
    n = None
    c = None
    if node is not None:
        n = int(node)
    if clock_seq is not None:
        c = int(clock_seq)
    return u.uuid1(node=n, clock_seq=c)


def gen_uuidv2(
    node: str = None, clock_seq: str = None, domain: str = 'PERSON', id: int = 0
) -> str:
    ''' '''
    # https://stackoverflow.com/questions/20910653/how-to-shift-bits-in-a-2-5-byte-long-bytes-object-in-python
    # https://github.com/google/UUID/blob/v1.0.0/version1.go  UUIDv1
    # https://github.com/google/UUID/blob/v1.0.0/dce.go       UUIDv2 is built upon UUIDv1
    x = gen_uuidv1(node=node, clock_seq=clock_seq)
    match domain.upper():
        case 'PERSON':
            uu = u.UUID(bytes=x.bytes)
            return uu
        case 'GROUP':
            uu = u.UUID(bytes=x.bytes)
            return uu
        case 'ORG':
            uu = u.UUID(bytes=x.bytes)
            return uu
        case _:
            raise ValueError


def gen_uuidv3(namespace: str = 'DNS', name: str = '') -> str:
    ''' '''
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


def gen_uuidv4() -> str:
    ''' '''
    return u.uuid4()


def gen_uuidv5(namespace: str = 'DNS', name: str = '') -> str:
    ''' '''
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


def gen_uuidv6(node: str = None, clock_seq: str = None) -> str:
    ''' '''
    # uuid1_to_uuid6() from https://github.com/oittaa/uuid6-python/blob/main/src/uuid6/__init__.py
    uu = gen_uuidv1(node=node, clock_seq=clock_seq)
    h = uu.hex
    h = h[13:16] + h[8:12] + h[0:5] + '6' + h[5:8] + h[16:]
    return u.UUID(hex=h, is_safe=uu.is_safe)


# def gen_uuidv7() -> str:
#     ''' '''
#     return u.uuid7()


# def gen_uuidv8() -> str:
#     ''' '''
#     return u.uuid8()


@click.command()
@click.help_option('--help', '-h')
@click.option(
    '--alphabet',
    '-a',
    help='Alphabet to use to decode/encode a short UUID (default base58)',
    default='123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz',
)
@click.option(
    '--clock_seq',
    '-c',
    help='Clock sequence [14-bit number] to use for UUIDv6/v2/v1',
    default=None,
)
@click.option(
    '--domain',
    '-d',
    help='Domain [PERSON/GROUP/ORG] to use for UUIDv2',
    default='PERSON',
)
@click.option(
    '--encoding',
    '-e',
    help='Encoding [BASE58/ZBASE32/NONE] to use for shortening UUID',
    default='BASE58',
)
@click.option(
    '--id',
    '-i',
    help='ID to use for UUIDv2',
    default='0',
)
@click.option(
    '--name',
    '-n',
    help='Name to use for UUIDv5/v3',
    default='',
)
@click.option(
    '--namespace',
    '-ns',
    help='Namespace [DNS/OID/URL/X500] to use for UUIDv5/v3',
    default='DNS',
    show_default=True,
)
@click.option(
    '--node',
    '-o',
    help='Node [48-bit MAC] to use for UUIDv6/v2/v1',
    default=None,
)
@click.option(
    '--type',
    '-t',
    help='Type (version) [UUIDv6/v5/v4/v3/v2/v1] of UUID to generate',
    default='4',
    show_default=True,
)
@click.option(
    '--uuid',
    '-u',
    help='Existing UUID to shorten or lengthen',
    default=None,
)
def main(
    alphabet, clock_seq, domain, encoding, id, name, namespace, node, type, uuid
) -> None:
    '''Generate a shortened (encoded) UUID'''

    if uuid:
        try:
            luu = u.UUID(uuid)
        except ValueError:
            # It might be a short UUID already
            luu = shortuuid.ShortUUID(alphabet=alphabet).decode(string=uuid)
    else:
        # A non-empty name but default type means we probably want UUIDv5
        if name and type == '4':
            type = '5'

        match type.upper():
            case '1':
                luu = gen_uuidv1(node=node, clock_seq=clock_seq)
            case '2':
                luu = gen_uuidv2(node=node, clock_seq=clock_seq, domain=domain, id=id)
            case '3':
                luu = gen_uuidv3(namespace=namespace, name=name)
            case '4':
                luu = gen_uuidv4()
            case '5':
                luu = gen_uuidv5(namespace=namespace, name=name)
            case '6':
                luu = gen_uuidv6(node=node, clock_seq=clock_seq)
            # case '7':
            #     luu = gen_uuidv7()
            # case '8':
            #     luu = gen_uuidv8()
            case _:
                raise ValueError

    match encoding.upper():
        case 'BASE58':
            suu = shortuuid.ShortUUID(alphabet=alphabet).encode(luu)
            print(suu)
        case 'ZBASE32':
            suu = zbase32.encode(luu.bytes)
            print(suu)
        case 'NONE':
            print(luu)
        case _:
            raise ValueError


if __name__ == '__main__':
    main()


# https://github.com/python/cpython/blob/main/Lib/uuid.py
# https://docs.python.org/3/library/uuid.html
# https://docs.python.org/3/library/typing.html

# https://github.com/skorokithakis/shortuuid
# https://pypi.org/project/shortuuid/

# default alphabet '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base57)
# desired alphabet '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base58)
# 122 bits of entropy for UUIDs
