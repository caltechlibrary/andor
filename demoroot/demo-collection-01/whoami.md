+++
markup = "mmark"
+++


# Who am I?


<section>
<div id="whoami"></div>
<p>
<a href="access">JSON view</a><p>
</section>


<!-- START: Who Am I? function -->

<script src="/scripts/CL.js"></script>

<script src="/scripts/andor.js"></script>

<script>
(function (window, document) {
    'use strict';
    let u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        elem = document.getElementById("whoami");

    if (elem !== undefined) {
        AndOr.whoAmI(elem, c_name);
    }
}(window, document));
</script>

<!--   END: Who Am I? function -->

