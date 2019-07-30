(function (window, document) {
    'use strict';
    let Demo = {};


    Demo.whoAmI = function (elem) {
        var div = document.createElement("div"),
            oReq = new XMLHttpRequest();
        div.innerHTML = "Retreiving your access profile&hellops;";
        elem.appendChild(div)
        oReq.addEventListener("progress", function(evt) {
            div.innertHTML = "progressing ...";
            console.log("Progressing");
        });
        oReq.addEventListener("load", function(evt) {
            let o = JSON.parse(oReq.response),
                user = o.user,
                workflows = o.workflows,
                blocks = [];
                let i = 0,
                    assignments = [];
                for (let key in workflows) {
                    assignments = assignments.concat(workflows[key].assign_to);
                }
                for (let i in assignments) {
                    if (assignments[i] == "*") {
                        assignments[i] = "any workflow or queue";
                    }
                }
                console.log("assigments", JSON.stringify(assignments));
                div.innerHTML = `<p>Welcome <em>${user.display_name}</em>,
<p>
You are a member of the following workflows
and queues&mdash; 
<ul><li>${user.member_of.join("</li><li>")}</li></ul>
<p>
You may assign objects to&mdash;
<ul><li>${assignments.join("</li><li>")}</li></ul>`;
        });
        oReq.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        oReq.open("GET", "/repository/access");
        oReq.send();
    }

    function listObject(ul, object_id) {
        let liReq = new XMLHttpRequest();
        liReq.addEventListener("load", function (evt) {
            let o = JSON.parse(liReq.response), 
                li = document.createElement("li"),
                anchor = document.createElement("a");
            anchor.setAttribute("href", "view.html?id="+object_id);
            anchor.innerHTML = o.title;
            li.appendChild(anchor);
            ul.appendChild(li);
        });
        liReq.open("GET", "/repository/objects/"+object_id);
        liReq.send();
    }

    Demo.listObjects = function(elem) {
        var div = document.createElement("div"),
            oReq = new XMLHttpRequest();
        div.innerHTML = "Retreiving objects profile&hellops;";
        elem.appendChild(div)
        oReq.addEventListener("progress", function(evt) {
            div.innertHTML = "progressing ...";
            console.log("Progressing");
        });
        oReq.addEventListener("load", function(evt) {
            let o_list = JSON.parse(oReq.response), blocks = [],
                ul = document.createElement("ul");
            div.innerHTML = "<h1>Objects list (unordered)</h1><p>";
            div.appendChild(ul);
            for (let i in o_list) {
                listObject(ul, o_list[i]);
            }
            console.log("loaded", JSON.stringify(o_list));
        });
        oReq.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        oReq.open("GET", "/repository/objects/");
        oReq.send();
    };

    Demo.viewObject = function(elem, object_id) {
        var div = document.createElement("div"),
            oReq = new XMLHttpRequest();
        div.innerHTML = "Retreiving object profile&hellops;";
        elem.appendChild(div)
        oReq.addEventListener("progress", function(evt) {
            div.innertHTML = "progressing ...";
            console.log("Progressing");
        });
        oReq.addEventListener("load", function(evt) {
            let o = JSON.parse(oReq.response);
                console.log("loaded", JSON.stringify(o));
                div.innerHTML = `<h1>${o.title}</h1>
<h2>URL</h2>
<div><a href="${o.official_url}">${o.official_url}</a></div>
<h2>abstract</h2>
<div class="abstract">${o.abstract}</div>
<h2>citation</h2>
<div class="citation">${o.official_cit}</div>
<p>`;
        });
        oReq.addEventListener("error", function(evt) {
            div.innerHTML = "error";
            console.log("Error happened", evt);
        });
        oReq.open("GET", "/repository/objects/"+ object_id);
        oReq.send();
    };
    window.Demo = Demo;
}(window, document));
