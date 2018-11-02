#!/usr/bin/env bash

name=goproj-tmpl
scriptDir="scripts"
BASEDIR=$(dirname "$0")
cd $BASEDIR
cd ..

go fmt ./... && go install && ${name}
