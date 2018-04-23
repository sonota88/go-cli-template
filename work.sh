#!/bin/bash

_fmt(){
  gofmt -w app.go
  gofmt -w src/main/model.go
}

_test(){
  (
    cd test/main
    go test
  )
}

case "$1" in
  fmt)
    _fmt
    ;;
  "test")
    _test
    ;;
  *)
    ;;
esac
