# Go Curabitur

Demo web application for using web socket on go language

## Prepare

GO TPATH

    export GOPATH=/path/to/project

BIN PATH

    export PATH=/path/to/project/bin:$PATH

## Development

Vendor

    make env

Asset

    npm install
    make dev-assets

Server

    make dev-server

## Release

Commands

    make env

    npm install
    make release
