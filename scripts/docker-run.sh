#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

pushd $ROOT > /dev/null
docker run -i -t cflondonservices/redisutils
popd > /dev/null
