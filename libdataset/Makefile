#
# Simple Makefile for building C-Shared library for Python3 module.
#
LIB_NAME = libdataset

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

ARCH = x86_64

OS = $(shell uname)

EXT = .so
ifeq ($(OS), Windows)
	EXT = .dll
	ARCH = x86_64
	os = windows
endif
ifeq ($(OS), Darwin)
	EXT = .dylib
	ARCH = $(shell arch)
	OS = macosx
endif
ifeq ($(OS), Linux)
	EXT = .so
	ARCH = $(shell arch)
	OS = linux
endif
ifeq ($(ARCH), i386)
	ARCH = amd64
endif
ifeq ($(ARCH), x86_64)
	ARCH = amd64
endif


build: $(LIB_NAME)$(EXT)

$(LIB_NAME)$(EXT): lib/$(LIB_NAME)$(EXT)
	
lib/$(LIB_NAME)$(EXT): $(LIB_NAME).go
	go build -buildmode=c-shared -o "lib/$(LIB_NAME)$(EXT)" "$(LIB_NAME).go"


clean:
	if [ -f "lib/$(LIB_NAME).dll" ]; then rm "lib/$(LIB_NAME).dll"; fi
	if [ -f "lib/$(LIB_NAME).dylib" ]; then rm "lib/$(LIB_NAME).dylib"; fi
	if [ -f "lib/$(LIB_NAME).so" ]; then rm "lib/$(LIB_NAME).so"; fi
	if [ -f "lib/$(LIB_NAME).h" ]; then rm "lib/$(LIB_NAME).h"; fi

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)


.FORCE:
