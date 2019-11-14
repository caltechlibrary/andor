#!/usr/bin/env python3
import sys
import os
import string
import secrets
from libdataset import dataset

def generate_development_secret(size = 1024):
    return ''.join(secrets.choice(string.ascii_lowercase + string.ascii_uppercase + string.digits + '.!@#$%^&*()-_+=') for _ in range(size))

cli_name = os.path.basename(sys.argv[0])
if len(sys.argv) != 4:
    print(f"USAGE: {cli_name} USER_COLLECTION_NAME ROLE_COLLECTION_NAME OBJECT_COLLECTION_NAME")
    print("E.g.")
    print(f"{cli_name} Users.ds Roles.ds Objects.ds")
    sys.exit(1)
for c_name in sys.argv[1:]:
    print(f"Initializing {c_name}")
    err = dataset.init(c_name)
    if err != "":
        print(f"Error {err}")
        sys.exit(1)

random_string = generate_development_secret(256)
config_py = f'''

import os

class Config(object):
    SECRET_KEY = os.getenv('SECRET_KEY') or '{random_string}'
    USERS = "{sys.argv[1]}"
    ROLES = "{sys.argv[2]}"
    OBJECTS = "{sys.argv[3]}"

'''

if os.path.exists("app/config.py"):
    print(f'''
Updating app/config.py setting USERS, ROLES, and OBJECTS.
    SECRET_KEY = " ... "
    USERS = "{sys.argv[1]}"
    ROLES = "{sys.argv[2]}"
    OBJECTS = "{sys.argv[3]}"
''')
else:
    print(f'''
Creating app/config.py setting USERS, ROLES, and OBJECTS.
    SECRET_KEY = " ... "
    USERS = "{sys.argv[1]}"
    ROLES = "{sys.argv[2]}"
    OBJECTS = "{sys.argv[3]}"
''')
with open("app/config.py", "w") as fp:
    fp.write(f'''{config_py}''')
