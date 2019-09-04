(function (window, document) {
   "use strict";
    let elem = document.getElementById("object-list"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        keys = u.searchParams.get("keys"),
        state = u.searchParams.get("state");

    if (elem !== undefined) {
        if (keys !== null && keys.includes(",")) {
           AndOr.viewObjectList(elem, c_name, keys.split(","));
        } else if (state !== undefined && state !== null && state !== "") {
           AndOr.viewAllObjects(elem, c_name, state);
        } else {
           AndOr.viewAllObjects(elem, c_name);
        }
    }
}(window, document));

