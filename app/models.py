from werkzeug.security import generate_password_hash, check_password_hash
from py_dataset import dataset

#
# User model holds the elements that the app needs to interact with
# user information such as display name, email, password and role.
#
class User:
    c_name = ''
    username = ''
    display_name = ''
    email = ''
    password = ''
    role = ''

    def __init__(c_name):
        self.c_name = c_name

    def get(username):
        if dataset.key_exists(self.c_name, username):
            return False
        user, err = dataset.read(self.c_name, username)
        if err != '':
            return False
        self.username = user['username']
        self.display_name = user['display_name']
        self.email = user['email']
        self.role = user['role']
        self.password = user['password']
        return True

    def save():
        c_name = self.c_name
        key = self.username
        user = {
            "username": self.username,
            "display_name": self.display_name,
            "email": self.email,
            "role": self.role,
            "password": self.password,
        }
        err = dataset.update(c_name, key, user)
        if err != '':
            return False
        return True

    def set_password(self, username, password):
        ok = self.get(username)
        if ok == False:
            return False
        self.password = generate_password_hash(password)
        return self.save()

    def check_password(self, password):
        c_name = self.c_name
        key = self.username
        if dataset.key_exists(c_name, key) == False:
            return False
        ok = self.get(key)
        if ok == False:
            return False
        return check_password_hash(self.password, password)

    def __str__(self):
        return self.display_name

    #FIXME: check permissions


class Role:
    c_name = ''
    role_name = ''
    objects = {
            "create": False,
            "delete": False,
            "read": False,
            "update": False
        }
    roles = {
            "create": False,
            "delete": False,
            "read": False,
            "update": False
        }
    users = {
            "create": False,
            "delete": False,
            "read": False,
            "update": False
        }


    def __init__(c_name):
        self.c_name = c_name

    def get(role_name):
        if dataset.key_exists(self.c_name, role_name):
            return False
        role, err = dataset.read(self.c_name, role_name)
        if err != '':
            return False
        self.role_name = role['role_name']
        self.objects = role['objects']
        self.roles = roles['roles']
        self.users = roles['users']
        return True

    def save():
        c_name = self.c_name
        key = self.role_name
        role = {
            "role_name": self.role_name,
            "objects": self.objects,
            "roles": self.roles,
            "users": self.users,
        }
        err = dataset.update(c_name, key, user)
        if err != '':
            return False
        return True

    def __str__(self):
        return self.role_name

