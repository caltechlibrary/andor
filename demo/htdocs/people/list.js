(function (window, document) {
   "use strict";
    let elem = document.getElementById("object-list"),
        u = new URL(window.location.href),
        c_name = AndOr.getCollectionName(u.pathname),
        keys = u.searchParams.get("keys"),
        state = u.searchParams.get("state");

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
                        state = objects[i]._State,
                        title = `${objects[i].family_name + ", " + objects[i].given_name} (edit ${key}, ${state})`,
                        href = "/" + collection_name + "/edit.html?cl_people_id="+key,
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
                    state = objects._State,
                    title = `${objects.family_name + ", " + objects.given_name} (edit ${key}, ${state})`,
                    href = "/" + collection_name + "/edit.html?cl_people_id="+key,
                    li = document.createElement("li"),
                    anchor = document.createElement("a");
                console.log("DEBUG objects ->", JSON.stringify(objects));
                console.log(`DEBUG key ${key}, state ${state}`, key, state);
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
                elem.innerHTML = `No objects found for ${collection_name}, ${state}`;
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


    if (elem !== undefined) {
        if (keys !== null && keys.includes(",")) {
           viewObjectList(elem, c_name, keys.split(","));
        } else if (state !== undefined && state !== null && state !== "") {
           viewAllObjects(elem, c_name, state);
        } else {
           viewAllObjects(elem, c_name);
        }
    }
}(window, document));

