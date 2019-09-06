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
            // And/Or required fields.
            "_State": "deposit",
            "_Key": ""
        },
        default_people = Object.assign({}, people),
        form_src = `
<style>
form.form-people {
    width: auto;
    height: auto;
    margin: 1.24em;
    padding: 1.24em;
}
form.form-people > div > div {
    padding-bottom: 0.72em;
}
form.form-people > div > div > input {
    min-width: 30em;
}
form.form-people > div > div > a {
    display: block;
}
form.form-people > div > div > textarea {
    min-width: 40em;
}
form.form-people > div > div > label {
    display: block; 
}
form.form-people > div > div > label.inline {
    display: inline-block; 
    width: 4em;
    padding-right:0;
    margin-right: 0;
}
form.form-people > div > div > input[type=checkbox] {
    min-width: 1em;
    padding-left: 1em;
    margin-left: 1em;
}
.errors, form.form-people > div > div.required > label {
    color: red;
    font-style: bold;
}
.controls {
    width: 10%;
    border: 0.024em solid green;
    padding:1.24em;
    margin: 1.24em;
    display: inline;
    float: left;
}
.controls > div > button {
    width: 100%;
    padding-top: 1em;
    padding-bottom: 1em;
    margin-top: 1em;
    margin-bottom: 1em;
}
.fields {
    padding: 0;
    margin: 0;
    display: inline;
    float: left;
}
</style>
<form class="form-people">
<div class="controls">
    <div>
        <div id="errors" class="errors"></div>
        <div id="status" class="status"></div>
        <hr>
        <button tabindex="21" id="create" disabled="true" title="create a new people record">Create</button>
        <button tabindex="22" id="read" title="retrieve a people record">Read</button>
        <button tabindex="23" id="save" disabled="true" title="update an existing people record, including assigning to a new state">Update</button>
        <button tabindex="24" id="reset" title="clear the current form and replace with an empty people record">Clear Form</button>
    </div>
</div><!-- END: .controls -->
<div class="fields">
    <div class="required">
        <label for="cl_people_id"><span title="this is a required field">CL PEOPLE ID</span> (<a href="https://feeds.library.caltech.edu/people/" target="_lookup">lookup</a>):</label>
        <input tabindex="1" type="text" id="cl_people_id" name="cl_people_id" value="{{cl_people_id}}" placeholder="e.g. Feynman-R-P" title="a Caltech People ID is required">
        <a id="cl_people_url" target="_lookup"></a>
    </div>
    <div>
        <label for="family_name">Family Name:</label>
        <input tabindex="2" type="text" id="family_name" name="family_name" value="{{family_name}}" placeholder="e.g. Feyman">
    </div>
    <div>
        <label for="given_name">Given Name:</label>
        <input tabindex="3" type="text" id="given_name" name="given_name" value="{{given_name}}" placeholder="e.g. Richard">
    </div>
    <div>
        <label for="thesis_id">Thesis ID (<a href="https://thesis.library.caltech.edu/cgi/search/advanced" target="_lookup">lookup</a>):</label>
        <input tabindex="4" type="text" id="thesis_id" name="thesis_id" value="{{thesis_id}}" placeholder="e.g. FEYNMAN-R-P">
        <a id="thesis_url" target="_lookup"></a>
    </div>
    <div>
        <label for="authors_id">Authors ID (<a href="https://authors.library.caltech.edu/cgi/search/advanced" target="_lookup">lookup</a>):</label>
        <input tabindex="5" type="text" id="authors_id" name="authors_id" value="{{authors_id}}" placeholder="e.g. FEYNMAN-R-P">
        <a id="authors_url" target="_lookup"></a>
    </div>
    <div>
        <label for="archivesspace_id">ArchivesSpace ID (<a href="https://collections.archives.caltech.edu/agents" target="_lookup">lookup</a>):</label>
        <input tabindex="6" type="text" id="archivesspace_id" name="archivesspace_id" value="{{archivesspace_id}}" placeholder="e.g. 3426">
        <a id="archivesspace_url" target="_lookup"></a>
    </div>
    <div>
        <label tabindex="7" for="directory_id">Directory ID <a href="https://directory.caltech.edu" target="_lookup">(lookup):</a></label>
        <input type="text" id="directory_id" name="directory_id" value="{{directory_id}}" placeholder="e.g. rpfeynman">
        <a id="directory_url" target="_lookup"></a>
    </div>
    <div>
        <label for="viaf">VIAF ID (<a href="http://viaf.org/" target="_lookup">lookup</a>):</label>
        <input tabindex="8" type="text" id="viaf" name="viaf" value="{{viaf}}" placeholder="e.g. 44298691">
        <a id="viaf_url" target="_lookup"></a>
    </div>
    <div>
        <label for="lcnaf">LCNAF (<a href="http://id.loc.gov/authorities/names.html" target="_lookup" title="Library of Congress Name Authority File">lookup</a>):</label>
        <input tabindex="9" type="text" id="lcnaf" name="lcnaf" value="{{lcnaf}}" placeholder="n50002729">
        <a id="lcnaf_url" target="_lookup"></a>
    </div>
    <div>
        <label for="isni">ISNI (<a href="http://www.isni.org/search" target="_lookup">lookup</a>):</label>
        <input tabindex="10" type="text" id="isni" name="isni" value="{{isni}}" placeholder="e.g. 0000 0001 2096 0218">
        <a id="isni_url" target="_lookup"></a>
    </div>
    <div>
        <label for="wikidata">Wikidata (<a href="https://www.wikidata.org/w/index.php?search=&search=&title=Special:Search&go=Go" target="_lookup">lookup</a>):</label>
        <input tabindex="11" type="text" id="wikidata" name="wikidata" value="{{wikidata}}" placeholder="Q39246">
        <a id="wikidata_url" target="_lookup"></a>
    </div>
    <div>
        <label for="snac">SNAC (<a href="https://snaccooperative.org/" target="_lookup">lookup</a>):</label>
        <input tabindex="12" type="text" id="snac" name="snac" value="{{snac}}" placeholder="e.g. ark:/99166/w6v69kzn">
        <a id="snac_url" target="_lookup"></a>
    </div>
    <div>
        <label for="orcid">ORCID (<a href="https://orcid.org/orcid-search/search/" target="_lookup">lookup</a>):</label>
        <input tabindex="13" type="text" id="orcid" name="orcid" value="{{orcid}}">
        <a id="orcid_url" target="_lookup"></a>
    </div>
    <div> 
        <label for="image">Image:</label>
        <input tabindex="14" type="url" id="image" name="image" value="{{image}}" placeholder="e.g. https://upload.wikimedia.org/wikipedia/en/4/42/Richard_Feynman_Nobel.jpg">
        <a id="image_url" target="_window"></a>
    </div>
    <div>
        <label for="educated_at">Educated At:</label>
        <textarea id="educated_at" name="educated_at" placeholder="e.g. Massachusetts Institute of Technology (S.B. 1939); Princeton University (Ph.D. 1942)">{{educated_at}}</textarea>
    </div>
    <div>
        <label class="inline" for="caltech">Caltech:</label>
        <input tabindex="15" type="checkbox" id="caltech" name="caltech" {{caltech}} title="Check if affiliated with Caltech">
    </div>
    <div>
        <label class="inline" for="jpl">JPL:</label>
        <input tabindex="16" type="checkbox" id="jpl" name="jpl" {{jpl}} title="check if affiliated with JPL">
    </div>
    <div>
        <label class="inline" for="faculty">Faculty:</label>
        <input tabindex="17" type="checkbox" id="faculty" name="faculty" {{faculty}} title="check if Caltech Faculty">
    </div>
    <div>
        <label class="inline" for="alumn">Alumn:</label>
        <input tabindex="18" type="checkbox" id="alumn" name="alumn" {{alumn}} title="check if Caltech Alumni">
    </div>
    <div>
        <label for="notes">Notes (internal use):</label>
        <textarea tabindex="19" id="notes" name="notes">{{notes}}</textarea>
    </div>
    <div>
        <label for="_State">Status:</label>
        <select tabindex="20" id="_State">
            <option value="{{_State}}" selected>{{_State}}</option>
            <option value="deposit">Deposit</option>
            <option value="review">Review</option>
            <option value="embargoed">Embargoed</option>
            <option value="published">Published</option>
            <option value="deleted">Deleted</option>
        </select>
    </div>
</div><!-- END: .fields -->
</form><!-- END: form.form-people -->
`;

    /**
     * Check for cl_people_id being passed in URL
     */
    key = page_url.searchParams.get("cl_people_id");
    if (! key) {
        key = "";
    }
    console.log("DEBUG key now ->", key, typeof key);


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
        for (let aKey in obj) {
             if (obj[aKey] === true) {
                 obj[aKey] = "checked";
             } else if (obj[aKey] === false) {
                 obj[aKey] = "";
             }
        }
        if (obj.cl_people_id === undefined || obj.cl_people_id === "") {
            if (obj._Key !== undefined && obj._Key !== "") {
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
            _State = form.querySelector("#_State"),
            _errors = form.querySelector("#errors"),
            _status = form.querySelector("#status"),
            create = form.querySelector("#create"),
            read = form.querySelector("#read"),
            save = form.querySelector("#save"),
            reset = form.querySelector("#reset"); 
        
        create.setAttribute("disabled", "disabled");
        save.setAttribute("disabled", "disabled");
        if (key === undefined || key === "" || key === null) {
            create.removeAttribute("disabled");
        } else  {
            save.removeAttribute("disabled");
        }

        // Setup rest of form.
        family_name.addEventListener("change", function(evt) {
            people.family_name = this.value;
        });
        given_name.addEventListener("change", function(evt) {
            people.given_name = this.value;
        });
        if (cl_people_id.value) {
            setupAnchor(cl_people_url, 
                'Check Feeds for ' + people.cl_people_id, 
                'https://feeds.library.caltech.edu/people/', 
                '',
                people.cl_people_id);
        }
        cl_people_id.addEventListener("change", function(evt) {
            let u = new URL(window.location.href),
                c_name = AndOr.getCollectionName(u.pathname),
                old_id = people.cl_people_id, old_family_name = people.cl_people_id;

            people.cl_people_id = capitalize_string(this.value, '-');
            people._Key = people.c_people_id;
            cl_people_id.value = people.cl_people_id;

            if (people.cl_people_id) {
                u.pathname = `/${c_name}/read/${people.cl_people_id}`;
                u.search = '';
                console.log("DEBUG looking up", u);
                CL.httpGet(u, 'application/json', function (data, err) {
                    if (err) {
                        console.log("DEBUG people.cl_people_id not found");
                        create.removeAttribute("disabled");
                        save.setAttribute("disabled", "disabled");
                        cl_people_url.innerHTML = 'Lookup people ids at feeds.library.caltech.edu';
                        cl_people_url.setAttribute('href', 'https://feeds.library.caltech.edu/people');
                        _status.innerHTML = `Ready to create a new records for ${people.cl_people_id} ${(new Date).toLocaleTimeString()}`;
                        family_name.focus();
                        return
                    }
                    console.log(`DEBUG people.cl_people_id ${people.cl_people_id} found`);
                    _status.innerHTML = `${people.cl_people_id} found ${(new Date()).toLocaleTimeString()}`;
                    setupAnchor(cl_people_url, 
                        'Check Feeds for ' + people.cl_people_id, 
                        'https://feeds.library.caltech.edu/people/', 
                        '',
                        people.cl_people_id);
                    create.setAttribute("disabled", "disabled");
                    read.removeAttribute("disabled");
                    read.focus();
                });
            } else {
                setupAnchor(cl_people_url, 
                    'Check Feeds ' + people.cl_people_id, 
                    'https://feeds.library.caltech.edu/people/', 
                    '',
                    people.cl_people_id);
                create.removeAttribute("disabled");
                save.setAttribute("disabled", "disabled");
            }
        });
        if (thesis_id.value) {
            setupAnchor(thesis_url, 
                'Check CaltechTHESIS (as author) for ' + thesis_id.value, 
                'https://thesis.library.caltech.edu/view/author/',
                '.html',
                thesis_id.value);
        }
        thesis_id.addEventListener("change", function (evt) {
            people.thesis_id = capitalize_string(this.value, '-');
            setupAnchor(thesis_url, 
                'Check CaltechTHESIS (as author) for ' + people.thesis_id, 
                'https://thesis.library.caltech.edu/view/author/',
                '.html',
                people.thesis_id);
            this.value = people.thesis_id;
        });
        if (authors_id.value) {
            setupAnchor(authors_url, 
                'Check CaltechAUTHORS for ' + authors_id.value, 
                'https://authors.library.caltech.edu/view/person-az/',
                '.html',
                authors_id.value);
        }
        authors_id.addEventListener("change", function(evt) {
            people.authors_id = capitalize_string(this.value);
            setupAnchor(authors_url, 
                'Check CaltechAUTHORS for ' + people.authors_id, 
                'https://authors.library.caltech.edu/view/person-az/',
                '.html',
                people.authors_id);
            this.value = people.authors_id;
        });
        if (archivesspace_id.value) {
            setupAnchor(archivesspace_url, 
                'Check Caltech Archives for ' + archivesspace_id.value, 
                'https://collections.archives.caltech.edu/agents/people/',
                '',
                archivesspace_id.value);
        }
        archivesspace_id.addEventListener("change", function(evt) {
            people.archivesspace_id = this.value;
            setupAnchor(archivesspace_url, 
                'Check Caltech Archives for ' + this.value, 
                'https://collections.archives.caltech.edu/agents/people/',
                '',
                this.value);
        });
        if (directory_id.value) {
            setupAnchor(directory_url, 
                'Check Caltech Directory for ' + directory_id.value, 
                'https://directory.caltech.edu/personnel/',
                '',
                directory_id.value);
        }
        directory_id.addEventListener("change", function(evt) {
            people.directory_id = this.value;
            setupAnchor(directory_url, 
                'Check Caltech Directory for ' + this.value, 
                'https://directory.caltech.edu/personnel/',
                '',
                this.value);
        });
        if (viaf.value) {
            setupAnchor(viaf_url, 
                'Check VIAF.org for ' + viaf.value, 
                'https://viaf.org/viaf/',
                '/',
                viaf.value);
        }
        viaf.addEventListener("change", function(evt) {
            people.viaf = this.value;
            setupAnchor(viaf_url, 
                'Check VIAF.org for ' + this.value, 
                'https://viaf.org/viaf/',
                '/',
                this.value);
        });
        if (lcnaf.value) {
            setupAnchor(lcnaf_url, 
                'Check LOC Name Authority File for ' + lcnaf.value, 
                'http://id.loc.gov/authorities/names/',
                '',
                lcnaf.value);
        }
        lcnaf.addEventListener("change", function(evt) {
            people.lcnaf = this.value;
            setupAnchor(lcnaf_url, 
                'Check LOC Name Authority File for ' + this.value, 
                'http://id.loc.gov/authorities/names/',
                '',
                this.value);
        });
        if (isni.value) {
            setupAnchor(isni_url, 
                'Check ISNI for ' + isni.value, 
                'http://isni.oclc.org/DB=1.2/SET=4/TTL=1/CMD?ACT=SRCH&IKT=6102&SRT=LST_nd&TRM=ISN%3A',
                '',
                isni.value);
        }
        isni.addEventListener("change", function(evt) {
            people.isni = this.value;
            setupAnchor(isni_url, 
                'Check ISNI for ' + this.value, 
                'http://isni.oclc.org/DB=1.2/SET=4/TTL=1/CMD?ACT=SRCH&IKT=6102&SRT=LST_nd&TRM=ISN%3A',
                '',
                this.value);
        });
        if (wikidata.value) {
            setupAnchor(wikidata_url, 
                'Check Wikidata for ' + wikidata.value, 
                'https://www.wikidata.org/wiki/',
                '',
                wikidata.value);
        }
        wikidata.addEventListener("change", function(evt) {
            people.wikidata = this.value;
            setupAnchor(wikidata_url, 
                'Check Wikidata for ' + this.value, 
                'https://www.wikidata.org/wiki/',
                '',
                this.value);
        });
        if (snac.value) {
            setupAnchor(snac_url, 
                'Check SNAC for ' + snac.value, 
                'https://snaccooperative.org/',
                '',
                snac.value);
        }
        snac.addEventListener("change", function(evt) {
            people.snac = this.value;
            setupAnchor(snac_url, 
                'Check SNAC for ' + this.value, 
                'https://snaccooperative.org/',
                '',
                this.value);
        });
        if (orcid.value) {
            setupAnchor(orcid_url, 
                'Check ORCID for ' + orcid.value, 
                'https://orcid.org/',
                '',
                orcid.value);
        }
        orcid.addEventListener("change", function(evt) {
            people.orcid = this.value;
            setupAnchor(orcid_url, 
                'Check ORCID for ' + this.value, 
                'https://orcid.org/',
                '',
                this.value);
        });
        if (image.value) {
            setupAnchor(image_url, 
                'Image preview ' + image.value, 
                '',
                '',
                image.value);
        }
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
            _status.innerHTML = `Creating ${people.cl_people_id}`;
            _errors.innerHTML = '';
            //FIXME: Check to see if key exists
            AndOr.createObject(c_name, people.cl_people_id, people, function(data, err) {
                if (err) {
                    _status.innerHTML = '';
                    _errors.innerHTML = `ERROR: Can not create ${people.cl_people_id}, API message ${err}`;
                    evt.preventDefault(); 
                    return;
                }
                _status.innerHTML = `Created ${people.cl_people_id} ${(new Date()).toLocaleTimeString()}`
                _errors.innerHTML = '';
                // We've created our object so let's disable create
                // and enable save for further edits
                create.setAttribute("disabled", "disabled");
                save.removeAttribute("disabled");
            });
            evt.preventDefault();
        }, false);
        read.addEventListener("click", function(evt) {
            let key = people.cl_people_id,
                u = new URL(window.location.href);
            if (key === undefined || key === "") {
                _status.innerHTML = '';
                _errors.innerHTML = 'ERROR: Missing CL PEOPLE ID, a required field';
                evt.preventDefault();
                return;
            }
            // NOTE: We're reloading the form on read.
            u.search = "?cl_people_id="+key;
            console.log("DEBUG reading", u);
            window.location = u.toString();
            evt.preventDefault();
        }, false);
        save.addEventListener("click", function (evt) {
            let key = people.cl_people_id;

            if (key === undefined || key === "") {
                _status.innerHTML = '';
                _errors.innerHTML = `ERROR: Missing "CL PEOPLE ID" field, ${(new Date()).toLocaleTimeString()}`;
                cl_people_id.focus();
                evt.preventDefault();
                return;
            }
            // Force the ._Key to match the key/people.cl_people_id value.
            people._Key = key;
            _status.innerHTML = `Saving ${people.cl_people_id}`
            _errors.innerHTML = '';
            //FIXME: Check to see if key exists
            console.log(`DEBUG calling AndOr.updateObject(${c_name}, ${key}, ...)`);
            AndOr.updateObject(c_name, key, people, function(data, err) {
                if (err) {
                    _status.innerHTML = '';
                    _errors.innerHTML = `ERROR: Can not save ${people.cl_people_id}, ${(new Date()).toLocaleTimeString()}`;
                    console.log(`ERROR: JSON API message ${err} ${(new Date()).toLocaleTimeString()}`);
                    evt.preventDefault(); 
                    return;
                }
                _status.innerHTML = `Saved ${people.cl_people_id} ${(new Date()).toLocaleTimeString()}`;
                _errors.innerHTML = '';
            });
            evt.preventDefault(); 
        }, false);
        reset.addEventListener("click", function(evt) {
            // NOTE: Need to clear any key/cl_people_id settings in
            // URL.
            let u = new URL(window.location.href);

            // On reset form by reloading the page without a cl_people_id defined.
            u.search = "";
            key = "";
            window.location = u.toString();
            evt.preventDefault();
        });
        return elem;
    }


    /**
     * Main, apply main logic for page.
     */
    if (key !== undefined && key !== null && key !== "") {
        div.innerHTML = "Retrieving " + key;
        let updateOK = false,
            tid = -1;
        AndOr.readObject(c_name, key, function(data, err) {
            if (err) {
                div.innerHTML = `ERROR: Record ${key} not found, JSON API message ${err}`;
                people = Object.assign({}, default_people);
                render_form(div, people, form_src, form_init);
                let _errors = div.querySelector('#errors');
                _errors.innerHTML = `ERROR: Record ${key} not found, JSON API message ${err}`;
                clearInterval(tid);
                return;
            }
            people = Object.assign(people, data);
            render_form(div, people, form_src, form_init);
            let dt = new Date(),
                _status = div.querySelector('#status');
            _status.innerHTML = `Retrieved ${key} ${dt.toLocaleTimeString()}`;
            updateOK = true;
            if (tid >= 0) {
                clearInterval(tid);
            }
        });
        // We're polling for results here ...
        tid = setInterval(function() {
            if (updateOK == true) {
                clearInterval(tid);
            }
        }, 1 * 1000);
    } else {
        people = Object.assign({}, default_people);
        render_form(div, people, form_src, form_init);
        let _status = div.querySelector('#status'),
            read = div.querySelector('#read');
        read.setAttribute("disabled", "disabled");
        _status.innerHTML = `Ready to create a new record ${(new Date()).toLocaleTimeString()}`;
    }

    window.People = people; //DEBUG
}(document, window));
