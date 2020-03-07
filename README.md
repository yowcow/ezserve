ezserve
=======

Easily serve static files in a directory

HOW TO USE
----------

    ezserve [-addr <address>] [-root <directory>] [-quiet]

* `-addr`: Address to bind (default to `:10080`)
* `-root`: Directory to serve (relative to current, default to `.`)
* `-quiet`: Quiet output (default to `false`)

INSTALL
-------

    go get -u -v github.com/yowcow/ezserve
