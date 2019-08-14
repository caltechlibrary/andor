/**
 * andor.js provides a common browser object for working with the
 * And/Or web service.
 *
 * @author R. S. Doiel, <rsdoiel@library.caltech.edu>
 */
(function (window, document) {
    'use strict';
    let AndOr = new Object();

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
            url = "/" + collection_name + "/access/";
       
        elem.appendChild(div)
        req.addEventListener("load", function(evt) {
            let o = JSON.parse(req.response),
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
        req.open("GET", url);
        req.send();
    }

    /**
     * viewObjectList takes a DOM element and renders a view
     * of the keys provided.
     *
     * @param elem is a DOM element
     * @param collection_name is a string holding the collection name
     * @param keys is a list of object keys in the collection
     * @param filter is a dataset type key filter.
     */
    function viewObjectList(elem, collection_name, keys, label) {
        let req = new XMLHttpRequest(),
            h2 = document.createElement("h2"),
            ul = document.createElement("ul"),
            url = "/" + collection_name + "/read/" + keys.join(",");
        
        if (label === undefined || label === null ) {
            h2.innerHTML = '';
        } else {
            h2.innerHTML = label;
        }
        elem.appendChild(h2);
        elem.appendChild(ul);

        req.addEventListener("load", function(evt) {
            let objects = JSON.parse(req.response);
            for (let i = 0; i < objects.length; i++) {
                let key = objects[i]._Key,
                    title = objects[i].title,
                    href = "/" + collection_name + "/view.html?key="+key,
                    li = document.createElement("li"),
                    anchor = document.createElement("a");
                anchor.innerHTML = title;
                anchor.setAttribute("href", href);
                li.innerHTML = "";
                li.appendChild(anchor);
                ul.appendChild(li);
            }
        });
        req.addEventListener("error", function(evt) {
            console.log("Error happened", evt);
        });
        req.open("GET", url);
        req.send();
    }

    /**
     * viewAllObjects takes a DOM element and renders a list of all objects
     *
     * @param elem is a DOM element
     * @param collection_name is a string holding the collection name
     * @param keys is a list of object keys in the collection
     */
    function viewAllObjects(elem, collection_name, state) {
        let req = new XMLHttpRequest(),
            url = "/" + collection_name + "/keys/",
            label = "List all objects";
        
        if (state !== undefined && state !== null && state !== "") {
            url += `/state/${state}`;
            label = `${state} objects`;
        }

        req.addEventListener("load", function(evt) {
            let keys = JSON.parse(req.response);
            console.log("DEBUG label", label);
            viewObjectList(elem, collection_name, keys, label);
        });
        req.addEventListener("error", function(evt) {
            elem.innerHTML = "error";
            console.log("Error happened", evt);
        });
        req.open("GET", url);
        req.send();
    }

    /**
     * viewObject takes a DOM element and object id and renders a view
     * of the object.
     *
     * @param elem is a DOM element
     * @param collection_name is a string holding the collection name
     * @param key is an object key in the collection
     */
    function viewObject(elem, collection_name, key) {
        let div = document.createElement("div"),
            req = new XMLHttpRequest(), 
            url = "/" + collection_name + "/read/" + key;

        elem.appendChild(div)

        req.addEventListener("load", function(evt) {
            let o = JSON.parse(req.response),
                creator_names = [];
                div.innerHTML = "";
                for (let i = 0; i < o.creators.length; i++) {
                    let display_name = o.creators[i].display_name;
                    creator_names.push(display_name)
                }
                div.innerHTML = `<h1>${o.title}</h1>
<h2>Description</h2>
<div class="description">${o.description}</div>
<h2>Author(s)</h2>
<div class="creators">${creator_names.join(", ")}</div>
<h2>Pub Date</h2>
<div class="pub-date">${o.pubDate}</div>
<p>`;
        });
        req.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        req.open("GET", url);
        req.send();
    }

    /**
     * createObject takes a collection name and a template
     * and renders an form that is ready to collect the fields
     * and perform a post to /create/KEY and retrieves an empty
     * object.
     */

    /**
     * editObject takes an object and sets up a form.
     *
     * @param elem is the DOM element that is updated with object's 
     * attributes. The contents of the object will be replaced by 
     * a rendered template literal.
     * @param collection_name
     * @param key to object (empty means new object)
     * @param template liter that renders the form.
     * @param defaultObject is what will be used if the Key is not provided
     */
    function editObject(elem, collection_name, key, template) {
        let req = new XMLHttpRequest(), 
            url = "/" + collection_name + "/";

        if (key == undefined || key === null || key === "") {
            url = "/" + collection_name + "/create/" + key,
            object = defaultObject;
            return 
        }
    }

    /* Attach our functions to our object for export */
    AndOr.whoAmI = whoAmI;
    AndOr.viewAllObjects = viewAllObjects;
    AndOr.viewObjectList = viewObjectList;
    AndOr.viewObject = viewObject;
    AndOr.editObject = editObject;

    /* Now update our global object */
    if (window.AndOr == undefined) {
        window.AndOr = {};
    }
    window.AndOr = Object.assign(window.AndOr, AndOr);
}(window, document));
