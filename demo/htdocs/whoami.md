<div id="whoami"></div>
<p>
<a href="access">JSON view</a><p>

<script src="/andor.js"></script>
<script>
(function (window, document) {
    'use strict';
    let c_name = "demo-people",
        elem = document.getElementById("whoami");

    if (elem !== undefined) {
        AndOr.whoAmI(elem, c_name);
    }
}(window, document));
</script>
