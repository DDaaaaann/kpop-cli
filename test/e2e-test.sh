#!/usr/bin/env bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <matrix-os>"
  exit 1
fi

os=$1

echo "Running end-to-end test on $os"
