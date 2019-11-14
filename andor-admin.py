import string
import os
import sys
import shutil
import getpass
from libdataset import dataset
from dotenv import load_dotenv
load_dotenv(dotenv_path=".flaskenv")
from werkzeug.security import generate_password_hash
from app import models

if not os.path.exists('config.py'):
    print(f'Nothing to administor.')
    sys.exit(1)

from config import Config
cli_name = os.path.basename(sys.argv[0])
cfg = Config()
flask_env = os.getenv('FLASK_ENV') or ''

def display_help(argv = {}):
    print(f'''
USAGE: {cli_name} VERB [VERB_OPTIONS]

{cli_name} provides the primary means of managing an
And/Or repository. It is used to managed users, assign roles
reset passwords, etc.

Verbs:

    help         display this help message
    add-user     add a new users to the system
    disable-user disable a user account
    email        set a user's email address
    display-name set a user's display name
    password     set a user's password
    assign-role  assign or update a user's role
    create-role  defines a role
    edit-role    changes a role's permissions
    delete-role  deletes a role

Verbs except for help require one or more parameters.
Envoking the verb without a parameter will display a
usage statement for the verb.

''')

def usage_add_user():
    print(f'''
USAGE: {cli_name} add-user USERNAME EMAIL DISPLAY_NAME

E.g. {cli_name} add-user jsteinbeck "jsteinbeck@localhost" "John Steinbeck"

''')

def add_user(argv):
    c_name = cfg.USERS
    if len(argv) != 3:
        usage_add_user()
        return False
    username, email, display_name = argv[0], argv[1], argv[2]
    if dataset.key_exists(c_name, username) == True:
        print(f'{username} already exists in {c_name}')
        return False
    user = {
        'id': username,
        'username': username,
        'email': email,
        'display_name': display_name,
    }
    err = dataset.create(c_name, username, user)
    if err != '':
        return False
    print(f'NOTE: {username} will need to have a role and password set.')
    return True


def usage_disable_user():
    print(f'''
USAGE: {cli_name} disable-user USERNAME

E.g. {cli_name} disable-user jsteinbeck

''')

def disable_user(argv):
    c_name = cfg.USERS
    if len(argv) != 1:
        usage_disable_user()
        return False
    username = argv[0]
    if dataset.key_exists(c_name, username) == False:
        print(f'{username} not found in {c_name}')
        return False
    u = models.User(c_name)
    if u.get(username) == False:
        print(f'{username} not found in {c_name}')
        return False
    u.is_active = False
    return u.save()




def usage_password():
    print(f'''
USAGE: {cli_name} password USERNAME

You will be prompted to enter a password for the user, it will
optionally send an email to the user with the new password.

E.g. {cli_name} password admin

''')

def set_password(argv):
    c_name = cfg.USERS
    if len(argv) < 1:
        return usage_password()
    username = argv[0]
    pw1, pw2 = '', ' '
    i = 0
    while pw1 != pw2:
        pw1 = getpass.getpass(
            prompt = f'Please enter a new password for {username}: ')
        pw2 = getpass.getpass(
            prompt = f'Please enter the password a second time: ')
        if pw1 != pw2:
            i += 1
        if i > 3:
            print(f'Passwords do not match, exiting')
            return False
    u, err = dataset.read(c_name, username)
    if err != '':
        print(f'{username} not found in {c_name}, {err}')
        return False
    u['password'] = generate_password_hash(pw1)
    err = dataset.update(c_name, username, u)
    if err != '':
        print(f'{username} in {c_name}, {err}')
        return False
    return True

def usage_set_email():
    print(f'''
USAGE: {cli_name} email USERNAME EMAIL_ADDRESS

This will set a user's email address.

E.g. {cli_name} email 'admin' 'root@localhost'

''')

def set_email(argv):
    c_name = cfg.USERS
    if len(argv) != 2:
        return usage_set_email()
    username, email = argv[0], argv[1]
    if dataset.key_exists(c_name, username) == True:
        u = models.User(c_name)
        if u.get(username):
            u.email = email
            return u.save()
        return True
    print(f'{username} does not exist in {c_name}')
    return False

def usage_set_display_name():
    print(f'''
USAGE: {cli_name} display-name USERNAME DISPLAY_NAME

This will set a user's display name.

E.g. {cli_name} display-name 'admin' 'Repository Administrator'

''')


def set_display_name(argv):
    c_name = cfg.USERS
    if len(argv) != 2:
        return usage_set_email()
    username, display_name = argv[0], argv[1]
    if dataset.key_exists(c_name, username) == True:
        u = models.User(c_name)
        if u.get(username):
            u.display_name = display_name
            return u.save()
        return True
    print(f'{username} does not exist in {c_name}')
    return False


