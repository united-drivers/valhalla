GOPKG :=		github.com/united-drivers/valhalla/valhalla/pb
SOURCES ?=		$(shell find . -name "*.proto")
GOTEMPLATE_INPUT_DIR ?=	.
GOTEMPLATE_OUTPUT_DIR ?= .
DOCKER_OPTS ?=		--rm
DOCKER_ENV_OPTS ?=	-e "GOTEMPLATE_INPUT_DIR=$(GOTEMPLATE_INPUT_DIR)" -e "GOTEMPLATE_OUTPUT_DIR=$(GOTEMPLATE_OUTPUT_DIR)" -e "SOURCES=$(SOURCES)"

all: getproto generate

getproto:
	git clone --depth=1 --recursive https://github.com/valhalla/valhalla
	rsync -avzL ./valhalla/proto/ proto
	rm -rf valhalla

.PHONY: deps
deps:

$(SOURCES): deps
	mkdir -p pb
	protoc -Iproto --gogo_out=plugins=grpc:./pb --govalidators_out=plugins=grpc:./pb --grpc-gateway_out=logtostderr=true:./pb/ "$@"

.PHONY: build
build: $(SOURCES)

generate:
	docker run $(DOCKER_OPTS) $(DOCKER_ENV_OPTS) -e "GIT_TAG=$(GIT_TAG)" -v "$(PWD):$(PWD)" -w "$(PWD)" --entrypoint=/bin/sh uniteddrivers/protoc -xec "make build"
