+++
title = "Collections, Objects and Attachments"
markup = "mmark"
+++


# Collections, Objects and Attachments

Collections hold objects. Objects are expressed as JSON
and stored as JSON documents. Each document must have
a unique ID in the collection.  The minimum object dataset 
supports is one holding a `._Key` attribute and value. 
The `._Key` reflects the unique id in the collection. The 
minimum And/Or object holds a `._Key` and `._State` 
attributes.  The later is used to map an object to a workflow 
and apply the capabilities expressed in And/Or's roles.

## The plumbing

**And/Or** is based on Caltech Library's [dataset](https://caltechlibrary.github.io/dataset) 
tool. __dataset__ is a simple JSON object store organized 
as a directory containing a pairtree[^1].
Each branch of the pair tree contains an object. 
You can think of this is Fedora extra-light[^2]. The dataset 
tool provides some basic CRUD[^3] operations as well as 
a few extended concepts like key lists and data frames. 
Not all of dataset's features have been mapped into And/Or
at this time.  And/Or is intended to be a light weight
web service on top of dataset's collection(s). Where the
dataset tool is single user And/Or provides multi-user support
because it is a web service and using the dataset package as if it 
web server were a single user. In effect it is a proxy
of many users behaving as a single coordated user from the
point of view of the dataset collection. 

A limitation of this approach is operating on a dataset
collection by the dataset CLI while it also is being
use by And/Or is likely to cause problems for everyone
do the lack of lock enforcement.

[^1]: Pairtrees are a way of organizing files on disc by mapping a key to a directory path that avoids too many files in the file directory, see https://confluence.ucop.edu/display/Curation/PairTree

[^2]: Fedora (not to be confused with RedHat), is an enterprise repository services offered by Duraspace, see https://duraspace.org/fedora/

[^3]: CRUD is database jargon for "Create, Read, Update and Delete".

