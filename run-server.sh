#!/bin/bash
set -e

bash ./build.sh

ENVIRONMENT=${1:-development}

./cmd/go-gin-boilerplate --config=$ENVIRONMENT
