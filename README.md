vasques
---

[![Build Status](https://cloud.drone.io/api/badges/scbizu/vasques/status.svg)](https://cloud.drone.io/scbizu/vasques)

A Go tool to convert `.json` files to `.proto`(mainly proto3)files.

## READ AT FIRST

vasques chooses the [buf's styleguide](https://buf.build/docs/style-guide) as the generated proto3 coding style since buf provides the robust protobuf file linting tools .

## Currently support types

|JSON Types|Protobuf Types|Additional Notes|
|--|--|--|
|string|string/int64|string would be treat as int64 at some cases due to the javascript's  precision|
|int|int32||
|float|double||
|boolean|bool||
|Object|Message||

## Install

```
# install from source
$ go get github.com/scbizu/vasques #Can be used as a binary directly with your $GOPATH set
# or outside $GOPATH/bin and GO111MODULE = ON
# install homebrew-gomod at first
$ brew install FiloSottile/gomod/brew-gomod
$ brew gomod  github.com/scbizu/vasques
```

or download the complied release from GitHub Release

## Usage 

```
$ vasques --json path/to/your_json_file.json   -d path/to/your_proto_file.proto
```

Explore more example json to proto files in the example directory


