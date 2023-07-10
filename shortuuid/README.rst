::

    $ ./shortuuid -u PxhKe6exuntRFbEqgNqaVe
    PxhKe6exuntRFbEqgNqaVe
    $ ./shortuuid -u b9efce8a-592d-47d8-8e5b-7ad8660dc7d1
    PxhKe6exuntRFbEqgNqaVe
    $ ./shortuuid -u PxhKe6exuntRFbEqgNqaVe -e none
    b9efce8a-592d-47d8-8e5b-7ad8660dc7d1
    $ ./shortuuid -u b9efce8a-592d-47d8-8e5b-7ad8660dc7d1 -e none
    b9efce8a-592d-47d8-8e5b-7ad8660dc7d1

Python bug!!! (weird leading 1 when first byte is 0???)::

    $ ./pyshortuuid.py -u 05ef55cd-f0e4-4676-852a-765a6c384fcc
    1jWNAus869vomZgM7VmzT5
    $ ./pyshortuuid.py -u 1jWNAus869vomZgM7VmzT5 -e none
    05ef55cd-f0e4-4676-852a-765a6c384fcc
    $ ./pyshortuuid.py -u jWNAus869vomZgM7VmzT5 -e none
    05ef55cd-f0e4-4676-852a-765a6c384fcc
    $ ./shortuuid -u jWNAus869vomZgM7VmzT5 -e none
    05ef55cd-f0e4-4676-852a-765a6c384fcc
    $ ./shortuuid -u 1jWNAus869vomZgM7VmzT5 -e none
    00000000-0000-0000-0000-000000000000

* https://docs.crunchybridge.com/api-concepts/eid/
* https://brandur.org/fragments/base32-slugs
* https://github.com/taskcluster/slugid-go
* https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
* https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-07.html
* https://blog.devgenius.io/analyzing-new-unique-identifier-formats-uuidv6-uuidv7-and-uuidv8-d6cc5cd7391a
* https://medium.com/geekculture/the-wild-world-of-unique-identifiers-uuid-ulid-etc-17cfb2a38fce

::

    # Dump info about a binary that was already built
    go version -m shortuuid
    go tool objdump shortuuid

    # Show all possible GOOS/GOARCH combos for builds
    go tool dist list

    # Show help for linker options for builds
    go tool link
