
# add-workflow

Workflows are a central organzation of **AndOr** for
permissioning as well as organizating curitorial tools.
The `add-workflow` verb will trigger an interactive
session for defining a workflow. In addition to
the interactive mode of `add-workflow` you can
supply a ready made JSON object and workflow name.
This makes it handle to script setup of repositories
built on **AndOr**.

## Interactive usage

In this example we are going to create an "editor" workflow
and have **AndOr** prompt interactive for settings
of "editor"

```bash
    bin/AndOr add-workflow editor
```

Or for loading an predefined workflow called "editor" from
a file called editor.json.

```bash
    bin/AndOr add-workflow editor editor.json
```

The workflows are stored in a 
[dataset](https://github.com/caltechlibrary/dataset) 
collection called "workflows.AndOr". You can export a 
workflow from one deployment to another by using 
the **dataset** command. In this example we're
retrieving a workflow called editor from the dataset
collection stored at `~/RadioPlays/workflows.AndOr`
and importing them into ~/Film/workflows.AndOr`.

```bash
    # Export our workflow from RadioPlays
    dataset read ~/RadioPlays/workflows.AndOr editor \
        > ~/Film/editor.json
    cd ~/Film
    # Import our workflow into Film's workflows.AndOr
    AndOr add-workflow editor editor.json
```


