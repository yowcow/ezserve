ezserve
=======

Easily serve static files in a directory

HOW TO USE
----------

    ezserve [-addr <address>] [-root <directory>] [-quiet]

* `-addr`: Address to bind (default to `:10080`)
* `-root`: Directory to serve (relative to current, default to `.`)
* `-quiet`: Quiet output (default to `false`)
* `-cert`: TLS cert file (default to empty)
* `-key`: TLS key file (default to empty)

INSTALL
-------

    go get -u -v github.com/yowcow/ezserve

HOW TO CREATE KEY/CERT
----------------------

https://gist.github.com/yowcow/8b31fda462fe59f2d9638a7b8e124f4a
