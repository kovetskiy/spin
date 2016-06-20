****
spin
****

.. image:: http://i.imgur.com/JyfyJg9.gif
   :alt: usage example

The universal tool that provides the dead simple progress indicator.

Installation
============

**spin** is go-gettable::

 go get github.com/kovetskiy/spin

Usage
=====

See ``spin --help`` for command line options.

``-i --stdin-as-indicator``   Use stdin data as progress indicator.

``-o --write-stdin``          Write stdin to spinner's stdout on exit.

``-s --status <string>``      Use specified ``<string>`` as spinner status.

``-t --interval <millisec>``  Use specified ``<millisec>`` as spinner iteration interval. [default: 100]

All of your need it's just pass any stdin to spin like as following::

 sleep 2 | spin -s 'Loading... '

And your shell will spawn ``spin`` process and terminate it when ``sleep`` exited.

If you want to indicate real progress you can use flag ``-i`` and spin will
use stdin data as progress indicator::

 git clone --progress https://github.com/kovetskiy/dotfiles 2>&1 | spin -i -s 'Clonning... '

License
=======

MIT.
