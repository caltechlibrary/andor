
# Installation

This software is a proof of concept. 

## Prequisites

Python v3.7 or better installed, Flask 1.1.x, and `py_dataset` v0.1.x installed.  
The example installation steps use [Miniconda](https://docs.conda.io/en/latest/miniconda.html "Miniconda installation page") version of Python 3.  It also assumes a POSIX like
system supported by Miniconda.

```bash
    pip install python-dotenv
    pip install flask
    pip install flask_wtf
```

## Installing And/Or Source

Use git to clone the And/Or repository. Checkout the pythonic-flask branch.
Create a ".flaskenv" file to hold your flask environment.

```bash
    git clone https://github.com/caltechlibrary/andor
    cd andor
    git checkout pythonic-flask
    vi .flaskenv # Add FLASH_APP=andor.py to your .flaskenv python environment file.
```

## Preparing the base repostiories

The first program, `andor-repositories.py` will create the necessary dataset
collections as well as generate an appropriate `config.py` needed to run the
application.

```bash
    python3 andor-repositories.py Users.ds Roles.ds People.ds
    python3 andor-add-user.py Users.ds Admin # Add an admin user
    python3 andor-assign-role.py Roles.ds Admin # Assign an admin role to the Admin user.
```

## Running the web app

```bash
    flask run
```
