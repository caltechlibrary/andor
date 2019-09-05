(function (document, window) {
    'use strict';
    /**
     * Page data
     */
    let cl = Object.assign({}, window.CL),
        AndOr = Object.assign({}, window.AndOr),
        div = document.getElementById('example-output'),
        page_url = new URL(window.location.href),
        c_name = AndOr.getCollectionName(page_url.pathname),
        key = "",
        people = {
            "family_name": "",
            "given_name": "",
            "cl_people_id": "",
            "thesis_id": "",
            "authors_id": "",
            "archivesspace_id": "",
            "directory_id": "",
            "viaf": "",
            "lcnaf": "",
            "isni": "",
            "wikidata": "",
            "snac": "",
            "orcid": "",
            "image": "",
            "educated_at": "",
            "caltech": false,
            "jpl": false,
            "faculty": false,
            "alumn": false,
            "notes": "",
            "_State": "deposit"
        },
        default_people = Object.assign({}, people),
        form_src = `
<style>
form.form-example {
    width: auto;
    height: auto;
    margin: 1.24em;
    padding: 1.24em;
}
form.form-example > div {
    padding-bottom: 0.72em;
}
form.form-example > div > input {
    min-width: 30em;
}
form.form-example > div > a {
    display: block;
}
form.form-example > div > textarea {
    min-width: 40em;
}
form.form-example > div > label {
    display: block; 
}
form.form-example > div > label.inline {
    display: inline-block; 
    width: 4em;
    padding-right:0;
    margin-right: 0;
}
form.form-example > div > input[type=checkbox] {
    min-width: 1em;
    padding-left: 1em;
    margin-left: 1em;
}
form.form-example > div > button {
    display: inline;
    padding: 0.32em;
    margin: 0.64em;
}
</style>
<form class="form-example">
<div>
    <label for="family_name">Family Name:</label>
    <input type="text" id="family_name" name="family_name" value="{{family_name}}" placeholder="e.g. Feyman">
</div>
<div>
    <label for="given_name">Given Name:</label>
    <input type="text" id="given_name" name="given_name" value="{{given_name}}" placeholder="e.g. Richard">
</div>
<div>
    <label for="cl_people_id">CL PEOPLE ID (<a href="https://feeds.library.caltech.edu/people/" target="_lookup">lookup</a>):</label>
    <input type="text" id="cl_people_id" name="cl_people_id" value="{{cl_people_id}}" placeholder="e.g. Feynman-R-P">
    <a id="cl_people_url" target="_lookup"></a>
</div>
<div>
    <label for="thesis_id">Thesis ID (<a href="https://thesis.library.caltech.edu/cgi/search/advanced" target="_lookup">lookup</a>):</label>
    <input type="text" id="thesis_id" name="thesis_id" value="{{thesis_id}}" placeholder="e.g. FEYNMAN-R-P">
    <a id="thesis_url" target="_lookup"></a>
</div>
<div>
    <label for="authors_id">Authors ID (<a href="https://authors.library.caltech.edu/cgi/search/advanced" target="_lookup">lookup</a>):</label>
    <input type="text" id="authors_id" name="authors_id" value="{{authors_id}}" placeholder="e.g. FEYNMAN-R-P">
    <a id="authors_url" target="_lookup"></a>
</div>
<div>
    <label for="archivesspace_id">ArchivesSpace ID (<a href="https://collections.archives.caltech.edu/agents" target="_lookup">lookup</a>):</label>
    <input type="text" id="archivesspace_id" name="archivesspace_id" value="{{archivesspace_id}}" placeholder="e.g. 3426">
    <a id="archivesspace_url" target="_lookup"></a>
</div>
<div>
    <label for="directory_id">Directory ID <a href="https://directory.caltech.edu" target="_lookup">(lookup):</a></label>
    <input type="text" id="directory_id" name="directory_id" value="{{directory_id}}" placeholder="e.g. rpfeynman">
    <a id="directory_url" target="_lookup"></a>
</div>
<div>
    <label for="viaf">VIAF ID (<a href="http://viaf.org/" target="_lookup">lookup</a>):</label>
    <input type="text" id="viaf" name="viaf" value="{{viaf}}" placeholder="e.g. 44298691">
    <a id="viaf_url" target="_lookup"></a>
</div>
<div>
    <label for="lcnaf">LCNAF (<a href="http://id.loc.gov/authorities/names.html" target="_lookup" title="Library of Congress Name Authority File">lookup</a>):</label>
    <input type="text" id="lcnaf" name="lcnaf" value="{{lcnaf}}" placeholder="n50002729">
    <a id="lcnaf_url" target="_lookup"></a>
</div>
<div>
    <label for="isni">ISNI (<a href="http://www.isni.org/search" target="_lookup">lookup</a>):</label>
    <input type="text" id="isni" name="isni" value="{{isni}}" placeholder="e.g. 0000 0001 2096 0218">
    <a id="isni_url" target="_lookup"></a>
</div>
<div>
    <label for="wikidata">Wikidata (<a href="https://www.wikidata.org/w/index.php?search=&search=&title=Special:Search&go=Go" target="_lookup">lookup</a>):</label>
    <input type="text" id="wikidata" name="wikidata" value="{{wikidata}}" placeholder="Q39246">
    <a id="wikidata_url" target="_lookup"></a>
</div>
<div>
    <label for="snac">SNAC (<a href="https://snaccooperative.org/" target="_lookup">lookup</a>):</label>
    <input type="text" id="snac" name="snac" value="{{snac}}" placeholder="e.g. ark:/99166/w6v69kzn">
    <a id="snac_url" target="_lookup"></a>
</div>
<div>
    <label for="orcid">ORCID (<a href="https://orcid.org/orcid-search/search/" target="_lookup">lookup</a>):</label>
    <input type="text" id="orcid" name="orcid" value="{{orcid}}">
    <a id="orcid_url" target="_lookup"></a>
</div>
<div> 
    <label for="image">Image:</label>
    <input type="url" id="image" name="image" value="{{image}}" placeholder="e.g. https://upload.wikimedia.org/wikipedia/en/4/42/Richard_Feynman_Nobel.jpg">
    <a id="image_url" target="_window"></a>
</div>
<div>
    <label for="educated_at">Educated At:</label>
    <textarea id="educated_at" name="educated_at" placeholder="e.g. Massachusetts Institute of Technology (S.B. 1939); Princeton University (Ph.D. 1942)">{{educated_at}}</textarea>
</div>
<div>
    <label class="inline" for="caltech">Caltech:</label>
    <input type="checkbox" id="caltech" name="caltech" {{caltech}} title="Check if affiliated with Caltech">
</div>
<div>
    <label class="inline" for="jpl">JPL:</label>
    <input type="checkbox" id="jpl" name="jpl" {{jpl}} title="check if affiliated with JPL">
</div>
<div>
    <label class="inline" for="faculty">Faculty:</label>
    <input type="checkbox" id="faculty" name="faculty" {{faculty}} title="check if Caltech Faculty">
</div>
<div>
    <label class="inline" for="alumn">Alumn:</label>
    <input type="checkbox" id="alumn" name="alumn" {{alumn}} title="check if Caltech Alumni">
</div>
<div>
    <label for="notes">Notes (internal use):</label>
    <textarea id="notes" name="notes">{{notes}}</textarea>
</div>
<div>
    <label for="_State">Status:</label>
    <select id="_State">
        <option value="{{_State}}" selected>{{_State}}</option>
        <option value="deposit">Deposit</option>
        <option value="review">Review</option>
        <option value="embargoed">Embargoed</option>
        <option value="published">Published</option>
        <option value="deleted">Deleted</option>
    </select>
</div>
<div>
<button id="create">Create</button>
<button id="save">Save</button>
<button id="reset">Reset</button>
</div>
</form><!-- END: form.form-example -->
`;
    
    /**
     * Check for key or cl_people_id being password to form.
     */
    //key = page_url.searchParams.get("key");
    console.log("DEBUG key now ->", key, typeof key);
    
    key = page_url.searchParams.get("key"); 
    if (key === undefined || key === "" || key === null) {
        key = page_url.searchParams.get("cl_people_id"); 
console.log("DEBUG key now ->", key, page_url.searchParams.get('cl_people_id'));
        if (key === undefined || key === null) {
            key = "";
        }
    }  else {
        key = "";
    }
console.log("DEBUG key now ->", key);

    /**
     * Page functions
     */
    
    //console.log("DEBUG page_url", page_url);
    //console.log("DEBUG c_name", c_name);
    //console.log("DEBUG key ->", key, "<-");
    
    function setupAnchor(elem, link_text, prefix, suffix, value) {
        if (value === "" ) {
            elem.setAttribute("href", value);
            elem.innerHTML = "";
            return;
        }
        elem.setAttribute("href", prefix + value + suffix);
        elem.innerHTML = link_text;
    }
    
    function capitalize_word(w) {
        return w.charAt(0).toUpperCase()+w.slice(1);
    }
    
    function capitalize_string(s, sep = '-') {
        if (s === "") {
            return s;
        }
        return s.toLowerCase().split(sep).map(capitalize_word).join(sep);
    }
    
    function form_init() {
        let obj = this;
        // NOTE: Convert bool true to 'checked' and
        // false to empty string for using in 
        // checkbox input.
        for (let key in obj) {
             if (obj[key] === true) {
                 obj[key] = "checked";
             } else if (obj[key] === false) {
                 obj[key] = "";
             }
        }
        if (obj.cl_people_id === undefined || obj.cl_people_id === "") {
            if (obj._Key !== undefined) {
                obj.cl_people_id = obj._Key;
            } else if (key !== undefined && key !== "") {
                obj.cl_people_id = key;
                obj._Key = key;
            }
        }
        return true;
    }

    function render_form(elem, people, form_src, form_init) {
        elem.innerHTML = "";
        let field = CL.field(people, form_src, form_init);
console.log("DEBUG field", field);
        let form = CL.assembleFields(elem, field),
            /**
             * Now we add our event listeners and lookups using
             * vanilla JavaScript and CL.httpGet()..
             */
            family_name = form.querySelector("#family_name"),
            given_name = form.querySelector("#given_name"),
            cl_people_id = form.querySelector("#cl_people_id"),
            cl_people_url = form.querySelector("#cl_people_url"),
            thesis_id = form.querySelector("#thesis_id"),
            thesis_url = form.querySelector("#thesis_url"),
            authors_id = form.querySelector("#authors_id"),
            authors_url = form.querySelector("#authors_url"),
            archivesspace_id = form.querySelector("#archivesspace_id"),
            archivesspace_url = form.querySelector("#archivesspace_url"),
            directory_id = form.querySelector("#directory_id"),
            directory_url = form.querySelector("#directory_url"),
            viaf = form.querySelector("#viaf"),
            viaf_url = form.querySelector("#viaf_url"),
            lcnaf = form.querySelector("#lcnaf"),
            lcnaf_url = form.querySelector("#lcnaf_url"),
            isni_url = form.querySelector("#isni_url"),
            wikidata = form.querySelector("#wikidata"),
            wikidata_url = form.querySelector("#wikidata_url"),
            snac = form.querySelector("#snac"),
            snac_url = form.querySelector("#snac_url"),
            orcid = form.querySelector("#orcid"),
            orcid_url = form.querySelector("#orcid_url"),
            image = form.querySelector("#image"),
            image_url = form.querySelector("#image_url"),
            educated_at = form.querySelector("#educated_at"),
            caltech = form.querySelector("#caltech"),
            jpl = form.querySelector("#jpl"),
            faculty = form.querySelector("#faculty"),
            alumn = form.querySelector("#alumn"),
            notes = form.querySelector("#notes"),
            create = form.querySelector("#create"),
            _State = form.querySelector("#_State"),
            save = form.querySelector("#save"),
            reset = form.querySelector("#reset"); 
        
        family_name.addEventListener("change", function(evt) {
            people.family_name = this.value;
        });
        given_name.addEventListener("change", function(evt) {
            people.given_name = this.value;
        });
        cl_people_id.addEventListener("change", function(evt) {
            people.cl_people_id = capitalize_string(this.value, '-');
            setupAnchor(cl_people_url, 
                'Check Feeds for ' + people.cl_people_id, 
                'https://feeds.library.caltech.edu/people/', 
                '',
                people.cl_people_id);
            this.value = people.cl_people_id;
        });
        thesis_id.addEventListener("change", function (evt) {
            people.thesis_id = capitalize_string(this.value, '-');
            setupAnchor(thesis_url, 
                'Check CaltechTHESIS (as author) for ' + people.thesis_id, 
                'https://thesis.library.caltech.edu/view/author/',
                '.html',
                people.thesis_id);
            this.value = people.thesis_id;
        });
        authors_id.addEventListener("change", function(evt) {
            people.authors_id = capitalize_string(this.value);
            setupAnchor(authors_url, 
                'Check CaltechAUTHORS for ' + people.authors_id, 
                'https://authors.library.caltech.edu/view/person-az/',
                '.html',
                people.authors_id);
            this.value = people.authors_id;
        });
        archivesspace_id.addEventListener("change", function(evt) {
            people.archivesspace_id = this.value;
            setupAnchor(archivesspace_url, 
                'Check Caltech Archives for ' + this.value, 
                'https://collections.archives.caltech.edu/agents/people/',
                '',
                this.value);
        });
        directory_id.addEventListener("change", function(evt) {
            people.directory_id = this.value;
            setupAnchor(directory_url, 
                'Check Caltech Directory for ' + this.value, 
                'https://directory.caltech.edu/personnel/',
                '',
                this.value);
        });
        viaf.addEventListener("change", function(evt) {
            people.viaf = this.value;
            setupAnchor(viaf_url, 
                'Check VIAF.org for ' + this.value, 
                'https://viaf.org/viaf/',
                '/',
                this.value);
        });
        lcnaf.addEventListener("change", function(evt) {
            people.lcnaf = this.value;
            setupAnchor(lcnaf_url, 
                'Check LOC Name Authority File for ' + this.value, 
                'http://id.loc.gov/authorities/names/',
                '',
                this.value);
        });
        isni.addEventListener("change", function(evt) {
            people.isni = this.value;
            setupAnchor(isni_url, 
                'Check ISNI for ' + this.value, 
                'http://isni.oclc.org/DB=1.2/SET=4/TTL=1/CMD?ACT=SRCH&IKT=6102&SRT=LST_nd&TRM=ISN%3A',
                '',
                this.value);
        });
        wikidata.addEventListener("change", function(evt) {
            people.wikidata = this.value;
            setupAnchor(wikidata_url, 
                'Check Wikidata for ' + this.value, 
                'https://www.wikidata.org/wiki/',
                '',
                this.value);
        });
        snac.addEventListener("change", function(evt) {
            people.snac = this.value;
            setupAnchor(snac_url, 
                'Check SNAC for ' + this.value, 
                'https://snaccooperative.org/',
                '',
                this.value);
        });
        orcid.addEventListener("change", function(evt) {
            people.orcid = this.value;
            setupAnchor(orcid_url, 
                'Check ORCID for ' + this.value, 
                'https://orcid.org/',
                '',
                this.value);
        });
        image.addEventListener("change", function(evt) {
            people.image = this.value;
            setupAnchor(image_url, 
                'Image preview ' + this.value, 
                '',
                '',
                this.value);
        });
        educated_at.addEventListener("change", function(evt) {
            people.educated_at = this.value;
        });
        caltech.addEventListener("change", function(evt) {
            if (this.checked) {
                people.caltech = true;
            } else {
                people.caltech = false;
            }
        });
        jpl.addEventListener("change", function(evt) {
            if (this.checked) {
                people.jpl = true;
            } else {
                people.jpl = false;
            }
        });
        faculty.addEventListener("change", function(evt) {
            if (this.checked) {
                people.faculty = true;
            } else {
                people.faculty = false;
            }
        });
        alumn.addEventListener("change", function(evt) {
            if (this.checked) {
                people.alumn = true;
            } else {
                people.alumn = false;
            }
        });
        notes.addEventListener("change", function(evt) {
            people.notes = this.value;
        });
        _State.addEventListener("change", function(evt) {
            people._State = this.value;
        });
        create.addEventListener("click", function(evt) {
            console.log("DEBUG people before create", people);
            if (people.cl_people_id === undefined || people.cl_people_id === "") {
                evt.preventDefault();
                return;
            }
            console.log("DEBUG people payload ->", JSON.stringify(people));
            //FIXME: Check to see if key exists
            AndOr.createObject(c_name, people.cl_people_id, people, function(data, err) {
                if (err) {
                    console.log("DEBUG can't create object,", err);
                    evt.preventDefault(); 
                    return;
                }
                console.log("DEBUG createObject() -> data", data, " error ", err);
            });
            evt.preventDefault();
        }, false);
        save.addEventListener("click", function (evt) {
            console.log("DEBUG people.cl_people_id before save", people.cl_people_id, typeof people.cl_people_id);
            console.log("DEBUG saving people payload ->", JSON.stringify(people));
            if (people._Key !== undefined && people.cl_people_id === "") {
                people.cl_people_id = people._Key;
            }
            if (people.cl_people_id === undefined || people.cl_people_id === "") {
                evt.preventDefault();
                return;
            }
            //FIXME: Check to see if key exists
            console.log(`DEBUG calling AndOr.updateObject(${c_name}, ${people.cl_people_id}, ...)`);
            AndOr.updateObject(c_name, people.cl_people_id, people, function(data, err) {
                if (err) {
                    console.log("DEBUG can't save object,", err);
                    evt.preventDefault(); 
                    return;
                }
                console.log("DEBUG updateObject() -> data", data, " error ", err);
            });
            evt.preventDefault(); 
        }, false);
        reset.addEventListener("click", function(evt) {
            // NOTE: Need to clear any key/cl_people_id settings in
            // URL.
            let u = new URL(window.location.href);
            console.log("DEBUG before to change u.search", u.search);
            u.search = "";
            console.log("DEBUG after to change u.search", u.search);
            window.location = u;
            console.log("DEBUG after reset window");
            evt.preventDefault();
        });
    }


    /**
     * Main, apply main logic for page.
     */
    if (key !== undefined && key !== "") {
        let t1 = Date.now(), t2 = Date.now();
        console.log(`DEBUG (${(t2 - t1)/1000}) key is defined, waiting on readObject`);
        div.innerHTML = "Retrieving " + key;
        let updateOK = false,
            tid = -1;
        AndOr.readObject(c_name, key, function(data, err) {
            t2 = Date.now();
            console.log(`DEBUG (${(t2 - t1)/1000}) got data, err`, data, err);
            if (err) {
                    console.log("readObject() error", err);
                    clearInterval(tid);
                    return;
            }
            people = Object.assign(people, data);
            render_form(div, people, form_src, form_init);
            updateOK = true;
            if (tid >= 0) {
                clearInterval(tid);
            }
        });
        t2 = Date.now();
        console.log(`DEBUG (${(t2 - t1)/1000}) (waiting on) AndOr.readObject(${c_name}, ${key}, ...) timer`);
        // We're polling for results here ...
        tid = setInterval(function() {
            console.log(`DEBUG (${(t2 - t1)/1000}) (waiting on) AndOr.readObject(${c_name}, ${key}, ...) timer ${tid}`);
            if (updateOK == true) {
                console.log(`DEBUG (${(t2 - t1)/1000}) clearInterval timer id ${tid}`);
                clearInterval(tid);
            }
            t2 = Date.now();
        }, 1 * 1000);
    } else {
        console.log("DEBUG key is NOT defined, using default object");
        people = Object.assign({}, default_people);
        console.log("DEBUG render form with default_people", people.cl_people_id);
        render_form(div, people, form_src, form_init);
    }

    window.People = people; //DEBUG
}(document, window));
