#!/usr/bin/env bash

name=web-temp
scriptDir="scripts"
BASEDIR=$(dirname "$0")
cd $BASEDIR
cd ..

go fmt ./... && go install && ${name}
