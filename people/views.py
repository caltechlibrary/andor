from django.shortcuts import render

# Create your views here.

from django.views import generic

class PeopleView(generic.ListView):
    model = People


