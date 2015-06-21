#!/bin/bash

gofmt -w .

go run lfm-user.go -u goranche,poohica

go run lfm-tag.go -t rap | go run lfm-artist.go -a Drake
