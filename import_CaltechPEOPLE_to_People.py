from libdataset import dataset

if __name__ == '__main__':
    c_name1 = 'CaltechPEOPLE.ds'
    c_name2 = 'People.ds'
    keys = dataset.keys(c_name1)
    for key in keys:
        obj1, err = dataset.read(c_name1, key)
        if err != '':
            print(f'WARNING: read failed, {key} in {c_name1}, {err}')
        else:
            obj2 = {}
            for attr1 in obj1:
                attr2 = attr1.lower().replace(' ', '_').replace('(', '').replace(')', '')
                if attr2.count('_link') == 0 and attr2.count('autogenerate') == 0 and not attr2 in [ 'cn', 'family', 'given', '_key' ]:
                    obj2[attr2] = obj1[attr1]
                elif attr2 in [ 'caltech', 'alumn', 'jpl', 'faculty' ]:
                    obj2[attr2] = True
                    if obj2[attr2] == '':
                        obj2[attr2] = False
            if dataset.key_exists(c_name2, key):
                err = dataset.update(c_name2, key, obj2)
            else:
                err = dataset.create(c_name2, key, obj2)
            if err != '':
                print(f'WARNING: write failed, {key} in {c_name2}, {err}')
