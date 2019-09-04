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
     * getCollectionName takes a URL pathname and returns the dirname
     * as the collection name.
     */
    function getCollectionName(url_path) {
        let p = url_path.split("/")
        if (p.length > 1) {
            return p[1];
        }
        return "unknown";
    }

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
       
        console.log("DEBUG URL", url);
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
        let u = new URL(window.location.href);

        u.pathname = `/${collection_name}/read/${keys.join(",")}`;
        console.log("DEBUG viewObjectList url", u);

        CL.httpGet(u, 'application/json', function (objects, err) {
            console.log("DEBUG got objects, err", objects, err);
            if (err) {
                console.log("Error happened", err);
                elem.innerHTML = `Unable to read ${collection_name} keys ${keys.join(", ")}<p>${err}`;
                return;
            }
            let h2 = document.createElement("h2"),
                ul = document.createElement("ul");
        
            if (label === undefined || label === null ) {
                h2.innerHTML = '';
            } else {
                h2.innerHTML = label;
            }
            elem.appendChild(h2);
            elem.appendChild(ul);

            // Handle single object case and array of objects case
            if (Array.isArray(objects)) {
                for (let i = 0; i < objects.length; i++) {
                    let key = objects[i]._Key,
                        title = `${objects[i].family_name + ", " + objects[i].given_name} (${key}, ${objects[i]._State})`,
                        href = "/" + collection_name + "/view.html?key="+key,
                        li = document.createElement("li"),
                        anchor = document.createElement("a");
                    anchor.innerHTML = title;
                    anchor.setAttribute("href", href);
                    li.innerHTML = "";
                    li.appendChild(anchor);
                    ul.appendChild(li);
                }
            } else {
                let key = objects._Key,
                    title = `${objects.family_name + ", " + objects.given_name} (${key}, ${objects._State})`,
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
    }

    /**
     * viewAllObjects takes a DOM element and renders a list of all objects
     *
     * @param elem is a DOM element
     * @param collection_name is a string holding the collection name
     * @param keys is a list of object keys in the collection
     */
    function viewAllObjects(elem, collection_name, state) {
        let u = new URL(window.location.href),
            label = "List all objects";

        u.pathname = "/" + collection_name + "/keys/";

        if (state !== undefined && state !== null && state !== "") {
            u.pathname += `state/${state}`;
            label = `${state} objects`;
            //FIXME: Need to get keys first and filter by state.
        }
        CL.httpGet(u, 'application/json', function(keys, err) {
            if (err) {
                console.log("Error happened", err);
                elem.innerHTML = `Unable to read ${collection_name} keys ${state}<p>${err}`;
                return;
            }
            console.log("DEBUG label", label, keys);
            if (Array.isArray(keys)  && keys.length > 0) {
                viewObjectList(elem, collection_name, keys, label);
            } else {
                elem.innerHTML = `Now objects found for ${collection_name}, ${state}`;
            }
        });
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
            u = new URL(window.location.href);

        u.pathname = "/" + collection_name + "/read/" + key;

        elem.appendChild(div)
        CL.httpGet(u, 'application/json', function (o, err) {
            if (err) {
                div.innerHTML = `error ${err}`;
                console.log("Error happened", evt);
                return;
            }
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
    }

    /**
     * createObject takes a collection name, a key, a JavaScript object
     * and sends it to the And/Or API.
     *
     * @param collection_name - the name of the dataset collection being served by andor (minus '.ds' ext.)
     * @param key - the key for the JSON object document
     * @param obj - the JavaScript object to be saved as a JSON object document
     * @param callbackFn - the callback function to handle the response from the POST
     */
    function createObject(collection_name, key, obj, callbackFn) {
        let u = new URL(window.location.href),
            payload = '';

        if (collection_name === undefined || collection_name === "") {
            console.log('WARNING: AndOr.createObject() called without collection name');
            return false;
        }
        if (key === undefined || key === null || key === "") {
            console.log('WARNING: AndOr.createObject() called without key');
            return false;
        }
        if (obj === undefined || obj === null) {
            console.log('WARNING: AndOr.createObject() called without JavaScript object');
            return false;
        }
        u.pathname = '/' + collection_name + '/create/' + key; 
        //FIXME: need to clear any searchParams ....
        payload = JSON.stringify(obj);
        //FIXME: Need to check to see if key does not exists and before calling create API, otherwise
        // return an error.
        let r = new XMLHttpRequest();
        console.log("DEBUG posting to", u);
        console.log("DEBUG payload is", payload);
        CL.httpPost(u, 'application/json', payload, function(data, err) {
            callbackFn(data, err);
            if (err) {
                console.log('ERROR: post ' + u + ' - ' + err);
                return;
            }
            console.log('DEBUG Success! ', data);
        });
    }

    /**
     * updateObject takes an object and sets up a form.
     *
     * @param collection_name
     * @param key to object (empty means new object)
     * @param obj - object to send as update
     * @param callbackFn - is the callback to handle response from the POST
     */
    function updateObject(collection_name, key, obj, callbackFn) {
        let url = new URL(window.location.href),
            payload = '';

        if (collection_name === undefined || collection_name === "") {
            console.log('WARNING: AndOr.updateObject() called without collection name');
            return false;
        }
        if (key === undefined || key === null || key === "") {
            console.log('WARNING: AndOr.updateObject() called without key');
            return false;
        }
        if (obj === undefined || obj === null) {
            console.log('WARNING: AndOr.updateObject() called without JavaScript object');
            return false;
        }
        url.pathname = '/' + collection_name + '/update/' + key; 
        payload = JSON.stringify(obj);
        //FIXME: Need to check to see if key does not exists and before calling update API, otherwise
        // return an error.
        CL.httpPost(url, 'application/json', payload, function(data, err) {
            callbackFn(data, err);
            if (err) {
                console.log('ERROR: post ' + url + ' - ' + err);
                return;
            }
            console.log('DEBUG Success! ', data);
        });
    }

    /**
     * readObject takes an collection name and key and gets the
     * object and calls the callback function.
     *
     * @param collection_name
     * @param key to object (empty means new object)
     * @param callbackFn - the callback function that returns the object and error from GET
     */
    function readObject(collection_name, key, callbackFn) {
        let u = new URL(window.location.href),
            payload = '';

        if (collection_name === undefined || collection_name === "") {
            console.log('WARNING: AndOr.readObject() called without collection name');
            return false;
        }
        if (key === undefined || key === null || key === "") {
            console.log('WARNING: AndOr.readObject() called without key');
            return false;
        }
        // Folder our keys if we have more than one.
        if (Array.isArray(key)) {
            key = key.join(',');
        }
        u.pathname = '/' + collection_name + '/read/' + key; 
        //FIXME: Need to clear any search params
        CL.httpGet(u, 'application/json', function(data, err) {
            callbackFn(data, err);
            if (err) {
                console.log('ERROR: post ' + url + ' - ' + err);
                return;
            }
            console.log('DEBUG Success! ', data);
        });
    }


    /**
     * getKeys takes an collection name and returns a list of keys via
     * the callback function.
     *
     * @param collection_name
     * @param key to object (empty means new object)
     * @param callbackFn - the callback function that returns the object and error from GET
     */
    function getKeys(collection_name, callbackFn) {
        let u = new URL(window.location.href),
            payload = [];

        if (collection_name === undefined || collection_name === "") {
            console.log('WARNING: AndOr.readObject() called without collection name');
            return false;
        }
        u.pathname = '/' + collection_name + '/keys';
        //FIXME: Need to clear any search params ...
        //FIXME: Need to check to see if key does not exists and before calling read API, otherwise
        // return an error.
        CL.httpGet(u, 'application/json', function(keys, err) {
            callbackFn(keys, err);
            if (err) {
                console.log('ERROR: post ' + url + ' - ' + err);
                return;
            }
            console.log('DEBUG Success! ', data);
        });
    }


    

    /* Attach our functions to our object for export */
    AndOr.getCollectionName = getCollectionName;
    AndOr.whoAmI = whoAmI;
    AndOr.viewAllObjects = viewAllObjects;
    AndOr.viewObjectList = viewObjectList;
    AndOr.viewObject = viewObject;
    AndOr.getKeys = getKeys;
    AndOr.createObject = createObject;
    AndOr.updateObject = updateObject;
    AndOr.readObject = readObject;

    /* Now update our global object */
    if (window.AndOr == undefined) {
        window.AndOr = {};
    }
    window.AndOr = Object.assign(window.AndOr, AndOr);
}(window, document));
