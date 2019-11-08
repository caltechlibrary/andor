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


#andor$(EXT): bin/andor$(EXT)


#bin/andor$(EXT): *.go cmd/andor/andor.go
#	go build -o bin/andor$(EXT) cmd/andor/andor.go

#build: $(PROJECT_LIST) libandor 

#install: 
#	env GOBIN=$(GOPATH)/bin go install cmd/andor/andor.go

website: page.tmpl README.md nav.md INSTALL.md LICENSE css/site.css
	#cp -vR scripts demo/htdocs/
	./mk_website.py $(baseurl)

#test: clean bin/andor$(EXT)
#	go test

reset:
	python3 development_reset.py

cleanweb:
	if [ -f index.html ]; then rm *.html; fi
	if [ -f docs/index.html ]; then rm docs/*.html; fi

clean: 
	if [ -d bin ]; then rm -fR bin; fi

py_code:
	cp *.py dist/
	cp -fR py_libdataset dist/

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/

update_version:
	./update_version.py --yes

release: clean distribute_docs py_code

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
