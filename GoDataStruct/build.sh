#!/bin/bash

CURDIR=`pwd`
export GOPATH=$GOPATH:$CURDIR
export GOBIN=$CURDIR

#go install
go install ./src/main/main.go
