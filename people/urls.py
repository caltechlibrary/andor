
from django.urls import path
from . import views

urlpatterns = [
    path('', views.index, name=index),
    path('people/', views.People.as_view(), name='people'),
] 


