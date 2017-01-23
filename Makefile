SHELL=/bin/bash -o pipefail

auto_gen := .ci/auto-gen.sh
gen_dirs = query

.PHONY: dependencies
dependencies:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install
	@echo "Installing test dependencies..."
	go get github.com/mattn/goveralls

.PHONY: parse-gen
parse-gen:
	$(auto_gen) $(gen_dirs)

.PHONY: generate
generate: parse-gen

.PHONY: test
test:
	go test -v -covermode=count -coverprofile=cover.out ./...

.PHONY: coveralls
coveralls:
	goveralls -coverprofile=cover.out -service=travis-ci