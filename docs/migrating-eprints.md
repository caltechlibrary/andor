
# Migrating EPrints

Migration of any repository can be terribly challenging often with the
risk of lost metadata or nuanced metadata evolution. **AndOr** side
steps this issue by using EPrint's existing metadata scheme which
has proven itself for more than a decade. Caltech Library has an
existing set of tools for migrating content, as is, out of EPrints.
We use this [toolset](https://caltechlibrary.github.io/eprinttools)
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
    dataset init repository.ds
    ep -api $URL_TO_EPRINTS -dataset repository.ds \
        -export-with-docs -export all
```

That's it, you run those commands and wait. If you want to incrementally
update you collection you would swap out the `-export all` for an
approapriate export strategy.

**AndOr** is built on dataset collections so once the harvest is completed
you can fireup **AndOr** and test.

