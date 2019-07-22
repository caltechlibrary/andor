
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
    # the init command will create users.toml, workflows.toml
    # and andor.toml if they don't exist.
    bin/AndOr init repository.ds 

    # We need some workflows. We create them in a TOML file
    # called "workflows.toml".
    # When we start AndOr it "loads" three files --
    # users.toml, workflows.toml and andor.toml.
    $EDITOR workflows.toml

    # AndOr creating/managing users is done first by editing 
    # a TOML file called "users.toml". It will get "loaded" when
    # you start the AndOr service.
    $EDITOR users.toml

    # Startup the AndOr web service with webservice.toml 
    bin/AndOr start
```

**AndOr** by default runs at http://localhost:8248. You can 
change it by updating your "andor.toml" file created with
`bin/AndOr init`.

This is describing a proof of concept system. Don't expect 
this to work yet!

