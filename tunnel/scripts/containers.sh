#!/bin/bash -e

function network_bastion {
  case $1 in
    create)
      docker network create --driver bridge bastion > /dev/null
    ;;

    destroy)
      docker network rm bastion > /dev/null
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
    ;;

    destroy)
      docker stop ssh-server > /dev/null
      docker rm ssh-server > /dev/null
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
    ;;

    destroy)
      docker stop redis-server > /dev/null
      docker rm redis-server > /dev/null
    ;;
  esac
}

case $1 in
  create)
    network_bastion create
    redis_server    create
    ssh_server      create
  ;;

  destroy)
    redis_server    destroy
    ssh_server      destroy
    network_bastion destroy
  ;;

  *)
    echo "Usage: containers.sh {create|destroy}"
  ;;
esac
