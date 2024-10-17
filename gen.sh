#!/bin/zsh
go build -o orange main.go
chmod +x orange
./orange gen -c resources/config.dev.yaml -m open -d default -t user


