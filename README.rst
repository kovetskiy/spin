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

 See ``spin --help`` for command line options on test runners.

  `-i --stdin-as-indicator`   Use stdin data as progress indicator.
  -o --write-stdin          Write stdin to spinner's stdout on exit.
  -s --status <string>      Use specified <string> as spinner status.
  -t --interval <millisec>  Use specified <millisec> as spinner iteration interval.
                            [default: 100]
  -h --help                 Show this screen.
  --version                 Show version.

License
=======

MIT.
