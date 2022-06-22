#!/bin/bash

# get licenses if no cache exists
go install github.com/google/go-licenses@v1.2.1 && go-licenses save ./... --save_path third_party_licenses --force

# fetch dependencies
go get -d ./...
ls -la $HOME/go/pkg/mod
cp -r $HOME/go/pkg/mod .
