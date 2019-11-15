
# Installation

This software is a proof of concept. 

## Prequisites

Go 1.13 and dataset v0.1.x to compile `libdataset`, 
Python v3.7 or better installed, Flask 1.1.x, 
Flask\_WTF, Flash-Login, and 
[Lunr.py](https://lunr.readthedocs.io/en/latest/).
The example installation steps use 
[Miniconda](https://docs.conda.io/en/latest/miniconda.html "Miniconda installation page") version of Python 3.x.  It also assumes a POSIX like
system supported by Miniconda.

### Installing Go and libdataset

Go to https://golang.org/dl/ and follow download
and installation instructions for Go 1.13 or better.

Download source for dataset, checkout the "andor" branch.

```bash
    git clone https://github.com/caletchlibrary/dataset \
        src/github.com/caltechlibrary/dataset
    cd src/github.com/caltechlibrary/dataset
    git checkout andor
```

### Installing Python modules with Minicona

```bash
    pip install python-dotenv
    pip install flask
    pip install flask_wtf
    pip install flask-login
    pip install lunr[languages]
```

## Installing And/Or

Use git to clone the And/Or repository. Checkout the pythonic-flask branch.
Create a ".flaskenv" file to hold your flask environment.

```bash
    git clone https://github.com/caltechlibrary/andor \
        src/github.com/caltechlibrary/andor
    cd src/github.com/caltechlibrary/andor
    vi .flaskenv # Add FLASK_APP=andor.py  
                 # Add FLASK_ENV=development
    make clean
    cd libdataset
    make build
    cd ..
    make
```

At this point if all has gone well make will
run `andor-setup.py`.


## Preparing and Configuring the People repository

The first program, `andor-setup.py` will create the necessary dataset
collections as well as generate an appropriate `config.py` needed to run the
application.

```bash
    # Create your repositories and config.py file.
    python3 andor-setup.py Users.ds Roles.ds People.ds
    # Create an "Admin" role, answer y to all the questions.
    python3 andor-admin.py create-role 'Admin'
    # Add a user using andor-admin.py
    # add-user will prompt for a password.
    python3 andor-admin.py add-user 'admin' 'jane.doe@example.edu' 'Jane Doe'
    # Assign admin to the Admin role
    python3 andor-amdin.py assign-role 'admin' 'Admin'
    # Set the admin password
    python3 andor-admin.py password 'admin'
```

Now you should be ready to start your And/Or repository service and start
adding and currating your objects.

## Running the And/Or service

```bash
    flask run
```
