
# Setting Up Andor

**AndOr** is experimental, a proof of concept so you're expected
to clone the git repo and compile it before following the steps
to setup and run it.

# Setting up Andor

1. Clone the AndOr repository
2. Make AndOr
3. Create a dataset collection to hold our objects and attachments
3. Create a dataset collection called `workflow.AndOr`
4. Create a dataset collection called `users.AndOr`
5. Generate an admin-user.json document
6. Generate an admin-workflow.json document
7. Add admin-workflow.json to workflow.AndOr
8. Add admin-user.json to users.AndOr
9. Start up AndOr and view with your web browser

## Example: A simple repository using your ORCID for user id

```bash
    git clone https://github.com/caltechlibrary/AndOr.git
    cd AndOr
    make
    # repository.ds can be named something else if you like
    # the name/path to look for is set in the configuration.
    bin/AndOr init repository.ds
    bin/AndOr init workflow.AndOr
    bin/AndOr init users.AndOr
    # We're only creating a local secrets collection for demo purposes
    # Normally you'd authenticate through another service like OAuth 2.
    bin/AndOr init secrets.AndOr

    # AndOr create-user will prompt for needed information.
    bin/AndOr create-user
    # AndOr create-workflow will prompt for the needed information.
    bin/AndOr create-workflow
    # AndOr config-service will prompt for the needed information.
    bin/AndOr config-service

    # Startup the AndOr web service (we're hosting a single object
    # collection.
    bin/AndOr host repository.ds 
```

**AndOr** by default runs at http://localhost:8123. You can use
additional options to control the port, hostname and protocol.

The is a demo system, a proof of concept. Don't expect this to work! 
It uses a local user id to secret collection. Normally you'd use
an OAuth source to authenticate with though the local secrets is 
convienent for localhost development environment.

