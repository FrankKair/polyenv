=======
Polyenv
=======

.. image:: https://img.shields.io/pypi/v/polyenv.svg
        :target: https://pypi.python.org/pypi/polyenv

.. image:: https://img.shields.io/travis/FrankKair/polyenv.svg
        :target: https://travis-ci.org/FrankKair/polyenv

Polyenv allows you to run any programming language.

Installation
------------

.. code:: sh
  
  $ pip3 install polyenv


Usage
-----

.. code:: sh

  $ polyenv language path_to_file


Calling `Python 3`:

.. code:: sh

  $ polyenv python3 ~/Desktop/example.py
  "Hello, from Python"


Calling `Ruby`:

.. code:: sh

  $ polyenv ruby ~/Desktop/example.rb
  "Hello, from Ruby"


SSL Error
---------

If you come across the following error:

.. code:: bash

    ssl.SSLError: [SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed (_ssl.c:777)

Just run:

.. code:: bash

    $ /Applications/Python\ 3.6/Install\ Certificates.command
