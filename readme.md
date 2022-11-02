# BTV Lite API

BTV Lite API is a GQL api that share the GraphQL schema with the main BTVPlatform API.
It is backed by a readonly sqlite database that ban be obtained from the main API.

**WARNING**: This is experimental and not stable!

## Goals

The goal is to provide a lightweight system for making BTVPlatform clients work offline

# Development

## Requirements

* [GoLang >=1.18](https://go.dev/)
* [gqlgen >=0.17.20](https://github.com/99designs/gqlgen)
* [sqlc >=1.15](https://github.com/kyleconroy/sqlc)

### Optional stuff

* `make` if you want to use the `Makefile` commands
* [fresh](https://github.com/pilu/fresh) for automatic rebuilds

## Run dev

Note: requires `fresh`


```shell
make run DB_PATH=/path/to/db.sqlite3
```

## Generate code for updated GQL schema

This command assumes that you have the [BTV Platform repository](https://github.com/bcc-code/brunstadtv) checked out at
`../brunstadtv/` relative to this readme file.

If you need to change the source of the GQL schema definition files, look in `./gqlgen.yml` file in this directory.

```shell
make gqlgen
```

## Update SQL statements

```shell
make sqlc
```

## Compile

```shell
make build # <-- Current arch
make build_all # <-- Cross compile for all specified OSs and architectures
```

Results are in the `./dist` folder

# Getting the DB

Send the following GQL query to the BTV Platform

```graphql
query {
	export {
		dbVersion
		url
	}
}
```

The returned URL  will be valid for 1 hour.

Note that the information is filtered based on the requesting user. This means that without a valid bearer token you 
will only receive public data.