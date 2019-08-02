
# AndOr Object Scheme

**AndOr** does NOT enforce scheme shape for the JSON it stores.
If does add two specific attributes to the objects for management.
`._Key` and `._State`. The key holds the objects unique identifier
used in **AndOr** and queue holds the state of an object. E.g.
if an object is in a "published" state the value of `._State` would
be "published".

```json
    {
        ... /* object's scheme here */
        "_Key": "SOME_KEY_VALUE_HERE",
        "_State": "published"
    }       
```

