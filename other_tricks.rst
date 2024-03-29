Python Stuff
------------

* http://blog.miguelgrinberg.com/post/designing-a-restful-api-with-python-and-flask
* http://flask-restful.readthedocs.org/en/0.3.5/quickstart.html#a-minimal-api
* http://locust.io/
* https://pypi.python.org/pypi/pip2pi/0.6.8
* https://github.com/pypiserver/pypiserver
* http://book.pythontips.com/en/latest/
* http://qpleple.com/add-progress-bars-to-your-python-loops/
* https://click.palletsprojects.com/en/7.x/utils/#showing-progress-bars
* https://codingdose.info/2019/06/15/how-to-use-a-progress-bar-in-python/
* https://github.com/niltonvolpato/python-progressbar/blob/master/examples.py
* https://bhave.sh/micropython-docker/
* https://til.simonwillison.net/python/stdlib-cli-tools

::

    echo ${JSON_GOOP} | python -m json.tool  # pretty-print
    python -m base64 -h                      # encode/decode base64 strings
    python -m calendar                       # show a full year calendar

::

    import itertools

    for perm in itertools.permutations("thing to use for anagram"):
        print(''.join(perm))

    # Examples:
    # "setec astronomy" is an anagram for "too many secrets"
    # "keynote shogun" is an anagram for "not enough keys"


Bash Stuff
----------

::

    kl5() { echo "${1}"|tr ABC 2|tr DEF 3|tr GHI 4|tr JKL 5|tr MNO 6|tr PQRS 7|tr TUV 8|tr WXYZ 9;}
    # kl5 877-PINK-EYE
    # 877-7465-393


HDWallet Zipapp
---------------

Do the following::

    # Install all pip requirements for the app in an empty directory
    python -m pip install ${MY_PIP_OPTIONS} --target ${MY_APP_DIRECTORY}

    # Cull a bunch of unnecessary gunk here!!!
    rm -rf ${MY_APP_DIRECTORY}/*dist-info
    find ${MY_APP_DIRECTORY} -name '*.py[coz]' | xargs rm
    find ${MY_APP_DIRECTORY} -name 'test*.py' | xargs rm
    for ((i=0; i<10; i++)); do find ${MY_APP_DIRECTORY} -type d -empty | xargs rm -r; done

    # Do any monkey-patching you want here!!!
    # Prepare your top-level script here!!!
    cp ${MY_PYTHON_SCRIPT} ${MY_APP_DIRECTORY}/__main__.py

    # Bundle everything in the directory into an executable zip file
    python -m zipapp -p '/usr/bin/env python' ${MY_APP_DIRECTORY}

Replace all the "sha3.keccak_256()" calls with just "sha3_256()" in
hdwallet/hdwallet.py and hdwallet/libs/base58.py and all "import sha3" lines
with "from hashlib import sha3_256" (assumes you are on Python 3.6+).  This
removes the redundant, non-pure-Python module.

You'll also have to replace all the "mnemonic/wordlist/\*.txt" files with an
equivalent .py file for each language, import them into the appropriate place
at the top of the mnemonic module and do an "eval(language)" and some split()
magic in its init() to get it to spit out the correct stuff.  This is needed
because I haven't yet figured out a good way to read a file packed inside the
zip file that's currently executing.

* https://docs.python.org/3/library/zipapp.html
* https://stackoverflow.com/questions/5355694/python-can-executable-zip-files-include-data-files


Unix Timestamps
---------------

::

    >>> from datetime import datetime
    >>> datetime.fromtimestamp(1649451192)
    datetime.datetime(2022, 4, 8, 16, 53, 12)


IPv4 Calculations
-----------------

::

    >>> import ipaddress
    >>> list(ipaddress.ip_network('163.123.192.190/29', strict=False).hosts())
    [IPv4Address('163.123.192.185'), IPv4Address('163.123.192.186'),
    IPv4Address('163.123.192.187'), IPv4Address('163.123.192.188'),
    IPv4Address('163.123.192.189'), IPv4Address('163.123.192.190')]
    >>> ipaddress.ip_network('163.123.192.190/29', strict=False).num_addresses
    8
    >>> len(list(ipaddress.ip_network('163.123.192.190/29', strict=False).hosts()))
    6
    >>> ipaddress.ip_network('163.123.192.190/29', strict=False).network_address
    IPv4Address('163.123.192.184')
    >>> ipaddress.ip_network('163.123.192.190/29', strict=False).broadcast_address
    IPv4Address('163.123.192.191')
    >>> ipaddress.ip_network('163.123.192.190/29', strict=False).netmask
    IPv4Address('255.255.255.248')
    >>> ipaddress.ip_network('163.123.192.190/29', strict=False).hostmask
    IPv4Address('0.0.0.7')

::

    import ipaddress
    ip1 = int(ipaddress.IPv4Address('10.1.200.202'))
    ip2 = int(ipaddress.IPv4Address('10.1.200.207'))
    print(ip2 - ip1 + 1)  # how many exist between these addresses?


Binary Subnet Mask:	11111111.11111111.11111111.11111000
Binary ID:	10100011011110111100000010111110
Integer ID:	2742796478
Hex ID:	0xa37bc0be
in-addr.arpa:	190.192.123.163.in-addr.arpa
IPv4 Mapped Address:	::ffff:a37b.c0be
6to4 Prefix:	2002:a37b.c0be::/48
