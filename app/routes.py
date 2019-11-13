from flask import render_template, flash, redirect, url_for, request, escape, abort
from flask_login import current_user, login_user, logout_user, login_required
from app import app, cfg, login_manager
from app.forms import LoginForm
from app.models import User
from py_dataset import dataset

@login_manager.user_loader
def load_user(user_id):
    c_name = cfg.USERS
    if dataset.has_key(c_name, user_id) == False:
        flash(f'DEBUG load_user({user_id}), failed, {user_id} not in {c_name}')
        return None
    u = User(user_id)
    flash(f'DEBUG load_user({user_id}) -> {u.username} {u.display_name} is_authenticated: {u.is_authenticated}, is_active: {u.is_active}, is_anonymous: {u.is_anonymous}')
    return u

@app.route('/')
@app.route('/index')
def index():
#    if current_user.is_authenticated:
#        user = {'username': current_user.username, 'display_name': current_user.display_name, 'is_authenticated' : True}
#    elif current_user.is_anonymous:
#        user = {'username': 'anonymous', 'display_name': 'Anonymous', 'is_authenticated': False}
#    else:
#        user = {'username': current_user.username, 'display_name': current_user.display_name, 'is_authenticated': False}
    posts = [
        {
            'author': {'username': 'John'},
            'title': "John's And/Or repository item",
        },
        {
            'author': {'username': 'Sarah'},
            'title': 'Strange moons of Jupiter item',
        }
    ]
    return render_template('index.html', title='Home', user = current_user, posts = posts)


@app.route('/login', methods = ["GET", "POST"])
def login():
    if current_user.is_authenticated:
        flash(f'DEBUG current user is {current_user.username}, redirecting')
        return redirect(url_for('index'))
    form = LoginForm()
    if form.validate_on_submit():
        username = form.username.data
        password = form.password.data
        remember_me = form.remember_me.data
        u = User(username)
        if u.check_password(password) == False:
            flash(f'DEBUG {u.username} or password not found in {u.c_name}')
            flash('Invalid username or password')
            return abort(401)
        login_user(user = u, remember=remember_me, fresh = True)
        flash(f'DEBUG {current_user} -> current_user.is_authenticated {current_user.is_authenticated}')
        flash(f'DEBUG current user is successfully logged in, {u.username}')
        flash(f'DEBUG redirecting to /')
        flash('Logged in successfully.')
        return redirect(url_for('index'))
    return render_template('login.html', title="Sign in", user = current_user, form=form)

@app.route("/logout")
def logout():
    logout_user()
    return redirect(url_for('index'))
