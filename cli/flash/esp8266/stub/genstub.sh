#!/bin/bash

scriptdir=$(realpath $(dirname $0))

set -e -x

make -C "$scriptdir" \
  LIBS="../../esp/stub/miniz.c uart.c" \
  STUB="stub.json"
