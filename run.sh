#!/bin/bash

_get_project_dir() {
  local real_path="$(readlink --canonicalize "$0")"
  (
    cd "$(dirname "$real_path")"
    pwd
  )
}


# --------------------------------
# Main

CURRENT_DIR=$(pwd)
PROJECT_DIR=$(_get_project_dir)

if [ "$1" = "build" ]; then
  go build
else
  go build
  if [ $? -eq 0 ]; then
    "${PROJECT_DIR}/go-cli-template" simple "${CURRENT_DIR}" "${PROJECT_DIR}" "$@"
  fi
fi
