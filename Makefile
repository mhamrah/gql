SHELL=/bin/bash -o pipefail

auto_gen := .ci/auto-gen.sh
gen_dirs = parser

.PHONY: dependencies
dependencies:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install
	@echo "Installing test dependencies..."
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/goyacc
	go get github.com/go-playground/overalls

.PHONY: generate
generate:
	$(auto_gen) $(gen_dirs)

.PHONY: test
test: generate
	go test -v ./...

.PHONY: coveralls
coveralls:
 	overalls -project=github.com/mhamrah/gql -covermode=count
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci