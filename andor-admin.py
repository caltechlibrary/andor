import os
import sys
import shutil
from dotenv import load_dotenv
load_dotenv(dotenv_path=".flaskenv")

if not os.path.exists('config.py'):
    print(f'Nothing to administor.')
    sys.exit(1)

from config import Config
cli_name = os.path.basename(sys.argv[0])
cfg = Config()
flask_env = os.getenv('FLASK_ENV') or ''
#print(f"cfg.USERS -> {cfg.USERS}")
#print(f"cfg.ROLES -> {cfg.ROLES}")
#print(f"cfg.OBJECTS -> {cfg.OBJECTS}")
#print(f"FLASK_ENV -> {flask_env}")

def display_help(params = {}):
    print(f'''
USAGE: {cli_name} VERB [VERB_OPTIONS]

{cli_name} provides the primary means of managing an
And/Or repository. It is used to managed users, assign roles
reset passwords, etc.

Verbs:

    help        display this help message
    add-user    add a new users to the system
    password    set a users password
    assign-role assign or update a roles role
    create-role defines a role
    edit-role   changes a role's permissions
    delete-role deletes a role

Verbs except for help require one or more parameters.
Envoking the verb without a parameter will display a
usage statement for the verb.

''')

def usage_add_user():
    print(f'''
USAGE: {cli_name} add-user USERNAME EMAIL DISPLAY_NAME

E.g. {cli_name} add-ser jsteinbeck "jsteinbeck@localhost" "John Steinbeck"

''')

def add_user(argv):
    print(f'DEBUG in add_user({argv})')
    if len(argv) != 3:
        usage_add_user()
        return False
    username, email, display_name = argv[0], argv[1], argv[2]
    print(f'DEBUG Adding user {username}, email {email}, display name {display_name}')
    print(f'DEBUG add_user() not implemented.')
    return False


def usage_password():
    print(f'''
USAGE: {cli_name} password USERNAME

You will be prompted to enter a password for the user, it will
optionally send an email to the user with the new password.

E.g. {cli_name} password admin

''')

def set_password(argv):
    if len(params) < 1:
        return usage_password()
    username = argv[0]
    print(f'DEBUG set password for username {username}')
    print(f'DEBUG set_password() not implemented.')
    return False


def usage_assign_role():
    print(f'''
USAGE: {cli_name} assign-role USERNAME ROLE_NAME

This will assign a user to a role. 

E.g. {cli_name} assign-role admin Admin

''')

def assign_role(argv):
    if len(params) != 2:
        return usage_asign_role()
    username, role = argv[0], argv[1]
    print(f'DEBUG assign-role {username} to {role}')
    print(f'DEBUG assign_role() not implemented.')
    return False

def usage_create_role():
    print(f'''
USAGE {cli_name} create-role ROLE_NAME

Provides an interactive set of prompts to create a new role
in Andor. Roles had a name and then a set of CRUD values 
for access to the user, role and objects collections.

E.g. {cli_name} create-role Publisher

''')

def create_role(argv):
    if len(argv) != 1:
        usage_create_role()
        return False
    print(f'DEBUG create_role() not implemented.')
    return False

def usage_edit_role():
    print(f'''
USAGE {cli_name} edit-role ROLE_NAME

Provides an interactive prompt to edit an existing role.

E.g. {cli_name} edit-role Publisher
''')

def edit_role(argv):
    if len(argv) != 1:
        usage_edit_role()
        return False
    print(f'DEBUG edit_role() not implemented.')
    return False

def usage_delete_role():
    print(f'''
USAGE {cli_name} delete-role ROLE_NAME

Removes a role from the system.

E.g. {cli_name} delete-role Publisher
''')

def delete_role(argv):
    if len(argv) != 1:
        usage_delete_role()
        return False
    print(f'DEBUG delete_role() not implemented.')
    return False

#
# Main cli logic
#
verbs = {
    "help": display_help,
    "add-user": add_user,
    "password": set_password,
    "assign-role": assign_role,
    "create-role": create_role,
    "edit-role": edit_role,
    "delete-role": delete_role,
}

if __name__ == __main__:
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
