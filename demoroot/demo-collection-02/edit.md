+++
markup = "mmark"
+++


## Edit


<section>
<div id="object-edit">
<form id="edit-object">
    <div>
        <label>Key</label>
        <input name="key" type="text" value="" placeholder="Object key">
    </div>
    <div>
        <label>Title</label>
        <input name="title" type="text" value="" placeholder="title">
    </div>
    <div>
        <label>Pub Date</label>
        <input name="pubDate" type="date" value="" placeholder="pubDate">
    </div>
    <div>
        Creators Widget would go here ...
    </div>
    <div>
        <label>Description</label><br />
        <textarea name="description" placeholder="description goes here"></textarea>
    </div>
    <div>
        <input type="submit" value="Update"> 
        <input type="reset" value="Clear">
    </div>
</form>
</div>
<p>
<a id="object-json-view">JSON view</a><p>
</section>

<!-- START: Edit Object Form -->

<script src="/scripts/CL.js"></script>

<script src="/scripts/andor.js"></script>

<script>
(function (window, document) {
   "use strict";
    var 
        elem = document.getElementById("object-edit"),
        anchor = document.getElementById("object-json-view"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        objectID = u.searchParams.get("key");

    if (objectID !== null && anchor !== undefined) {
        anchor.setAttribute("href", `/${c_name}/read/${objectID}`)
    } else {
        anchor.setAttribute("href", `/${c_name}/create/${objectID}`)
        anchor.innerHTML = "No object id specified, view List";
    }
    if (obejctID !== null && elem !== undefined) {
        Andor.editObject(elem, objectID);
    }
}(window, document));
</script>

<!--   END: Edit Object Form -->

