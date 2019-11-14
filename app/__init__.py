
from flask import Flask
from app import config

app = Flask(__name__)
app.config.from_object(config.Config)
cfg = config.Config()
# Flask-login expects app.secret_key instead of config object
app.secret_key = cfg.SECRET_KEY

from flask_login import LoginManager

login_manager = LoginManager(app)
# Iniatialize the Flask login manager.
login_manager.init_app(app)

from app import routes
