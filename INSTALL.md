
# Installation

This software is a proof of concept. As such you need to have
go version 1.12 or better to compile it.  Typically the
following steps are done to install (example is Posix based).

```bash
    git clone https://github.com/caltechlibrary/AndOr
    cd AndOr
    go build -o bin/AndOr cmd/AndOr/AndOr.go
```

You can then run the following to see if the tool works.

```bash
    bin/AndOr -help
```

See [Setting Up AndOr](docs/Setting-Up-AndOr.html) for details.

