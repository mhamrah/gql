SHELL=/bin/bash -o pipefail

auto_gen := .ci/auto-gen.sh
parser_dir := parser

parse-gen:
	$(auto_gen) $(parser_dir)

generate: parse-gen
