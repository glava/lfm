#!/bin/bash

gofmt -w .

go build 

go install

lfm -u goranche,poohica

lfm -t rap | lfm -a Drake
