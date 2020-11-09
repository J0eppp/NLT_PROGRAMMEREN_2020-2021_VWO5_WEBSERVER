#!/bin/bash

curl localhost:7070/api/v1/product?product=8718906445338 | jq --tab
