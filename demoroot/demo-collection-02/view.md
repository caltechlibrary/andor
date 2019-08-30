+++
markup = "mmark"
+++

# View Object


<section>
<div id="object-view"></div>
<p>
<a id="object-json-view">JSON view</a><p>
</section>

<!-- START: View Object -->

<script src="/scripts/CL.js"></script>

<script src="/scripts/andor.js"></script>

<script>
(function (window, document) {
   "use strict";
    var elem = document.getElementById("object-view"),
        anchor = document.getElementById("object-json-view"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        objectID = u.searchParams.get("key");
    if (objectID && anchor !== undefined) {
        anchor.setAttribute("href", `/${c_name}/read/${objectID}`)
    } else {
        anchor.setAttribute("href", `/${c_name}/list.html`)
        anchor.innerHTML = "No object id specified, view List";
    }
    if (elem !== undefined && objectID !== null) {
        AndOr.viewObject(elem, c_name, objectID);
    }
}(window, document));
</script>

<!--   END: View Object -->

