#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = andor

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).py | cut -d\'  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

PROJECT_LIST = andor

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


build: libdataset config

website: page.tmpl README.md nav.md INSTALL.md LICENSE css/site.css
	python3 mk_website.py $(baseurl)

test: clean
	python3 test_libdataset.py

run:
	flask run

reset:
	python3 development_reset.py

libdataset: libdataset/lib/libdataset.h

libdataset/lib/libdataset.h:
	cd libdataset && $(MAKE)

config:
	python3 andor-setup.py Staff.ds Roles.ds People.ds
	dataset create Roles.ds 'Admin' '{"users": {"create": true, "read": true, "update": true, "delete": true}, "roles": {"create": true, "read": true, "update": true, "delete": true}, "objects": {"create": true, "read": true, "update": true, "delete": true}}'
	dataset create Roles.ds 'Editor' '{"users": {"create": false, "read": false, "update": false, "delete": false}, "roles": {"create": false, "read": false, "update": false, "delete": false}, "objects": {"create": true, "read": true, "update": true, "delete": true}}'
	dataset create Roles.ds 'Depositor' '{"users": {"create": false, "read": false, "update": false, "delete": false}, "roles": {"create": false, "read": false, "update": false, "delete": false}, "objects": {"create": true, "read": true, "update": false, "delete": false}}'
	python3 andor-admin.py add-user admin "$(USERNAME)@localhost" "Repository Admin"
	python3 andor-admin.py assign-role admin Admin
	python3 andor-admin.py password admin

cleanweb:
	if [ -f index.html ]; then rm *.html; fi
	if [ -f docs/index.html ]; then rm docs/*.html; fi

clean: 
	if [ -f libdataset/lib/libdataset.h ]; then rm -fR libdataset/lib/*; fi
	if [ -d libdataset/__pycache__ ]; then rm -fR libdataset/__pycache__; fi

update_version:
	./update_version.py --yes

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	./mk_website.py
	bash publish.bash

spellcheck:
	aspell -c README.md
	aspell -c INSTALL.md
	aspell -c INSTALL.md
	aspell -c README.md
	aspell -c TODO.md
	aspell -c docs/Object-Scheme.md
	aspell -c docs/Oral-Histories-as-Proof-of-Concept.md
	aspell -c docs/Queue-Scheme.md
	aspell -c docs/Reference.md
	aspell -c docs/Schedule.md
	aspell -c docs/Setting-Up-AndOr.md
	aspell -c docs/User-Scheme.md
	aspell -c docs/Workflow-Scheme.md
	aspell -c docs/Workflow-Use-Cases.md
	aspell -c docs/add-user-workflow.md
	aspell -c docs/check.md
	aspell -c docs/config.md
	aspell -c docs/index.md
	aspell -c docs/init.md
	aspell -c docs/list-users.md
	aspell -c docs/list-workflow.md
	aspell -c docs/load-user.md
	aspell -c docs/load-workflow.md
	aspell -c docs/migrating-eprints.md
	aspell -c docs/remove-user-workflow.md
	aspell -c docs/remove-user.md
	aspell -c docs/remove-workflow.md
	aspell -c docs/start.md

FORCE:
