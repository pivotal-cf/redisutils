#!/bin/bash

function kill_and_wait() {
  declare pidfile="$1" timeout="${2:-25}" sigkill_on_timeout="${3:-1}"

  if [ ! -f "${pidfile}" ]; then
    echo "Pidfile ${pidfile} doesn't exist"
    exit 0
  fi

  local pid
  pid=$(head -1 "${pidfile}")

  if [ -z "${pid}" ]; then
    echo "Unable to get pid from ${pidfile}"
    exit 1
  fi

  if ! pid_is_running "${pid}"; then
    echo "Process ${pid} is not running"
    rm -f "${pidfile}"
    exit 0
  fi

  echo "Killing ${pidfile}: ${pid} "
  kill "${pid}"

  if ! wait_pid_death "${pid}" "${timeout}"; then
    if [ "${sigkill_on_timeout}" = "1" ]; then
      echo "Kill timed out, using kill -9 on ${pid}"
      kill -9 "${pid}"
      sleep 0.5
    fi
  fi

  if pid_is_running "${pid}"; then
    echo "Timed Out"
    exit 1
  else
    echo "Stopped"
    rm -f "${pidfile}"
  fi
}
