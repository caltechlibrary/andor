
# Action items

## Bugs

## Next

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
