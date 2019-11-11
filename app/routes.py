from flask import render_template, flash, redirect, url_for, request, escape
from flask_login import current_user, login_user
from app import app, cfg, login
from app.forms import LoginForm
from app.models import User

@login.user_loader
def load_user(user_id):
    c_name = cfg.ROLES
    if dataset.has_key(c_name, user_id) == False:
        return None
    u = User(c_name)
    if u.get(user_id) == False:
        return None
    print(f'DEBUG User loaded {u.username} {u.display_name}')
    return u

@app.route('/')
@app.route('/index')
def index():
    if current_user.is_anonymous:
        user = {"username": "anonymous", "display_name": "Anonymous"}
    else:
        user = {"username": current_user.username, "display_name": current_user.display_name}
    print(f'DEBUG user is {user}')
    posts = [
        {
            'author': {'username': 'John'},
            'title': "John's And/Or repository item",
        },
        {
            'author': {'username': 'Sarah'},
            'title': "Strange moons of Jupiter item",
        }
    ]
    return render_template('index.html', title='Home', user = user, posts = posts)

@app.route('/login', methods = ["GET", "POST"])
def login():
    if current_user.is_authenticated:
        print(f'DEBUG current user is {current_user.username}, redirecting')
        next = request.args.get('next')
        if not is_safe_url(next):
            return flask.about(400)
        return redirect(next or url_for('index'))
    form = LoginForm()
    if form.validate_on_submit():
        c_name = cfg.USERS
        username = form.username.data
        password = form.password.data
        remember_me = form.remember_me.data
        u = User(c_name)
        if u.get(username) == False or u.check_password(password) == False:
            print(f'DEBUG {username} or password not found in {c_name}')
            flash('Invalid username or password')
            return flask.about(400)
        login_user(user = u)
        print(f'DEBUG current user is logged in, {u.username}')
        flash('Logged in successfully.')
        next = request.args.get('next')
        if not is_safe_url(next):
            return flask.about(400)
        return redirect(next or url_for('index'))
    return render_template('login.html', title="Sign in", user = current_user, form=form)

