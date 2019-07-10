
# Action items

## Bugs

## Next

+ [ ] Figure out a simple auth mechanism to use for the proof of concept
    + [ ] Evaluate using OAuth 2 with library's Google Apps for Educatin
    + [ ] Read up on go-cloud package for easy authentication options
    + [ ] If not assume we're behind a web server that can provide the authenticaition mechanism like Shibboleth (e.g. Apache, NginX)
+ [ ] Read up on channels so I can create multi-user safe access to our dataset collections
+ [ ] andor-init intializes the necessarily dataset collections to run andor
+ [ ] andor-genuser generates a JSON document which can be added to users.andor
+ [ ] andor-genworkflow generates a JSON document which can be added to workflows.andor
+ [ ] andor (web service)
    + [ ] JSON read API calls
    + [ ] JSON write API calls
    + [ ] Web UI based on content in htdocs

## Someday, Maybe

+ [ ] Wrap dataset to accept TOML files as import files
