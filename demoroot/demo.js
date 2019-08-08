(function (window, document) {
    'use strict';
    let Demo = {};


    Demo.whoAmI = function (elem) {
        var div = document.createElement("div"),
            oReq = new XMLHttpRequest();
        div.innerHTML = "Retreiving your access profile&hellops;";
        elem.appendChild(div)
        oReq.addEventListener("progress", function(evt) {
            div.innerHTML = "progressing ...";
        });
        oReq.addEventListener("load", function(evt) {
            let o = JSON.parse(oReq.response),
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
            for (let i in assignments) {
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
        oReq.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        oReq.open("GET", "/demo-collection-01/access/");
        oReq.send();
    }

    function viewObjectList(elem, keys) {
        var readReq = new XMLHttpRequest(),
            h2 = document.createElement("h2"),
            ul = document.createElement("ul");
        
        h2.innerHTML = "List objects";
        elem.appendChild(h2);
        elem.appendChild(ul);

        readReq.addEventListener("load", function(evt) {
            let objects = JSON.parse(readReq.response);
            for (let i = 0; i < objects.length; i++) {
                let key = objects[i]._Key,
                    title = objects[i].title,
                    url = "/demo-collection-01/view.html?key="+key,
                    li = document.createElement("li"),
                    anchor = document.createElement("a");
                anchor.innerHTML = title;
                anchor.setAttribute("href", url);
                li.innerHTML = "";
                li.appendChild(anchor);
                ul.appendChild(li);
            }
        });
        readReq.addEventListener("error", function(evt) {
            console.log("Error happened", evt);
        });
        readReq.open("GET", "/demo-collection-01/read/" + keys.join(","));
        readReq.send();
    }

    Demo.listObjects = function(elem) {
        var keysReq = new XMLHttpRequest();
        elem.innerHTML = "Retreiving objects keys&hellops;";
        keysReq.addEventListener("load", function(evt) {
            let keys = JSON.parse(keysReq.response);
            elem.innerHTML = "";
            viewObjectList(elem, keys);
        });
        keysReq.addEventListener("error", function(evt) {
            elem.innerHTML = "error";
            console.log("Error happened", evt);
        });
        keysReq.open("GET", "/demo-collection-01/keys/");
        keysReq.send();
    };

    Demo.viewObject = function(elem, object_id) {
        var div = document.createElement("div"),
            oReq = new XMLHttpRequest();
        div.innerHTML = "Retreiving object profile&hellops;";
        elem.appendChild(div)
        oReq.addEventListener("progress", function(evt) {
            div.innerHTML = "progressing ...";
        });
        oReq.addEventListener("load", function(evt) {
            let o = JSON.parse(oReq.response),
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
        oReq.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        oReq.open("GET", "/demo-collection-01/read/"+ object_id);
        oReq.send();
    };
    window.Demo = Demo;
}(window, document));
