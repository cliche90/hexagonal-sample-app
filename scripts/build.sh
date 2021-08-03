#!/bin/bash

build_path="$(dirname "$0")/../cmd/server/"
go build "${build_path}"