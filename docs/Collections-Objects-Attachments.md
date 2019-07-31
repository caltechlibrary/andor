+++
title = "Collections, Objects and Attachments"
markup = "mmark"
+++


# Collections, Objects and Attachments

Collections hold objects. Objects are expressed as JSON
documents. They have a unique ID in the collection.
The minimum object for dataset includes a `._Key` 
attribute and value. The `._Key` reflects the unique id
in the collection. The minimum And/Or object holds a `._Key` 
and `._Queue` attributes. The later is used to map an object
to a workflow and apply the capabilities expressed in And/Or's
roles.

Objects as JSON documents but sometimes that isn't enough.
It is commont to want to the JSON object as metadata about
something else like a PDF document. Since encoding documents
like PDF into valid JSON is extremely ineffecient And/Or
via dataset supports the concept of "attachments". This is
like attachments in email. In email you might write a memo
to someone and include the files that the memo is about. In
And/Or and dataset you might "attach" the files to the JSON
object. The attachments are versioned using a semver[^1].
In this way dataset and And/Or form the foundation on which
most repository systems are built.

## The plumbing

And/Or is based on Caltech Library's [dataset](https://caltechlibrary.github.io/dataset) tool. dataset is a simple JSON
object store organized as a directory containing a pairtree[^2].
Each brand of the pair tree contains an object and any "attached"
files associated with he object. You can think of this is 
Fedora light[^3]. The dataset tool provides some basic CRUD[^4]
operations as well as a few extended concepts like key lists
and data frames. And/Or is inteneded to be a light weight
web service on top of dataset's collection(s). Where the
dataset tool is single user And/Or provides multi-user support
because it is a web service and can use the simpler dataset
package as if it was a single user. In effect it is a proxy
of many users behaving as a single coordated users from the
point of view of the dataset collection.

[^1]: Semver is short for semantic versioning, see https://semver.org/

[^2]: Pairtrees are a way of organizing files on disc by mapping a key to a directory path that avoids too many files in the file directory, see https://confluence.ucop.edu/display/Curation/PairTree

[^3]: Fedora (not to be confused with RedHat), is an enterprise repository services offered by Duraspace, see https://duraspace.org/fedora/
