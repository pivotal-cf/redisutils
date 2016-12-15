#!/bin/bash

REDISUTILS_DIR=$GOPATH/src/github.com/pivotal-cf/redisutils

while [[ $# -gt 0 ]]; do
  case $1 in
      --verbose)
        VERBOSE="-v"
      ;;

      *)
        echo "Usage: test.sh [--verbose]"
        exit 1
      ;;
  esac
  shift
done

cp -r $HOME/redisutils/vendor/* $GOPATH/src
mkdir -p $REDISUTILS_DIR
cp -r $HOME/redisutils/* $REDISUTILS_DIR

cd $REDISUTILS_DIR
ginkgo $VERBOSE -r --race
