#!/bin/bash

set -e

. "$(dirname $0)/variables.sh"

find . -name "*_gen.go" -exec rm {} \;

for p in $@; do
  echo "Generating: $PACKAGE/$p"
  go generate $PACKAGE/$p
done
#