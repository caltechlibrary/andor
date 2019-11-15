from flask import render_template, flash, redirect, url_for, request, escape, abort
from flask_login import current_user, login_user, logout_user, login_required
from app import app, cfg, login_manager
from app.forms import LoginForm, PeopleForm, SearchForm
from app.models import User, People
from lunr import lunr
from libdataset import dataset
import time
from datetime import datetime
import sys

@login_manager.user_loader
def load_user(user_id):
    c_name = cfg.USERS
    if dataset.key_exists(c_name, user_id) == False:
        return None
    u = User(user_id)
    return u

# Home page
@app.route('/')
@app.route('/index')
def index():
    return render_template('index.html', title='Home', user = current_user)

# List People
@app.route('/people')
def people_list():
    if current_user.is_authenticated == False:
        flash(f'Must be logged in to curate people')
        return redirect(url_for('index'))
    pg = request.args.get('pg', 1, type=int)
    c_name = cfg.OBJECTS
    keys = dataset.keys(c_name)
    if len(keys) == 0:
        flash(f'No people in {c_name} repository, add some.')
        return redirect(url_for('people_new'))
    keys.sort()
#    if pg > 1:
#        pg -= 1
#    else:
#        pg = 0
#    offset = pg * 10 
#    if len(keys) > 10:
#        keys = keys[offset:offset+10] 
    objects, err = dataset.read_list(c_name, keys)
    if err != '':
        flash(f"Can't read {c_name}, {err}")
        objects = []
    return render_template('people_list.html', title='List People', user = current_user, objects = objects)


# Search People
@app.route('/people/search', methods = [ "GET", "POST" ])
def people_search():
    if current_user.is_authenticated == False:
        flash(f'Must be logged in to curate people')
        return redirect(url_for('index'))
    objects = []
    form = SearchForm()
    if form.validate_on_submit():
        #FIXME: do search here ...
        c_name = cfg.OBJECTS
        keys = dataset.keys(c_name)
        if len(keys) == 0:
            flash(f'No people in {c_name} repository, add some.')
            return redirect(url_for('people_new'))
#        if pg > 1:
#            pg -= 1
#        else:
#            pg = 0
#        offset = pg * 10 
#        if len(keys) > 10:
#            keys = keys[offset:offset+10] 
        objects, err = dataset.read_list(c_name, keys)
        if err != '':
            flash(f"Can't read {c_name}, {err}")
            objects = []
        else:
            # NOTE: we want to save a dict of keys to object
            # for when we want to assemble our results list.
            oMap = dict(zip(keys, objects))
            idx = lunr(
                    ref = "_Key",
                    fields= [ 
                        'cl_people_id', 'family_name',
                        'given_name', 'thesis_id',
                        'authors_id', 'archivesspace_id',
                        'directory_id', 'viaf',
                        'lcnaf', 'isni',
                        'wikidata', 'snac',
                        'orcid', 'image',
                        'educated_at', 'caltech',
                        'jpl', 'faculty',
                        'alumn', 'notes',
                    ],
                    documents = objects)
            results = []
            try:
                results = idx.search(form.query.data)
            except Exception as err:
                flash(f'Search error: {err}')
            objects = []
            for result in results:
                key = result['ref']
                if key in oMap:
                    objects.append(oMap[key])
    return render_template('people_search.html', title='People Search', user = current_user, objects = objects, form = form)




@app.route('/people/new', methods = [ "GET", "POST" ])
def people_new():
    if current_user.is_authenticated == False:
        flash(f'Must be logged in to curate people')
        return redirect(url_for('index'))
    form = PeopleForm()
    if form.validate_on_submit():
        people = People()
        people.cl_people_id = form.cl_people_id.data
        people.family_name = form.family_name.data
        people.given_name = form.given_name.data
        people.thesis_id = form.thesis_id.data
        people.authors_id = form.authors_id.data
        people.archivesspace_id = form.archivesspace_id.data
        people.directory_id = form.directory_id.data
        people.viaf = form.viaf.data
        people.lcnaf = form.lcnaf.data
        people.isni = form.isni.data
        people.wikidata = form.wikidata.data
        people.snac = form.snac.data
        people.orcid = form.orcid.data
        people.image = form.image.data
        people.educated_at = form.educated_at.data
        people.caltech = form.caltech.data
        people.jpl = form.jpl.data
        people.faculty = form.faculty.data
        people.alumn = form.alumn.data
        people.notes = form.notes.data
        c_name = cfg.OBJECTS
        key = people.cl_people_id
        now = datetime.now()
        if dataset.key_exists(c_name, key):
            err = dataset.update(c_name, key, people.to_dict())
            if err != '':
                flash(f'WARNING: failed to update {key} in {c_name}, {err}')
            else:
                flash(f'{people.cl_people_id} updated {now}')
        else:
            err = dataset.create(c_name, key, people.to_dict())
            if err != '':
                flash(f'WARNING: failed to create {key} in {c_name}, {err}')
            else:
                flash(f'{people.cl_people_id} created {now}')
        return redirect(url_for('people_edit', cl_people_id = people.cl_people_id))
    return render_template('people_edit.html', title="New People", user = current_user, form=form)


