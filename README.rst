=======
Polyenv
=======


.. image:: https://img.shields.io/pypi/v/polyenv.svg
        :target: https://pypi.python.org/pypi/polyenv

.. image:: https://img.shields.io/travis/FrankKair/polyenv.svg
        :target: https://travis-ci.org/FrankKair/polyenv


Polyenv allows you to run any programming language.


* Free software: MIT license
* Documentation: https://polyenv.readthedocs.io.





Installation
------------
`$ pip3 install polyenv`


Usage
-----

`$ polyenv language path_to_file`

Calling a `Python 3` script

.. code:: python

  polyenv python3 ~/Desktop/example.py
    "Hello, from Python"


Calling a `Ruby` script

.. code:: python

  polyenv ruby ~/Desktop/example.py
    "Hello, from Ruby"



Features
--------

* TODO


SSL Error
---------

If you come across the following error:

.. code:: bash

    ssl.SSLError: [SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed (_ssl.c:777)

Just run
.. code:: bash

    $ /Applications/Python\ 3.6/Install\ Certificates.command


Credits
-------

This package was created with Cookiecutter_ and the `audreyr/cookiecutter-pypackage`_ project template.

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`audreyr/cookiecutter-pypackage`: https://github.com/audreyr/cookiecutter-pypackage
