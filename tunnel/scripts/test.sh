#!/bin/bash -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $( dirname $DIR ) )

GREEN='\033[0;32m'
NC='\033[0m' # No Colour

function echo_green {
  echo -e "${GREEN}$1${NC}"
}

function network_bastion {
  case $1 in
    create)
      docker network create --driver bridge bastion > /dev/null
      echo_green "Created network bastion"
    ;;

    destroy)
      docker network rm bastion > /dev/null
      echo_green "Destroyed network bastion"
    ;;
  esac
}

function redis_server {
  case $1 in
    create)
      docker run \
        --network=bastion \
        --name redis-server \
        -d -i cflondonservices/redis-server \
         > /dev/null
      echo_green "Started redis-server container"
    ;;

    destroy)
      docker stop ssh-server > /dev/null
      docker rm ssh-server > /dev/null
      echo_green "Destroyed redis-server container"
    ;;
  esac
}

function ssh_server {
  case $1 in
    create)
      docker run \
        --network=bastion \
        --name ssh-server \
        -d -P cflondonservices/ssh-server \
         > /dev/null
      echo_green "Started ssh-server container"
    ;;

    destroy)
      docker stop redis-server > /dev/null
      docker rm redis-server > /dev/null
      echo_green "Destroyed ssh-server container"
    ;;
  esac
}

network_bastion create
redis_server    create
ssh_server      create

export REDIS_HOST=$( docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-server )
export REDIS_PORT="6379"
SSH_ADDRESS=$( docker port ssh-server 22 )
IFS=: read SSH_HOST SSH_PORT <<< "${SSH_ADDRESS}"
export SSH_HOST=${SSH_HOST}
export SSH_PORT=${SSH_PORT}
export SSH_USER="vcap"
export SSH_PASSWORD="funky92horse"

pushd $ROOT/tunnel > /dev/null
  ginkgo -v .
popd > /dev/null

redis_server    destroy
ssh_server      destroy
network_bastion destroy
