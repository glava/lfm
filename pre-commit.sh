#!/bin/bash
go test ./lastfm

gofmt -w .

go build

go install

#smoke run

lfm -u goranche,poohica

lfm -t rap | lfm -a Drake
