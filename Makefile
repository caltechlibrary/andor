#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = andor

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

PROJECT_LIST = AndOr

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


AndOr$(EXT): bin/AndOr$(EXT)


bin/AndOr$(EXT): *.go cmd/AndOr/AndOr.go
	go build -o bin/AndOr$(EXT) cmd/AndOr/AndOr.go

build: $(PROJECT_LIST) libAndOr

install: 
	env GOBIN=$(GOPATH)/bin go install cmd/AndOr/AndOr.go

website: page.tmpl README.md nav.md INSTALL.md LICENSE css/site.css
	./mk_website.py $(baseurl)

test: clean bin/AndOr$(EXT)
	go test

cleanweb:
	if [ -f index.html ]; then rm *.html; fi
	if [ -f docs/index.html ]; then rm docs/*.html; fi

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d repository.ds ]; then rm -fR respository.ds; fi
	if [ -d test_repo.ds ]; then rm -fR test_repo.ds; fi
	if [ -f andor.toml ]; then rm andor.toml; fi
	if [ -f workflows.toml ]; then rm workflows.toml; fi
	if [ -f users.toml ]; then rm users.toml; fi

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/AndOr
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/AndOr.exe
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/AndOr
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/AndOr cmd/AndOr/AndOr.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	#FIXME: use go.mod instead
	#bash package-versions.bash > dist/package-versions.txt

update_version:
	./update_version.py --yes

release: clean AndOr.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	./mk_website.py
	bash publish.bash

FORCE:
