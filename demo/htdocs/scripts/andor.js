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
        console.log("DEBUG url_path", url_path);
        let p = url_path.split("/")
        if (p.length > 1) {
            console.log("DEBUG collection name", p[1], "parts ->", p);
            return p[1];
        }
        return "unknown";
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

        console.log(`DEBUG updateObject(${collection_name}, ${key}, ...)`);
        if (collection_name === undefined || collection_name === "") {
            console.log('WARNING: AndOr.updateObject() called without collection name');
            return false;
        }
        if (key === undefined || key === null || key === "") {
            console.log('WARNING: AndOr.updateObject() called without key', key, obj);
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
