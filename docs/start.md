
# Starting the web service

Example 

```bash
    AndOr start
    AndOr start /etc/andor.toml
```

**AndOr** web service is started with the "start" verb.
By default it looks for "andor.toml" to load its configuration.
You may also provide a specific path to that file. That file
is used to setup the web service's hostname, port, path
to the repositories, and where to find the workflows and users
TOML files. If something is wrong (e.g. can't find a toml file,
or they don't parse correctly) it will abort start up with an
exit code of 1.  If started you'll see an "OK" printed followed
by logging of web requests.