def usage_assign_role():
    print(f'''
USAGE: {cli_name} assign-role USERNAME ROLE_NAME

This will assign a user to a role. 

E.g. {cli_name} assign-role admin Admin

''')

def assign_role(argv):
    c_name = cfg.USERS
    if len(argv) != 2:
        return usage_asign_role()
    username, role = argv[0], argv[1]
    if dataset.key_exists(c_name, username) == False:
        print(f'{username} does not exist in {c_name}')
        return False
    u, err = dataset.read(c_name, username)
    if err != '':
        print(f'Failed to read {username} from {c_name}, {err}')
        return False
    u['role'] = role
    err = dataset.update(c_name, username, u)
    if err != '':
        print(f'Failed to write {username} to {c_name}, {err}')
        return False
    return True


def usage_create_role():
    print(f'''
USAGE {cli_name} create-role ROLE_NAME

Provides an interactive set of prompts to create a new role
in Andor. Roles had a name and then a set of CRUD values 
for access to the user, role and objects collections.

E.g. {cli_name} create-role Publisher

''')

def create_role(argv):
    c_name = cfg.ROLES
    if len(argv) != 1:
        usage_create_role()
        return False
    role_name = argv[0]
    if dataset.key_exists(c_name, role_name):
        print(f'{role_name} already exists in {c_name}')
        return False
    role = { 
            f'{cfg.USERS}': { 'create': False, 'read': False, 'update': False, 'delete': False},
            f'{cfg.ROLES}': { 'create': False, 'read': False, 'update': False, 'delete': False},
            f'{cfg.OBJECTS}': { 'create': False, 'read': False, 'update': False, 'delete': False},
            }
    for c_name in [ cfg.USERS, cfg.ROLES, cfg.OBJECTS ]:
        print(f'Collection {c_name}')
        perms = role[c_name]
        for perm in [ 'create', 'read', 'update', 'delete' ]:
            y_or_n = input(f'allow {perm}? [Y/n] ').lower()
            if y_or_n in [ 'y', 'yes' ]:
                role[c_name][perm] = True
            else:
                role[c_name][perm] = False
    c_name = cfg.ROLES
    err = dataset.create(c_name, role_name, role)
    if err != '':
        print(f'ERROR: {err}')
        return False
    return True

def usage_edit_role():
    print(f'''
USAGE {cli_name} edit-role ROLE_NAME

Provides an interactive prompt to edit an existing role.

E.g. {cli_name} edit-role Publisher
''')

def edit_role(argv):
    c_name = cfg.ROLES
    if len(argv) != 1:
        usage_create_role()
        return False
    role_name = argv[0]
    if not dataset.key_exists(c_name, role_name):
        print(f'{role_name} does not exist in {c_name}')
        return False
    role, err = dataset.read(c_name, role_name)
    if err != '':
        print(f'ERROR: {err}')
        return False
    for c_name in [ cfg.USERS, cfg.ROLES, cfg.OBJECTS ]:
        print(f'Collection {c_name}')
        perms = role[c_name]
        for perm in [ 'create', 'read', 'update', 'delete' ]:
            val = role[c_name][perm]
            t_or_f = input(f'{perm} is {val}, T(rue)/F(alse)/K(eep)? ').lower()
            if t_or_f in [ 't', 'true' ]:
                role[c_name][perm] = True
            if t_or_f in [ 'f', 'false' ]:
                role[c_name][perm] = False
    c_name = cfg.ROLES
    err = dataset.update(c_name, role_name, role)
    if err != '':
        print(f'ERROR: {err}')
        return False
    return True

def usage_delete_role():
    print(f'''
USAGE {cli_name} delete-role ROLE_NAME

Removes a role from the system.

E.g. {cli_name} delete-role Publisher
''')

def delete_role(argv):
    c_name = cfg.ROLES
    if len(argv) != 1:
        usage_delete_role()
        return False
    role_name = argv[0]
    if dataset.key_exists(c_name, role_name):
        err = dataset.delete(c_name, role_name)
        if err != '':
            return False
        return True
    print(f'{role_name} not found in {c_name}')
    return False

#
# Main cli logic
#
verbs = {
    "help": display_help,
    "add-user": add_user,
    "disable-user": disable_user,
    "email": set_email,
    "display-name": set_display_name,
    "password": set_password,
    "assign-role": assign_role,
    "create-role": create_role,
    "edit-role": edit_role,
    "delete-role": delete_role,
}

if __name__ == '__main__':
    if len(sys.argv) < 2:
        verb = "help"
        display_help()
        sys.exit(0)
    else:
        verb = sys.argv[1]
    
    if verb in verbs:
        ok = verbs[verb](sys.argv[2:])
        if ok != True:
            sys.exit(1)
    else:
        display_help()
        sys.exit(1)
else:
    print(f'Run tests not implemented.')
