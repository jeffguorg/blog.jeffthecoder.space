#!/usr/bin/env bash

# package resources into go code
pushd internal/template
rm bindata.go -fv
go-bindata -o bindata.go -pkg template .
popd