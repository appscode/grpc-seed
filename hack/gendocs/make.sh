#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/plow/hack/gendocs
go run main.go
popd
