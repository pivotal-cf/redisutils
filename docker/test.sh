#!/bin/bash

REDISUTILS_DIR=$GOPATH/src/github.com/pivotal-cf/redisutils

cp -r $HOME/redisutils/vendor/* $GOPATH/src
mkdir -p $REDISUTILS_DIR
cp -r $HOME/redisutils/* $REDISUTILS_DIR

cd $REDISUTILS_DIR
shift
ginkgo "$@"
