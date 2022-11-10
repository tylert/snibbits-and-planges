#!/usr/bin/env python

import uuid as u

import click
import shortuuid


def gen_uuidv1(node: str = None, clock_seq: str = None) -> str:
    ''' '''
    return u.uuid1(node=node, clock_seq=clock_seq)


def gen_uuidv2(
    node: str = None, clock_seq: str = None, domain: str = None, id: int = 0
) -> str:
    ''' '''
    # https://github.com/google/UUID/blob/v1.0.0/version1.go  UUIDv1
    # https://github.com/google/UUID/blob/v1.0.0/dce.go       UUIDv2 is built upon UUIDv1
    match domain.upper():
        case 'PERSON':
            uuid = u.uuid1(node=node, clock_seq=clock_seq)
            return uuid
        case 'GROUP':
            uuid = u.uuid1(node=node, clock_seq=clock_seq)
            return uuid
        case 'ORG':
            uuid = u.uuid1(node=node, clock_seq=clock_seq)
            return uuid
        case _:
            raise ValueError


def gen_uuidv3(name: str = None, namespace: str = None) -> str:
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


def gen_uuidv5(name: str = None, namespace: str = None) -> str:
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


# def gen_uuidv6() -> str:
#     ''' '''
#     return u.uuid6()


# def gen_uuidv7() -> str:
#     ''' '''
#     return u.uuid7()


# def gen_uuidv8() -> str:
#     ''' '''
#     return u.uuid8()


@click.command()
@click.option(
    '--alphabet',
    '-a',
    help='Alphabet to use to decode/encode a short UUID (default base58)',
    default='123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz',
)
@click.option(
    '--clock_seq',
    '-c',
    help='Clock sequence to use for UUIDv2/v1 sequence number',
    default=None,
)
@click.option(
    '--encoding',
    '-e',
    help='Encoding to use for shortening UUID - BASE58/NONE',
    default='BASE58',
)
@click.option(
    '--name',
    '-n',
    help='Name to use for the UUIDv5/v3 hash',
    default='',
)
@click.option(
    '--namespace',
    '-ns',
    help='Namespace to use for UUIDv5/v3 hash - DNS/OID/URL/X500',
    default='DNS',
    show_default=True,
)
@click.option(
    '--node',
    '-o',
    help='NodeID [interface name] to use for UUIDv2/v1 MAC - RANDOM/eth0/etc.',
    default=None,
)
@click.option(
    '--typeuuid',
    '-t',
    help='Version [type] of UUID to generate - UUIDv5/v4/v3/v2/v1',
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
def main(alphabet, clock_seq, encoding, name, namespace, node, typeuuid, uuid):
    '''Generate a shortened (encoded) UUID'''

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

        match typeuuid.upper():
            case '1':
                luu = gen_uuidv1(node=node, clock_seq=clock_seq)
            # case '2':
            #     luu = gen_uuidv2(node=node, clock_seq=clock_seq, domain=domain, id=id)
            case '3':
                luu = gen_uuidv3(name=name, namespace=namespace)
            case '4':
                luu = gen_uuidv4()
            case '5':
                luu = gen_uuidv5(name=name, namespace=namespace)
            # case '6':
            #     luu = gen_uuidv6()
            # case '7':
            #     luu = gen_uuidv7()
            # case '8':
            #     luu = gen_uuidv8()
            case _:
                raise ValueError

    suu = shortuuid.ShortUUID(alphabet=alphabet).encode(luu)
    match encoding.upper():
        case 'BASE58':
            print(suu)
        case 'NONE':
            print(luu)
        case _:
            raise ValueError


if __name__ == '__main__':
    main()


# https://docs.python.org/3/library/uuid.html
# https://docs.python.org/3/library/typing.html
# https://github.com/skorokithakis/shortuuid
# https://pypi.org/project/shortuuid/
# https://en.wikipedia.org/wiki/Universally_unique_identifier
# https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-04  UUIDv6/v7/v8
# https://click.palletsprojects.com/en/8.1.x/

# default alphabet '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base57)
# desired alphabet '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base58)
# 122 bits of entropy for UUIDs
