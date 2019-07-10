
# Install Andor

**andor** is experimental, a proof of concept so you're expected
to clone the git repo and compile it before following the steps
to setup and run it.

# Setting up Andor

1. Clone the andor repository
2. Make andor
3. Create a dataset collection to hold our objects and attachments
3. Create a dataset collection called `workflow.andor`
4. Create a dataset collection called `users.andor`
5. Generate an admin-user.json document
6. Generate an admin-workflow.json document
7. Add admin-workflow.json to workflow.andor
8. Add admin-user.json to users.andor
9. Start up andor and view with your web browser

## Example: A simple repository using your ORCID for user id

```bash
    git clone https://github.com/caltechlibrary/andor.git
    cd andor
    make
    dataset init repository.ds
    dataset init workflow.andor
    dataset init users.andor
    dataset init secrets.andor

    # andor-genuser will prompt for needed information.
    bin/andor-genuser admin-user.json
    # andor-genworkflow will prompt for the need information.
    bin/andor-genworkflow admin-workflow.json

    dataset create -i admin-workflow.json workflow.andor admin
    dataset create -i admin-user.json user.andor admin

    bin/andor repository.ds 
```

**andor** by default runs at http://localhost:8123. You can use
additional options to control the port, hostname and protocol.

The is a demo system, a proof of concept. Don't expect this to work! 
It uses a local user id to secret collection. Normally you'd use
an OAuth source to authenticate with though the local secrets is 
convienent for localhost development environment.

