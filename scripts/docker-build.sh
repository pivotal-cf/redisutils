#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

pushd $ROOT
docker build -t cflondonservices/redisutils .
popd
