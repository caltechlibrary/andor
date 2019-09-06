
# Installation

This software is a proof of concept. As such you need to have
Go version 1.12 or better to compile it and then to manually
install it.  Typically the following steps are done to install 
(example is POSIX system based).

```bash
    git clone https://github.com/caltechlibrary/andor
    cd andor
    go build -o bin/andor cmd/andor/andor.go
```

You can then run the following to see if the tool works.

```bash
    bin/andor -help
```

See [Setting Up AndOr](docs/Setting-Up-AndOr.html) for details.

