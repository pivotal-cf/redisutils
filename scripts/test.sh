#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

PACKAGE_DIR=/root/go/src/github.com/pivotal-cf/redisutils

while [[ $# -gt 0 ]]; do
  case $1 in
      --unit-only)
        UNITONLY=true
      ;;

      --)
        shift
        GINKGO_ARGS="$@"
        break
      ;;

      *)
        echo "Usage: test.sh [--unit-only]"
        exit 1
      ;;
  esac
  shift
done

if [ "$UNITONLY" = true ]; then
  ginkgo monit/ --race
  exit $?
fi

GINKGO_ARGS=${GINKGO_ARGS:-". -r --race"}

echo "Building docker image..."
$DIR/docker-build.sh > /dev/null

echo "Running Ginkgo test suites in docker..."
pushd $ROOT > /dev/null
docker run \
  -v $ROOT:/home/vcap/redisutils \
  -i -t cflondonservices/redisutils \
  /home/vcap/test.sh ${GINKGO_ARGS}
popd > /dev/null
