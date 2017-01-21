SHELL=/bin/bash -o pipefail

auto_gen := .ci/auto-gen.sh
gen_dirs = query

parse-gen:
	$(auto_gen) $(gen_dirs)

generate: parse-gen

test: generate
	go test ./...