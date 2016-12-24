#!/bin/bash

cp -r redisutils $HOME/redisutils
sudo -E -u vcap \
  PATH=$PATH:$GOPATH/bin \
  $HOME/redisutils/docker/test.sh . -r -v
