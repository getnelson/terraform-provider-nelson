TEST ?= $$(go list ./... | grep -v 'vendor')
PKG		:= github.com/getnelson/terraform-provider-nelson

export GO111MODULE=on
export TESTARGS=-race

build:
	go install

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -count=1

format: tools.goimports
	@echo "--> formatting code with 'goimports' tool"
	@goimports -local $(PKG) -w -l $(GOFILES)

tools.goimports:
	@command -v goimports >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing goimports"; \
		GO111MODULE=off go get golang.org/x/tools/cmd/goimports; \
	fi

build-binaries:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

.PHONY: build test testacc format tools.goimports
