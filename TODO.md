
# Action items

## Bugs

+ [ ] Editing the collection when And/Or is running to trigger a refresh of collection.json in memory
+ [ ] view methods should only be defined in a custom repository's JS, andor.js should only support retrieval methods specific to the And/Or JSON API.
+ [x] init should generate a roles.toml and users.toml for editing if they do not exist
+ [ ] adding, changing the cl_people_id should trigger a lookup and set the create/save button state
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
