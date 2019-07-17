
# Setting Up AndOr

**AndOr** is a proof of concept so you're expected
to clone the git repo and compile it before following the steps
to setup and run it.

# Steps

1. Clone the AndOr repository
2. Run Make to compile it
3. Initialize dataset collections for user, workflows and repository
    + also creates an example workflows.toml and users.toml you can edit
4. Harvest content from EPrints repository being migrated
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
    # workflows.AndOr and users.AndOr must user those names.
    bin/AndOr init repository.ds workflows.AndOr users.AndOr

    # We need some workflows. We create them in a TOML file
    # and in this example the file is called "workflows.toml".
    # We then "load" the workflows into AndOr.
    # NOTE: Loading only adds/updates workflows in AndOr.
    $EDITOR workflows.toml
    bin/AndOr load-workflow workflows.toml

    # AndOr create users is done first by create/editing a TOML
    # file, in this example users.toml, then "loading" 
    # it into AndOr. 
    # NOTE: Loading ONLY adds/updates users in AndOr.
    $EDITOR users.toml
    bin/AndOr load-user user.toml

    # Configure our AndOr web service (e.g. set hostname, port, 
    # protocol, collection name(s))
    bin/AndOr config > webservice.toml
    $EDITOR webservice.toml

    # Startup the AndOr web service with webservice.toml 
    bin/AndOr start webservice.toml
```

**AndOr** by default runs at http://localhost:8248. You can 
change generating a TOML based config file, editing it and then
using it to start AndOr.

This is describing a proof of concept system. Don't expect 
this to work yet!

