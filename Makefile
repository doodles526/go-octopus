SHELL := /bin/bash
PKG = github.com/doodles526/go-octopus
SUBPKG_NAMES = finger
SUBPKGS = $(addprefix $(PKG)/, $(SUBPKG_NAMES))
PKGS = $(PKG) $(SUBPKGS)

.PHONY: test $(PKGS)

test: $(PKGS)

$(GOLINT):
	go get github.com/golang/lint/golint
	go build github.com/golang/lint/golint

$(PKGS): $(GOLINT)
ifneq ($(NOLINT),1)
	@PATH=$(PATH):$(GOPATH)/bin golint $(GOPATH)/src/$@*/**.go
	@echo ""
endif
	go get -d -t $@
ifeq ($(COVERAGE),1)
	go test -cover -coverprofile=$(GOPATH)/src/$@/c.out $@ -test.v
	go tool cover -html=$(GOPATH)/src/$@/c.out
else
	go test $@ -test.v
endif
