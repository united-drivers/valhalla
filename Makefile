GOPKG :=		github.com/united-drivers/valhalla/pb
GOPATH ?=		$(HOME)/go
SOURCES ?=		$(shell find . -name "*.proto")
GOTEMPLATE_INPUT_DIR ?=	.
GOTEMPLATE_OUTPUT_DIR ?= .
DOCKER_OPTS ?=		--rm
DOCKER_ENV_OPTS ?=	-e "GOTEMPLATE_INPUT_DIR=$(GOTEMPLATE_INPUT_DIR)" -e "GOTEMPLATE_OUTPUT_DIR=$(GOTEMPLATE_OUTPUT_DIR)" -e "SOURCES=$(SOURCES)"

all: getproto generate

getproto:
	git clone --depth=1 --recursive https://github.com/valhalla/valhalla
	rsync -avzL ./valhalla/proto/ pb
	rm -rf valhalla

.PHONY: deps
deps:

$(SOURCES): deps
	$(eval CURRENT := $(shell echo $@ | sed 's/.proto//' | sed 's/pb\///'))
ifdef sed
	sed -i '1isyntax = "proto2";\noption go_package = "$(GOPKG)/$(CURRENT)";\n' $@
else
	mkdir -p pb
	mkdir -p $(shell echo $@ | sed 's/.proto//')
	protoc -Ipb --gogo_out=plugins=grpc:$(GOPATH)/src --govalidators_out=plugins=grpc:$(GOPATH)/src --grpc-gateway_out=logtostderr=true:$(GOPATH)/src "$@"

	@# fix grpc-gateway path, should be fixed on upstream repo
	@if [ -f $(OUTDIR)/go-grpc/$(GOPKG)/$(GOPKG)/$(dir $*)/*.pb.gw.go ]; then mv $(OUTDIR)/go-grpc/$(GOPKG)/$(GOPKG)/$(dir $*)/*.pb.gw.go $(OUTDIR)/go-grpc/$(GOPKG)/$(dir $*)/; fi
endif

.PHONY: build
build: $(SOURCES)

generate:
	docker run $(DOCKER_OPTS) $(DOCKER_ENV_OPTS) -e "GIT_TAG=$(GIT_TAG)" -v "$(PWD):$(PWD)" -w "$(PWD)" --entrypoint=/bin/sh uniteddrivers/protoc -xec "make build GOPATH=$(GOPATH) sed=1 && make GOPATH=$(GOPATH) build"
