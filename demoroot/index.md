+++
markup = "mmark"
+++


# Welcome to [And/Or](https://github.com/caltechlibrary/andor)

This is a demo of And/Or[^1], a web base multi-user version of 
[dataset](https://github.com/caltechlibrary/dataset). It is a web
service designed to build curation tools for collections of JSON
objects.

[demo-collection-01.ds](/demo-collection-01/)
: demonstrates the basic user, role, object state, and impact on viewing

[demo-collection-02.ds](/demo-collection-02/)
: demonstrates CRUD operations on a repository

Each demo collection featuers a Web UI created with static
files. Aside from the And/Or JSON object service which provides
content the rest of the UI is constructed from HTML, CSS, and JavaScript.
Each collection can have it's own UI as the name of the collection is
mapped to the static host content for the specific collection
(e.g `/demo-collection-01/`).
Common functionality is defined via JavaScript and implemented in a
file called `/andor.js` which contains a root object calls `AndOr`. 

The demo configuration creates the following users - ester, innez, 
jane, bea and millie. They use a common password "hello". If running
demo under BasicAUTH then use a private browser window to switch
between logins.


[^1]: And/Or is named after a character in the [Ruby](https://www.zbs.org/index_new.php/store/ruby) stories produced by [ZBS](https://www.zbs.org) 
