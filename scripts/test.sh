#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

PACKAGE_DIR=/root/go/src/github.com/pivotal-cf/redisutils

while [[ $# -gt 0 ]]; do
  case $1 in
      --unit-only)
        UNITONLY=true
      ;;

      *)
        echo "Usage: test.sh [--unit-only]"
        exit 1
      ;;
  esac
  shift
done

if [ "$UNITONLY" = true ]; then
  ginkgo monit/
  exit $?
fi

echo "Building docker image..."
$DIR/docker-build.sh > /dev/null

echo "Running Ginkgo test suites in docker..."
docker run -i -t cflondonservices/redisutils $PACKAGE_DIR/docker/test.sh
