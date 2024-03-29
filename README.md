And/Or
=====================================================

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg?style=flat-square)](https://choosealicense.com/licenses/bsd-3-clause)

<!-- [![Latest release](https://img.shields.io/badge/Latest_release-0.0.1-b44e88.svg?style=flat-square)](http://shields.io) -->



**And/Or** is a proof of concept simple object repository based on 
Caltech Library's [dataset](https://caltechlibrary.github.io/dataset "And/Or requires v0.1.x")
tool.  It implements an web service for curating a simple object
repository using Python and a custom go implementation of a C-Shared
library - `libdataset`.


Table of contents
-----------------

* [Introduction](#introduction)
* [Installation](install.html)
* [Documentation](docs/)
* [Known issues and limitations](#known-issues-and-limitations)
* [Getting help](#getting-help)
* [Contributing](contributing.html)
* [License](#license)
* [Authors and history](#authors-and-history)
* [Acknowledgments](#authors-and-acknowledgments)


Introduction
------------

**And/Or** is an extremely light weight object repository. 
It builds on [dataset](https://github.com/caltechlibrary/dataset).
It uses Python's Flask package for hosting a web service providing
an asynchronous interface for curating a dataset collection.  
__dataset__ has proven to be a useful tool for managing library 
metadata using a data science approach.  It is built for continious 
migration dataflows.  It lacks multi-user curation support or the 
convienence of having web browser based edit forms. **And/Or** is a 
prototype for extending libdataset.go with a Python managed
service that can then be used to create extremely light weight 
object repository system using Python.


Installation
------------

See [INSTALL.md](install.html). This software is experimental.
There are no pre-compiled binaries provided. This software 
is largely written in Python 3.7 and packages such
as Flask, FlaskWTF as well as Go 1.13 for libdataset. The prototype 
was developed using [Miniconda]() based Python distribution.

[INSTALL.md](install.html) provides details for compiling, 
installing and configing the prototype.


Known issues and limitations
----------------------------

This is a proof-of-concept project. It SHOULD NOT be used
in any production setting.  It is ONLY suitable for demonstrating
an approach to building light weight object repositories.

Getting help
------------

You can contact us via GitHub [issue tracker](https://github.com/caltechlibrary/andor/issues).

Contributing
------------

See [CONTRIBUTING.md](contributing.html)


License
-------

Software produced by the Caltech Library is Copyright (C) 2019, Caltech.  
This software is freely distributed under a BSD/MIT type license.  
Please see the [LICENSE](LICENSE) file for more information.


Authors and history
---------------------------

[Robert](https://rsdoiel.github.io) is the culprit responsible 
for this proof of concept


Acknowledgments
---------------

This work was funded by the California Institute of Technology Library.

<div align="center">
  <br>
  <a href="https://www.caltech.edu">
    <img width="100" height="100" src="assets/caltech-round.svg">
  </a>
</div>

