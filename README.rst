carbon-table
============

.. image:: https://travis-ci.org/yunstanford/carbon-table.svg?branch=master
    :alt: build status
    :target: https://travis-ci.org/yunstanford/carbon-table

.. image:: https://codecov.io/gh/yunstanford/carbon-table/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/yunstanford/carbon-table


Carbon-Table backends with Gin and Trie Tree that supports fast resolving Graphite-Like wildcards query.
This is designed to put behind ``carbon-relay-ng`` and regsiter ``sendAllMatch`` route. That way, Graphite-Web
doesn't need to send requests to all Carbon-Cache Instances for wildcard queries.


.. figure:: ./example/arch.png
   :align: center
   :alt: arch.png


-------
Install
-------

.. code::

    curl https://raw.githubusercontent.com/yunstanford/carbon-table/master/get-carbon-table.sh | sh


----------
Quickstart
----------

.. code::

    # http:8080, tcp:3000
    ./carbon-table

    # send metric data
    echo "test.foo 7 `date +%s`" | nc localhost 3000
    echo "test.bar 7 `date +%s`" | nc localhost 3000

    # expand
    curl http://localhost:8080/metric/pattern/test.*/

    [{"Query":"test.bar","IsLeaf":true},{"Query":"test.foo","IsLeaf":true}]

    # expand
    curl http://localhost:8080/metric/query/test.*/

    ["test.bar","test.foo"]