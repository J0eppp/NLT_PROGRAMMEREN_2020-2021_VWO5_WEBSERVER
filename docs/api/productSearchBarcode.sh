#!/bin/bash

curl localhost:8080/api/v1/product?product=8718906445338 | jq --tab
