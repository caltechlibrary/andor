#!/bin/bash

echo "Checking for andor program"
ANDOR="bin/andor"
if [[ ! -f "$ANDOR" ]]; then
    make clean
    make
    make website
fi
echo "Starting $ANDOR with people-andor.toml"
cd demo  && "../$ANDOR" start people-andor.toml
