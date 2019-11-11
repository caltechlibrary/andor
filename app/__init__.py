
from flask import Flask
from config import Config

app = Flask(__name__)
app.config.from_object(Config)
cfg = Config()

from flask_login import LoginManager

login = LoginManager(app)

from app import routes
