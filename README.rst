emdr-announcer-go
=================

:Status: Production
:Author: Greg Taylor
:License: BSD

This is an EMDR_ announcer written in Go_.

.. _Go: http://golang.org/
.. _EMDR: http://readthedocs.org/docs/eve-market-data-relay/

Install
-------

* Install Go_.
* ``sudo go get github.com/alecthomas/gozmq``
* From within your ``emdr-relay-announcer`` dir: ``go build emdr-relay-announcer.go``
* You should now be able to run the announcer: ``./emdr-announcer-go``

License
-------

This project, and all contributed code, are licensed under the BSD License.
A copy of the BSD License may be found in the repository.
