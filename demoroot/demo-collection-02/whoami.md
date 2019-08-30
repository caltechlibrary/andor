+++
markup = "mmark"
+++


## Who am I?


<section>
<div id="whoami"></div>
<p>
<a href="access">JSON view</a><p>
</section>

<!-- START: Who am I? -->

<script src="/scripts//CL.js"></script>

<script src="/scripts/andor.js"></script>

<script>
(function (document) {
   "use strict";
    var elem = document.getElementById("whoami"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname);

    if (elem !== undefined) {
        AndOr.whoAmI(elem, c_name);
    }
}(document));
</script>

<!--   END: Who am I? -->

