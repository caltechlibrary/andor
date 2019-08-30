+++
markup = "mmark"
+++


## List Objects

<section>
<ul id="object-list"></ul>
</section>

<!-- START: List Objects -->

<script src="/scripts/CL.js"></script>

<script src="/scripts/andor.js"></script>

<script>
(function (window, document) {
   "use strict";
    var elem = document.getElementById("object-list"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        state = u.searchParams.get("state");

    if (elem !== undefined && state !== null) {
        AndOr.viewAllObjects(elem, c_name, state);
    } else {
        AndOr.viewAllObjects(elem, c_name);
    }
}(window, document));
</script>

<!--   END: List Objects -->

