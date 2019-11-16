from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, BooleanField, SubmitField, TextAreaField, DateTimeField
from wtforms.validators import DataRequired, URL, Optional, Length

class SearchForm(FlaskForm):
    query = StringField('Search', validators = [ DataRequired() ])
    submit = SubmitField('Go')

class LoginForm(FlaskForm):
    username = StringField('Username', validators = [ DataRequired() ])
    password = PasswordField('Password', validators = [ DataRequired() ])
    remember_me = BooleanField('Remember me?')
    submit = SubmitField('Sign in')

class PeopleForm(FlaskForm):
    cl_people_id = StringField('CL People ID', validators = [ DataRequired() ])
    family_name = StringField('Family Name', validators = [ DataRequired() ])
    given_name = StringField('Given Name', validators = [ DataRequired() ])
    thesis_id = StringField('Thesis ID', validators = [])
    authors_id = StringField('Authors ID', validators = [])
    archivesspace_id = StringField('ArchivesSpace ID', validators = [])
    directory_id = StringField('Directory ID', validators = [])
    viaf = StringField('VIAF', validators = [])
    lcnaf = StringField('LCNAF', validators = [Optional()])
    isni = StringField('ISNI', validators = [Optional()])
    wikidata = StringField('wikidata', validators = [Optional()])
    snac = StringField('SNAC', validators = [Optional()])
    orcid = StringField('ORCID', validators = [Optional()])
    image = StringField('Image URL', validators = [URL(), Optional()])
    educated_at = TextAreaField('Educated At', validators = [Optional()])
    caltech = BooleanField('Caltech?')
    jpl = BooleanField('JPL?')
    faculty = BooleanField('Faculty?')
    alumn = BooleanField('Alumni?')
    notes = TextAreaField('Notes (internal)', validators = [Optional()])
    updated = DateTimeField('updated', format='%Y-%m-%d %H:%M:%S')
    submit = SubmitField('Save')

