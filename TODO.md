
# Action items

## Bugs

+ [x] init should generate a roles.toml and users.toml for editing if non exist

## Address Missing features

+ [ ] EPrints versions metadata as well as attachments
    + Could use a simple diff on a pretty printed object
    + Modify dataset to automatically version attachments
    + Modify dataset to automatically save a diff of JSON doc, e.g. OBJECT_ID.diff.YYYYMMDDHHMM
+ [ ] EPrints has a issues system
    + Probably link out Jira or GitHub if we're going to support issues
    + Worst case add another collection called "issues" and just store stuff there with the same ID as the Object record.
+ [ ] If an object is being "edited" (read before an update) indicate that in the UI of someone else logged in an editing objects - Stephen's suggestion

## Next

+ [x] Write up basic approach
+ [x] Write up use cases demonstrating how we address minimal requirements
+ [x] Write up some use cases based on existing EPrints usage
+ [x] Document proposed schema (e.g. role, AndOr user)
+ [x] Pick a simple auth mechanism for use in proof of concept
+ [ ] Read up on headers passed from BasicAUTH so we can use that for Oral Histories authentication and user id mapping
+ [x] Evaluate using diff to created versioned metadata
    + https://github.com/sergi/go-diff
    + https://godoc.org/golang.org/x/perf/internal/diff
+ [ ] Implement a proof of concept web service (AndOr) supporting crud operations on collection objects, listing objects and user/role rule enforcement
+ [ ] Implement search with Lunr for prototype
+ [ ] Design basic UI in HTML and JavaScript

## Someday, Maybe

+ [ ] Wrap dataset to accept TOML files as import files
+ [ ] Add metadata versioning
+ [ ] Add issue tracker integration
+ [ ] Create a public website view
+ [ ] Integrate with ArchivesSpace
+ [ ] Integrate with external services (e.g. orcid.org, datacite.org)
+ [ ] Implement next ID service for generating unique ids
+ [ ] Implement Thesis Deposit/ETD like service
