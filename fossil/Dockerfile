# https://fossil-scm.org/home/doc/trunk/www/build.wiki
FROM alpine:edge
RUN apk update                                                                       \
  && apk upgrade                                                                     \
  && apk add --no-cache                                                              \
    curl gcc make tcl                                                                \
    musl-dev                                                                         \
    openssl-dev zlib-dev                                                             \
    openssl-libs-static zlib-static                                                  \
  && curl                                                                            \
  "https://fossil-scm.org/home/tarball/fossil-src.tar.gz?name=fossil-src&uuid=trunk" \
  -o fossil-src.tar.gz                                                               \
  && tar xf fossil-src.tar.gz                                                        \
  && cd fossil-src                                                                   \
  && ./configure                                                                     \
    --static                                                                         \
    --disable-fusefs                                                                 \
    --with-th1-docs                                                                  \
    --with-th1-hooks                                                                 \
  && make
