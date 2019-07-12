
# AndOr Object Scheme

**AndOr** does really enforce an scheme shape for the JSON it stores.
It does add on fields to all records stored via its API. That field
is "\_Queue". It holds the current workflow queue assignment of the
object. For a newly created object that assignment comes from the
[AndOr user scheme](User-Schema.html). It will change as the object
moves through a workflow.  **dataset** the storage engine of **AndOr** also adds a field to the objects' schema. These include "\_Key". The **AndOr** JSON API can return the object clean using the ?clean=true URL parameter.

```json
    {
        ... /* object's scheme here */
        "_Key": "SOME_KEY_VALUE_HERE",
        "_WorkQueue": "published"
    }       
```

