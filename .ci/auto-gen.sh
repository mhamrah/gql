#!/bin/bash

set -e

. "$(dirname $0)/variables.sh"

echo $PACKAGE/$1

go generate $PACKAGE/$1

