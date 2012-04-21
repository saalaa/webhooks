webhooks
========

Introduction
------------

This is a simple implementation of a webhooks server that will listen for
queries and try to run scripts named after the path of the query it finds in
the `PATH` environment variable.

To start the server:

    $ PATH=hooks ./webhooks

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

License
-------

Released under a BSD License.

