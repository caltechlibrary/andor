from flask import render_template
from app import app

@app.route('/')
@app.route('/index')
def index():
    user = { 'username' : 'rsdoiel' }
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

@app.route('/login')
def login():
    form = LoginForm()
    return render_template('login.html', title="Sign in", form=form)
