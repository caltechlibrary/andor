+++
markup = "mmark"
+++


# Welcome to [And/Or](https://github.com/caltechlibrary/andor)

This is a demo of And/Or[^1], a JSON object web service based on
[dataset](https://github.com/caltechlibrary/dataset) collections. 
The demo dataset collection is "demo-colloction-01.ds". Web UI 
is built from static HTML/CSS/JavaScripts making http calls back
to the web service.  Web UI is located under `/demo-collection-01/`.

## Web UI

+ [Landing page](/demo-collection-01/) - should trigger a Basic AUTH loging
+ JSON views (raw And/Or access)
    + List of [Keys](/demo-collection-01/keys/) - a JSON list of object keys
    + Single [Object](/demo-collection-01/read/100) - a JSON view of read object id "100".
    + List of [Objects](/demo-collection-01/read/100,101,102,103) - An ordered array of objects with keys 100, 101, 102 and 103
+ HTML views (static pages using JavaScript to format raw output)
    + [List objects keys](/demo-collection-01/keys)
    + [Object](/demo-collection-01/view.html?id=100) - an HTML view of object id "100".
    + [Objects](/demo-collection-01/view.html?id=100,101,102,103) - an HTML view of object id "100", "101", "102", and "103".


[^1]: And/Or is named after a character in [Ruby](https://www.zbs.org/index_new.php/store/ruby) stories produced by [ZBS](https://www.zbs.org) 
