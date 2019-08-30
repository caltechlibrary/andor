
# List View


<div id="object-list"></div>

<!-- START: List people based on GET request in URL -->
<script src="/scripts/CL.js"></script>
<script src="/scripts/andor.js"></script>
<script>
(function (window, document) {
   "use strict";
    let elem = document.getElementById("object-list"),
        u = new URL(window.location.href),
        c_name = andor.getCollectionName(u.pathname),
        keys = u.searchParams.get("keys"),
        state = u.searchParams.get("state");

    if (elem !== undefined) {
        if (keys !== null && keys.includes(",")) {
           AndOr.viewObjectList(elem, c_name, keys.split(","));
        } else if (state !== null && state !== "") {
           AndOr.viewAllObjects(elem, c_name, state);
        } else {
           AndOr.viewAllObjects(elem, c_name);
        }
    }
}(window, document));
</script>

<!--   END: List people based on GET request in URL -->

