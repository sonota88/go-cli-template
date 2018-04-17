#!/bin/bash

_fmt(){
  gofmt -w app.go
  gofmt -w src/main/model.go
}

case "$1" in
  fmt)
    _fmt
    ;;
  *)
    ;;
esac
