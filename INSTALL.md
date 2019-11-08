
# Installation

This software is a proof of concept. 

## Prequisites

Python v3.7 or better installed, Flask 1.1.x, and `py_dataset` v0.1.x 
installed.  The example installation steps use 
[Miniconda](https://docs.conda.io/en/latest/miniconda.html "Miniconda installation page") version of Python 3.  It also assumes a POSIX like
system supported by Miniconda.

```bash
    pip install python-dotenv
    pip install flask
    pip install flask_wtf
```

## Installing And/Or

Use git to clone the And/Or repository. Checkout the pythonic-flask branch.
Create a ".flaskenv" file to hold your flask environment.

```bash
    git clone https://github.com/caltechlibrary/andor
    cd andor
    git checkout pythonic-flask
    vi .flaskenv # Add FLASK_APP=andor.py  
                 # Add FLASK_ENV=development
```

## Preparing the base repostiories

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
