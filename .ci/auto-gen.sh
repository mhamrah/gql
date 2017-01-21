#!/bin/bash

set -e

. "$(dirname $0)/variables.sh"

echo "Generating: $PACKAGE/$1"

find . -name "*_gen.go" -exec rm {} \;

go generate $PACKAGE/$1

cat $PACKAGE/$1/graphql_lexer_gen.go

