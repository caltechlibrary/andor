
# Migrating EPrints

Migration of any repository can be terribly challenging often with the
risk of lost metadata or nuanced metadata evolution. **AndOr** side
steps this issue by using EPrints's existing metadata scheme which
has proven itself for more than a decade. Caltech Library has an
existing set of tools for migrating content, as is, out of EPrints.
We use this [tool set](https://caltechlibrary.github.io/eprinttools)
regularly and it includes the ability to migrate not just the metadata
but also the attach documents.

This is what is required to migrate an EPrints' collection to 
**AndOr** based system.

1. Create an account (if one doesn't access) in EPrints with full access to the EPrints REST API
2. Use the [ep](https://caltechlibrary.github.io/eprinttools/docs/ep/) command line program to export the repository along with its documents.

In our proof of concept we've previously setup our EPrints repository
to have content harvested. We use an existing account (in this
example `$URL_TO_EPRINTS` would be the EPrints host and any authentication needed such as an account name and password).
Here's the steps with existing Caltech Library tools to migrate an
EPrints repository into a dataset collection.

```bash
    # Create an empty AndOr deployment
    bin/AndOr init repository.ds users.AndOr workflow.AndOr
    # Import our metadata and files from an existing EPrints
    ep -api $URL_TO_EPRINTS -dataset repository.ds \
        -export-with-docs -export all
```

That's it, you run those commands and wait. When the import
is complete you can add some workflows and users to curate your
new repository.

You could also using an incrementally migration/update from
the EPrints collection by swapping out the `-export all` for an
appropriate export strategy. In principle you could run the two
in parallel assuming the data flow was one direction (i.e.
from EPrints to AndOr).

**AndOr** is built on dataset collections. If there is a
workflows.AndOr, users.AndOr and a dataset collection
**AndOr** is ready to fire up and test.

