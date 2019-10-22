
# Action items

## Bugs

## Next

+ [ ] Evaluate impact of dropping explicit user/role model and implement andor as part of libdataset where Python provides the Web Service infrastructure rather than a stand alone services written in Go.

## Someday, Maybe ideas

+ [ ] Add configurable web hooks for create, update and delete actions
+ [ ] Expand And/Or command line verbs to either generate JSON configurations of a related andor system's collection
+ [ ] Dropping support form TOML and YAML to reduce external dependencies
+ [ ] File watch collection.json changes it needs to be reloaded into running AndOr instance, this will allow external processes to update frames or modify records
+ [ ] Remove basic auth support from And/Or service in favor of OAuth 2 tokens, demo using ORCID to identify curators
+ [ ] Migrate users.toml, roles.toml into JSON records for managing curation, the collection to use is defined in system json file used when calling `andor start`
+ [ ] User/roles to be defined as their own collection that is not publically visible but is used by And/Or for access
+ [ ] Change "lookup" to reflect `search for {{field_name}}` or `Go to {{field_value}}`
+ [ ] Consolidate Look up link and link to object's source
+ [ ] orcid, author_id, thesis_id and cl_people_id need to have a copy field button so they can be easily copied to be pasted into EPrints
+ [ ] Have create link take you to a create page, rename "create" button to save, save button should take you back to a read only view with link to edit item
+ [ ] Rename "update" to save, the edit page should be separate from the save page, when save is pressed go back to a read only view with link to edit item
+ [ ] Editing the collection when And/Or is running to trigger a refresh of collection.json in memory
    + add a file listener for collection.json to trigger refresh, it needs to pass through the mutex
+ [ ] Need a page to copy a record to a new record with a new cl_people_id
+ [ ] Need a page to merge to cl_people_id into one record with a field picker like in JabRef
+ [ ] if cl_people_id is empty the lookup to warning if you're going to overwrite a record

## Ideas

+ [ ] If an object is being "edited" (read before an update) indicate that in the UI of someone else logged in an editing objects - Stephen's suggestion
+ [ ] Implement search with Lunr for prototype
+ [ ] Integrate with external services (e.g. orcid.org, datacite.org)
+ [ ] Integrate with ArchivesSpace

## Blue Sky ideas

+ [ ] OAuth, OpenID, Shibboleth, and LDAP natively in andor service
+ [ ] Add metadata versioning
+ [ ] Add attachment support
+ [ ] Add issue tracker integration
+ [ ] Create a public website view
+ [ ] Implement next ID service for generating unique ids
+ [ ] Implement Thesis Deposit/ETD like service

## Completed

+ [x] The tabindex values is interacting with the next field focus negatively, find a different solution
    + remove tabindex
+ [x] The placeholder text obscures the field values and can't easily be styled
    + replace with title attribute.
+ [x] view methods should only be defined in a custom repository's JS, andor.js should only support retrieval methods specific to the And/Or JSON API.
+ [x] init should generate a roles.toml and users.toml for editing if they do not exist
