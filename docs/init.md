
# init

The `init` verb creates empty dataset collections. You can
list as many dataset collections as you want it will sequencially
go through the list creating them one at a time. If there is
an error it will abort on the collectiont that encounters an
error (e.g. permissions to create directories and files).

## When to use init

If you are setting up a new repository you will need to initialize
collections. A typical use case would look like the following.

```bash
    AndOr init MyRepo.ds users.AndOr workflow.AndOr
```

This command would create the following dataset collections
in the following order.

1. MyRepo.ds, the collection holding repository data
2. users.AndOr, a dataset collection holding curitorial user information
3. workflows.AndOr, a dataset collection hold workflow definitions

**AndOr** requires two collections named "users.AndOr" and "workflows.AndOr". These define the curatorial users and workflows available to **AndOr**.
**AndOr** can support one or more additional "repository" collections.
Repository collections holds the content the API will provide.
They can be called anything you like except "users.AndOr" or
"workflows.AndOr".


