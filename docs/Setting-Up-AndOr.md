
# Setting Up AndOr

**AndOr** is a proof of concept so you're expected
to clone the git repo and compile it before following the steps
to setup and run it.

# Steps

1. Clone the AndOr repository
2. Run Make to compile it
3. Initialize dataset collections for user, workflows and repository
4. Harvest content from EPrint repository being migrated
5. Add workflow(s) 
6. Add user(s)
7. Start up AndOr and view with your web browser

## Example: A simple empty repository

```bash
    git clone https://github.com/caltechlibrary/AndOr.git
    cd AndOr
    make
    # repository.ds can be named something else if you like
    # the name/path to look for is set in the configuration.
    bin/AndOr init repository.ds

    # AndOr create-user will prompt for needed information.
    bin/AndOr create-user
    # AndOr create-workflow will prompt for the needed information.
    bin/AndOr create-workflow
    # AndOr config-service will prompt for the needed information.
    bin/AndOr config-service

    # Startup the AndOr web service (we're hosting a single object
    # collection.
    bin/AndOr start
```

**AndOr** by default runs at http://localhost:8248. You can use
additional options to control the port, hostname and protocol.

This is describing a proof of concept system. Don't expect 
this to work!

