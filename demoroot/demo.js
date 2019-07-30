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

    Demo.listObjects = function(elem) {
        var li = document.createElement("li");
        li.innerHTML = "Hello World (list objects)";
        elem.appendChild(li);
    };

    Demo.viewObject = function(elem) {
        var div = document.createElement("div");
        div.innerHTML = "Hello World (view object)";
        elem.appendChild(div)
    };
    window.Demo = Demo;
}(window, document));
