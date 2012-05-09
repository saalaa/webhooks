webhooks
========


Introduction
------------

This is a simple implementation of a webhooks server that will listen for
queries and try to run anything named after the path of the query it finds in
the `PATH` environment variable.

To start the server:

    $ PATH=hooks ./webhooks


Examples
--------

After starting the server, the following examples provide a good overview of
how to use the server along with the possible queries and responses:

    $ time curl http://localhost:8005/test-ok
    {"code": 200, "message": "OK"}
    real	0m0.007s
    user	0m0.000s
    sys	0m0.004s

    $ time curl http://localhost:8005/test-async
    {"code": 200, "message": "OK"}
    real	0m2.019s
    user	0m0.008s
    sys	0m0.004s

    $ time curl http://localhost:8005/test-async?async=yes
    {"code": 200, "message": "OK"}
    real	0m0.015s
    user	0m0.008s
    sys	0m0.008s

    $ time curl http://localhost:8005/test-error
    {"code": 500, "message": "Internal Server Error"}
    real	0m0.010s
    user	0m0.008s
    sys	0m0.000s

    $ time curl http://localhost:8005/test-missing
    {"code": 501, "message": "Not Implemented"}
    real	0m0.012s
    user	0m0.000s
    sys	0m0.008s


Security Considerations
-----------------------

There are a couple of things you will want to do when using this server:

* Always restrict the `PATH`, either by setting it inline (as in the example)
  or by using `export`
* Make sure you have sufficient rights to run the scripts correctly but
  certainly not *too much*
* Rate limit the frontend web server if hooks come from a high activity SCM
  repository (for example) and/or perform heavy processing
* Run it inside a `chroot` if possible

The server should not be able to run arbitrary programs on your system. Don't
take my word for granted (and neither that of libraries/language used) and
check what the server can and can't do yourself.

I will not assume any responsibility whatsoever.

License
-------

Released under a BSD License.
