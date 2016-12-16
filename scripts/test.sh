#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

PACKAGE_DIR=/root/go/src/github.com/pivotal-cf/redisutils

while [[ $# -gt 0 ]]; do
  case $1 in
      --local)
        LOCAL=true
      ;;

      --)
        shift
        GINKGO_ARGS="$@"
        break
      ;;

      *)
        echo "Usage: test.sh [--local] [-- [ginkgo_arg, ...]]"
        exit 1
      ;;
  esac
  shift
done

if [ "$LOCAL" = true ]; then
  ginkgo ${GINKGO_ARGS}
  exit $?
fi

GINKGO_ARGS=${GINKGO_ARGS:-". -r --race --slowSpecThreshold=10"}

echo "Building docker image..."
$DIR/docker-build.sh > /dev/null

echo "Running Ginkgo test suites in docker..."
pushd $ROOT > /dev/null
docker run \
  -v $ROOT:/home/vcap/redisutils \
  -i -t cflondonservices/redisutils \
  /home/vcap/test.sh ${GINKGO_ARGS}
popd > /dev/null
