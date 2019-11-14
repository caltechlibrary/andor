import os
import sys
import shutil
from dotenv import load_dotenv
load_dotenv(dotenv_path=".flaskenv")

if not os.path.exists(os.path.join('app', 'config.py')):
    print(f'Nothing to reset.')
    sys.exit(1)

from app import config
cfg = config.Config()
flask_env = os.getenv('FLASK_ENV') or ''
print(f"cfg.USERS -> {cfg.USERS}")
print(f"cfg.ROLES -> {cfg.ROLES}")
print(f"cfg.OBJECTS -> {cfg.OBJECTS}")
print(f"FLASK_ENV -> {flask_env}")
if flask_env != 'development':
    print(f'''
WARNING: Cannot find FLASK_ENV in the OS environment or
it does not equal "development".

You are not running a development deployment. You need to manaully
remove the following files if you really want to reset and delete
the content of this And/Or repository!

    rm -fR {cfg.USERS} {cfg.ROLES} {cfg.OBJECTS} config.py

''')
    sys.exit(1)
print("Resetting the development environment")
if os.path.exists(cfg.USERS):
    shutil.rmtree(cfg.USERS)
if os.path.exists(cfg.ROLES):
    shutil.rmtree(cfg.ROLES)
if os.path.exists(cfg.OBJECTS):
    shutil.rmtree(cfg.OBJECTS)
if os.path.exists('config.py'):
    os.remove('config.py')
