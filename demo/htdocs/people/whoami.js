(function (window, document) {
    'use strict';
    let AndOr = window.AndOr,
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        elem = document.getElementById("whoami");

    /**
     * whoAmI takes a DOM element and renders a textual description
     * of a person's role and object assignment options.
     *
     * @param elem is a DOM element
     * @param collection_name is a string holding the collection name
     */
    function whoAmI (elem, collection_name) {
        let div = document.createElement("div"),
            req = new XMLHttpRequest(),
            u = new URL(window.location.href);

        u.pathname = "/" + collection_name + "/access/";
       
        console.log("DEBUG URL", u);
        elem.appendChild(div)
        CL.httpGet(u, 'application/json', function(data, err) {
            let o = data,
                user = o.user,
                roles = o.roles,
                assignments = [],
                blocks = [];
            div.innerHTML = "";
            for (let key in roles) {
                if (roles[key].assign_to !== undefined) {
                    assignments = assignments.concat(roles[key].assign_to);
                }
            }
            for (let i = 0; i < assignments.length; i++) {
                if (assignments[i] == "*") {
                    assignments[i] = "any queue";
                }
            }
            div.innerHTML = `<p>Welcome <em>${user.display_name}</em>,
<p>
You are a member of the following roles
and queues&mdash; 
<ul>${user.roles.length == 0 ? "none" : "<li>" + user.roles.join("</li><li>")+ "</li>"}</ul>
<p>
You may assign objects to&mdash;
<ul>${assignments.length == 0 ? "none": "<li>" + assignments.join("</li>\n<li>") + "</li>"}</ul>`;
        });
        req.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
    }

    if (elem !== undefined) {
        whoAmI(elem, c_name);
    }
}(window, document));
