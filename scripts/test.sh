#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

while [[ $# -gt 0 ]]; do
  case $1 in
      --local)
        LOCAL=true
      ;;

      --skip-docker-build)
        SKIP_BUILD=true
      ;;

      --)
        shift
        GINKGO_ARGS="$@"
        break
      ;;

      *)
        echo "Usage: test.sh [--local] [--skip-docker-build] [-- [ginkgo_arg, ...]]"
        exit 1
      ;;
  esac
  shift
done

if [ "$LOCAL" = true ]; then
  UNIT_TESTS="monit redisserver redisserver/config iredis"
  GINKGO_ARGS=${GINKGO_ARGS:-"--race ${UNIT_TESTS}"}
  ginkgo $GINKGO_ARGS
  exit $?
fi

GINKGO_ARGS=${GINKGO_ARGS:-". -r --race --slowSpecThreshold=15"}

if [ "$SKIP_BUILD" != true ]; then
  echo "Building docker image..."
  $DIR/docker-build.sh > /dev/null
fi

echo "Running Ginkgo test suites in docker..."
pushd $ROOT > /dev/null
docker run \
  -v $ROOT:/home/vcap/redisutils \
  -i -t cflondonservices/redisutils \
  /home/vcap/test.sh ${GINKGO_ARGS}
popd > /dev/null
