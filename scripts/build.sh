#!/usr/bin/env bash

# build binary

BASEDIR=$(dirname "$0")
echo "BASEDIR=$BASEDIR"
cd $BASEDIR
cd ..

name=web-temp

case $1 in
-n | -name)
    if [ "$2" == "" ]
    then
        echo "name required"
        exit 1
    fi
    name=$2
    ;;
*)
    echo "default binary name $name"
esac

echo "cross compile"

buildTime=`date "+%Y-%m-%d_%H:%M:%S"`
gitHash=`git rev-parse HEAD`
gitBranch=`git rev-parse --abbrev-ref HEAD`
goVersion=`go version`

echo "BuildTime: $buildTime"
echo "GitHash: $gitHash"
echo "GitBranch: $gitBranch"
echo "GoVersion: $goVersion"


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $name -ldflags "\
    -X 'main.BuildTime=$buildTime' \
    -X 'main.GitHash=$gitHash' \
    -X 'main.GitBranch=$gitBranch' \
    -X 'main.GoVersion=$goVersion'"

#version="[version]_`whoami`_`date "+%Y-%m-%d__%H:%M:%S"`_`git rev-parse --abbrev-ref HEAD`_`git rev-parse HEAD`"
#echo "$version"

##GOOS=linux or GOOS=darwin
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $name -ldflags "-X version.Version=$version"
