from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import UserMixin
from py_dataset import dataset
from app import cfg, login_manager



def NewUser(username, email, display_name):
    '''NewUser create a new user in cfg.USERS, then returns a new User().'''
    c_name = cfg.USERS
    user = {
        'c_name': c_name,
        'id': username,
        'username': username,
        'display_name': display_name,
        'email': email,
        'role': '',
        'password': ''
    }
    if dataset.has_key(c_name, username): 
        return None
    err = dataset.create(c_name, username, user)
    if err != '':
        return None
    u = User(username)
    return u

#
# User model holds the elements that the app needs to interact with
# user information such as display name, email, password and role.
#
class User(UserMixin):
    '''User is the model of users who can use the And/Or web service'''
    id = ''
    username = ''
    display_name = ''
    email = ''
    role = ''
    password = ''

    def __init__(self, username = ''):
        self.c_name = cfg.USERS
        user = {}
        if username != '':
            user, err = dataset.read(self.c_name, username)
            if err != '':
                print(f'error reading {username} in {cfg.USERS}, {err}')
        self.id = username
        self.username = user['username'] if 'username' in user else ''
        self.display_name = user['display_name'] if 'display_name' in user else ''
        self.email = user['email'] if 'email' in user else ''
        self.role = user['role'] if 'role' in user else ''
        self.password = user['password'] if 'password' in user else ''

    def save(self):
        c_name = self.c_name
        key = self.username
        if dataset.has_key(c_name, key):
            err = dataset.update(c_name, key, self)
            if err != '':
                return False
        else:
            err = dataset.create(c_name, key, self)
            if err != '':
                return False
        return True

    def set_password(self, password):
        self.password = generate_password_hash(password)
        return self.save()

    def check_password(self, password):
        return check_password_hash(self.password, password)

    def __str__(self):
        return self.display_name

    #FIXME: check permissions


#
# Role describes a set of permissions on objects, roles and 
# users collections. A role has a name and that is used to persist
# the role in the cfg.ROLES collection.
# 
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


    def __init__(self, c_name):
        self.c_name = c_name

    def get(self, role_name):
        if dataset.has_key(self.c_name, role_name) == False:
            return False
        role, err = dataset.read(self.c_name, role_name)
        if err != '':
            return False
        self.role_name = role['role_name']
        self.objects = role['objects']
        self.roles = roles['roles']
        self.users = roles['users']
        return True

    def save(self):
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

