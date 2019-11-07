from django.db import models
#from py_libdataset import dataset

# Create your models here.
class People(models.Model):
    """This class represents our CaltechPEOPLE object which we used to crosswalk data between our systems"""

    # Fields defines the columns in a GSheet
    CL_PEOPLE_ID = models.CharField(verbose_name = "Caltech Library People ID", max_length=256, primary_key = True, help_text = "The common ID used between Caltech Library systems, often the Creator ID from CaltechAUTHORS")
    Family_Name	= models.CharField(max_length=256, help_text = "The family of the individual, e.g. John Steinbeck would be 'Steinbeck'")
    Given_Name	= models.CharField(max_length=256, help_text = "The given name of an individual, e.g. John Steinbeck would be 'John'")
    Thesis_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The creator ID used in CaltechTHESIS")
    Advisor_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The advisor id used in CatechTHESIS")
    Authors_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The creator ID used in CaltechAUTHORS")
    ArchivesSpace_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The agent:person id used in ArchivesSpace")
    Directory_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The directory user id used in the Caltech Directory")
    VIAF_ID = models.CharField(max_length=256, blank = True, default = "", help_text = "The VIAF if avaialble")	
    LCNAF = models.CharField(max_length=256, blank = True, default = "", help_text = "The LCNAF if available")	
    ISNI = models.CharField(max_length=256, blank = True, default = "", help_text = "The ISNI if available")
    Wikidata = models.CharField(max_length=256, blank = True, default = "", help_text = "The wikidata id if available")
    SNAC = models.CharField(max_length=256, blank = True, default = "", help_text = "The SNAC if available")
    ORCID = models.CharField(max_length=256, blank = True, default = "", help_text = "The ORCID if available")	
    image = models.CharField(max_length=256, blank = True, default = "", help_text = "A URL to an image of the individual")
    educated_at	 = models.TextField(blank = True, default = "", help_text = "A free form text description of past education")
    Caltech = models.BooleanField(blank = True, default = "", help_text = "True if associated with Caltech")
    JPL	 = models.BooleanField(default = False, help_text = "True if associated with JPL")
    Faculty = models.BooleanField(default = False, help_text = "True if Caltech faculty")
    Alumn = models.BooleanField(default = False, help_text = "True if Caltech Alumni")
    Notes = models.TextField(blank = True, default = "", help_text = "Internal processing notes")
    Created = models.DateTimeField(auto_now_add = True, help_text = "Date object was created")
    Updated = models.DateTimeField(auto_now = True, help_text = "Date object was updated")

    # Metadata
    class Meta:
        ordering = [ '-family_name', '-given_name' ]

    # Methods
    def get_absolute_url(self):
        """Returns the url to access a particular instance of the model."""
        return reverse('people-detail-view', args=[str(self.CL_PEOPLE_ID)])

    def __str__(self):
        return self.Family_Name + ", " + self.Given_Name


