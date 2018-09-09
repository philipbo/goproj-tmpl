#!/usr/bin/env bash

scriptDir="scripts"
BASEDIR=$(dirname "$0")
cd $BASEDIR
cd ..

git pull
export GOPATH=/data/go

name=web-temp
cd /data/go/src/$name
go install
sudo supervisorctl restart $name
tail -f /data/go/logs/$name-error.log