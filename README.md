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