@app.route('/people/edit/<cl_people_id>', methods = [ "GET", "POST" ])
def people_edit(cl_people_id):
    if current_user.is_authenticated == False:
        flash(f'Must be logged in to curate people')
        return redirect(url_for('index'))
    people = People()
    people.load(cl_people_id)
    form = PeopleForm()
    if form.validate_on_submit():
        people.cl_people_id = form.cl_people_id.data
        people.family_name = form.family_name.data
        people.given_name = form.given_name.data
        people.thesis_id = form.thesis_id.data
        people.authors_id = form.authors_id.data
        people.archivesspace_id = form.archivesspace_id.data
        people.directory_id = form.directory_id.data
        people.viaf = form.viaf.data
        people.lcnaf = form.lcnaf.data
        people.isni = form.isni.data
        people.wikidata = form.wikidata.data
        people.snac = form.snac.data
        people.orcid = form.orcid.data
        people.image = form.image.data
        people.educated_at = form.educated_at.data
        people.caltech = form.caltech.data
        people.jpl = form.jpl.data
        people.faculty = form.faculty.data
        people.alumn = form.alumn.data
        people.notes = form.notes.data
        c_name = cfg.OBJECTS
        key = people.cl_people_id
        now = datetime.now()
        if dataset.key_exists(c_name, key):
            err = dataset.update(c_name, key, people.to_dict())
            if err != '':
                flash(f'WARNING: failed to update {key} in {c_name}, {err}')
            else:
                flash(f'{people.cl_people_id} updated {now}')
        else:
            err = dataset.create(c_name, key, people.to_dict())
            if err != '':
                flash(f'WARNING: failed to create {key} in {c_name}, {err}')
            else:
                flash(f'{people.cl_people_id} created {now}')
    else:
        form.cl_people_id.data = people.cl_people_id 
        form.family_name.data = people.family_name 
        form.given_name.data = people.given_name 
        form.thesis_id.data = people.thesis_id 
        form.authors_id.data = people.authors_id 
        form.archivesspace_id.data = people.archivesspace_id 
        form.directory_id.data = people.directory_id 
        form.viaf.data = people.viaf 
        form.lcnaf.data = people.lcnaf 
        form.isni.data = people.isni 
        form.wikidata.data = people.wikidata 
        form.snac.data = people.snac 
        form.orcid.data = people.orcid 
        form.image.data = people.image 
        form.educated_at.data =  people.educated_at 
        form.caltech.data = people.caltech 
        form.jpl.data = people.jpl 
        form.faculty.data = people.faculty 
        form.alumn.data = people.alumn 
        form.notes.data = people.notes 
    return render_template('people_edit.html', title="Edit Person", user = current_user, form=form)


@app.route('/login', methods = ["GET", "POST"])
def login():
    if current_user.is_authenticated:
        return redirect(url_for('index'))
    form = LoginForm()
    if form.validate_on_submit():
        username = form.username.data
        password = form.password.data
        remember_me = form.remember_me.data
        u = User(username)
        if u.check_password(password) == False:
            flash('Invalid username or password')
            return abort(401)
        login_user(user = u, remember=remember_me, fresh = True)
        #flash('Logged in successfully.')
        return redirect(url_for('people_list'))
    return render_template('login.html', title="Sign in", user = current_user, form=form)


@app.route("/logout")
def logout():
    logout_user()
    return redirect(url_for('index'))
