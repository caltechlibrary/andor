#!/usr/bin/env python3

import sys
from py_dataset import dataset

#
# This script loops through the poeple.ds dataset collection converting
# from our production data into demo data. This involves adding a 
# _State field as well as normalizing the field names we inheritted
# from the spreadsheet.
#

c_name = "people.ds"
keys = dataset.keys(c_name)
for key in keys:
    print(f"Importing {c_name}.{key}")
    data, err = dataset.read(c_name, key)
    if err != "":
        print(f"Error read {c_name}.{key}, {err}")
        sys.exit(1)
    # Make fieldname lower case
    obj = {
            "_Key": key,
            "_State": "deposit"
    }
    for field in data:
        fkey = field.lower()
        obj[fkey] = data[field]
    err = dataset.update(c_name, key, obj)
    if err != "":
        print(f"Error write {c_name}.{key}, {err}")
        sys.exit(1)
