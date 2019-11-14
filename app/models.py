from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import UserMixin
from libdataset import dataset
from app import cfg, login_manager
from dataclasses import dataclass, field



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
    if dataset.key_exists(c_name, username): 
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
        if dataset.key_exists(c_name, key):
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
        if dataset.key_exists(self.c_name, role_name) == False:
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

class People:
    cl_people_id = ''
    family_name = ''
    given_name = ''
    thesis_id = ''
    authors_id = ''
    archivesspace_id = ''
    directory_id = ''
    viaf = ''
    lcnaf = ''
    isni = ''
    wikidata = ''
    snac = ''
    orcid = ''
    image = ''
    educated_at = ''
    caltech = False
    jpl = False
    faculty = False
    alumn = False
    notes = ''

    def load(self, cl_people_id):
        c_name = cfg.OBJECTS
        if dataset.key_exists(c_name, cl_people_id):
            u, err = dataset.read(c_name, cl_people_id)
            if err != '':
                return err
            self.cl_people_id = u['cl_people_id'] if 'cl_people_id' in u else ''
            self.family_name = u['family_name'] if 'family_name' in u else ''
            self.given_name = u['given_name'] if 'given_name' in u else ''
            self.thesis_id = u['thesis_id'] if 'thesis_id' in u else ''
            self.authors_id = u['authors_id'] if 'authors_id' in u else ''
            self.archivesspace_id = u['archivesspace_id'] if 'archivesspace_id' in u else ''
            self.directory_id = u['directory_id'] if 'directory_id' in u else ''
            self.viaf = u['viaf'] if 'viaf' in u else ''
            self.lcnaf = u['lcnaf'] if 'lcnaf' in u else ''
            self.isni = u['isni'] if 'isni' in u else ''
            self.wikidata = u['wikidata'] if 'wikidata' in u else ''
            self.snac = u['snac'] if 'snac' in u else ''
            self.orcid = u['orcid'] if 'orcid' in u else ''
            self.image = u['image'] if 'image' in u else ''
            self.educated_at = u['educated_at'] if 'educated_at' in u else ''
            self.caltech = u['caltech'] if 'caltech' in u else False
            self.jpl = u['jpl'] if 'jpl' in u else False
            self.faculty = u['faculty'] if 'faculty' in u else False
            self.alumn = u['alumn'] if 'alumn' in u else False
            self.notes = u['notes'] if 'notes' in u else ''
        return ''

    def to_dict(self):
        o = {}
        o['cl_people_id'] = self.cl_people_id
        o['family_name'] = self.family_name
        o['given_name'] = self.given_name
        o['thesis_id'] = self.thesis_id
        o['authors_id'] = self.authors_id
        o['archivesspace_id'] = self.archivesspace_id
        o['directory_id'] = self.directory_id
        o['viaf'] = self.viaf
        o['lcnaf'] = self.lcnaf
        o['isni'] = self.isni
        o['wikidata'] = self.wikidata
        o['snac'] = self.snac
        o['orcid'] = self.orcid
        o['image'] = self.image
        o['educated_at'] = self.educated_at
        o['caltech'] = self.caltech
        o['jpl'] = self.jpl
        o['faculty'] = self.faculty
        o['alumn'] = self.alumn
        o['notes'] = self.notes
        return o


