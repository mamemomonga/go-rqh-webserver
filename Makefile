NAME   := rqh-webserver
APPPATH := github.com/mamemomonga/go-rqh-webserver

SRCS     := $(shell find src -type f -name '*.go')
ASSETS   := $(shell find assets -type f)
VERSION  := v$(shell cat version)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
BUILDER_DOCKER_IMAGE := $(NAME)-builder
export GOBIN := $(shell if [ -z "$$GOBIN" ]; then echo "$$GOPATH/bin"; else echo "$$GOBIN"; fi)

# -----------------------

.PHONY: build deps clean run release

build: bin/$(NAME) $(ASSETS_PUBLIC) $(ASSETS_TEMPLATES)

deps: $(GOBIN)/dep $(GOBIN)/packr
	$(GOBIN)/dep ensure -v

clean:
	cd src; packr clean
	rm -rf bin/$(NAME) vendor release

run:
	cd src; packr clean
	cd src/$(NAME); go run .

# -----------------------

release:
	mkdir -p release
	docker build --build-arg APPPATH=$(APPPATH) -t $(BUILDER_DOCKER_IMAGE) .
	docker run --rm $(BUILDER_DOCKER_IMAGE) tar cC /go/src/$(APPPATH)/release . | tar xC release

dcr-release-build: $(PACKR_FILES)
	cd src; packr -v -z
	mkdir -p release
	GOOS=linux   GOARCH=arm   $(MAKE) dcr-release-build-os-arch
	GOOS=linux   GOARCH=amd64 $(MAKE) dcr-release-build-os-arch
	GOOS=darwin  GOARCH=amd64 $(MAKE) dcr-release-build-os-arch
	GOOS=windows GOARCH=amd64 $(MAKE) dcr-release-build-os-arch
	cd release; mv $(NAME)-windows-amd64 $(NAME)-windows-amd64.exe
	chmod 755 release/*

dcr-release-build-os-arch:
	cd src/$(NAME); \
		CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o ../../release/$(NAME)-$(GOOS)-$(GOARCH)

# -----------------------

bin/$(NAME): $(SRCS)
	cd src; packr -v -z
	mkdir -p bin
	cd src/$(NAME); \
		go build $(LDFLAGS) -o ../../bin/$(NAME) 

# -----------------------

$(GOBIN)/dep:
	echo $(GOBIN)
	@if [ "$(go env GOARCH)" == "arm" ]; then \
		go get -v -u github.com/golang/dep/cmd/dep ;\
	else \
		curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh ;\
	fi

$(GOBIN)/packr:
	go get -v -u github.com/gobuffalo/packr/packr


