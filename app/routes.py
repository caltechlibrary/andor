from flask import render_template, flash, redirect, url_for, request, escape, abort
from flask_login import current_user, login_user, login_required
from app import app, cfg, login
from app.forms import LoginForm
from app.models import User
from py_dataset import dataset

@login.user_loader
def user_loader(user_id):
    c_name = cfg.ROLES
    if dataset.has_key(c_name, user_id) == False:
        return None
    u = User(c_name)
    if u.get(user_id) == False:
        return None
    print(f'DEBUG user_loader({user_id}) -> {u.username} {u.display_name}')
    return u

@app.route('/')
@app.route('/index')
def index():
    if current_user.is_authenticated:
        user = {"username": current_user.username, "display_name": current_user.display_name}
    elif current_user.is_anonymous:
        user = {"username": "anonymous", "display_name": "Anonymous"}
    else:
        user = {"username": 'unknown', "display_name": 'unknown'}
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
        return redirect(url_for('index'))
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
            return abort(401)
        login_user(user = u)
        print(f'DEBUG current user is successfully logged in, {u.username}')
        print(f'DEBUG redirecting to /')
        flash('Logged in successfully.')
        return redirect(url_for('index'))
    return render_template('login.html', title="Sign in", user = current_user, form=form)

@app.route("/logout")
@login_required
def logout():
    logout_user()
    return redirect(url_for('index'))
