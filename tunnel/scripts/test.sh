#!/bin/bash -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $( dirname $DIR ) )

function export_test_env_vars {
  export REDIS_HOST=$( docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-server )
  export REDIS_PORT="6379"
  SSH_ADDRESS=$( docker port ssh-server 22 )
  IFS=: read SSH_HOST SSH_PORT <<< "${SSH_ADDRESS}"
  export SSH_HOST=${SSH_HOST}
  export SSH_PORT=${SSH_PORT}
  export SSH_USER="vcap"
  export SSH_PASSWORD="funky92horse"
}

function run_tests {
  pushd $ROOT/tunnel > /dev/null
    ginkgo -v .
  popd > /dev/null
}

$ROOT/tunnel/scripts/containers.sh create

export_test_env_vars
run_tests

$ROOT/tunnel/scripts/containers.sh destroy
