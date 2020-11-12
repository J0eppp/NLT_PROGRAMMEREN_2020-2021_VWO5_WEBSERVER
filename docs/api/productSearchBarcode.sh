#!/bin/bash

curl localhost:8080/api/v1/product?product=8718906445338 | jq --tab
curl localhost:8080/api/v1/product?product=8710400197652 | jq --tab
