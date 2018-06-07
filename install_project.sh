#!/bin/bash
mkdir -p pkg bin
npm install newman
go get github.com/dustin/go-humanize
go install github.com/dustin/go-humanize