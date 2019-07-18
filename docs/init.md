
# init

The `init` verb creates empty dataset collections and also
"users.toml", "workflows.toml", and "andor.toml" TOML
files if they did not already exist.  You can list as many 
dataset collections as you want and the init command will 
attempt to initialize them one after the other.  If an error 
is encountered it will abort creating futher collections
after that.

## When to use init

If you are setting up a new repository you will need to initialize
collections. A typical use case would look like the following.

```bash
    AndOr init MyRepository.ds
```

This command would create the dataset collection and following files.

1. MyRepoisotory.ds, the collection holding repository data
2. andor.toml, the web service configuration use by `AndOr start`
3. users.toml, the users file by `AndOr start`
4. workflows.toml, the workflows file use by `AndOr start`

**AndOr** uses the three TOML files to configure how the web 
service works, known users and the workflows and capabilties
available to them.  Repository collections holds the content 
the API will provide. An **AndOr** provides access to 
the repositories listed in the "andor.toml" file. The users
workflows apply across repositories listed the "andor.toml" file..



