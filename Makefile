# From
# https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4

# Borrowed from:
# https://github.com/silven/go-example/blob/master/Makefile
# https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BINARY = myserver
VET_REPORT = vet.report
TEST_REPORT = tests.xml

GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOOS=$(shell go env GOOS)

VERSION?=?
COMMIT=$(shell git describe --dirty --abbrev=8 --tags --always --exclude '*')
DESCRIBE=$(shell git describe --dirty --abbrev=8 --tags --always)

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
ifeq ($(VERSION),?)
  LDFLAGS := -X myserver/config.Version=${DESCRIBE}
else
  LDFLAGS := -X myserver/config.Version=${VERSION}-${COMMIT}
endif


GCFLAGS = all=-N -l

# TODO. Build the project
all: myserver
linux: myserver-linux
windows: myserver-windows

link:
	BUILD_DIR=${BUILD_DIR}; \
	BUILD_DIR_LINK=${BUILD_DIR_LINK}; \
	CURRENT_DIR=${CURRENT_DIR}; \
	if [ "$${BUILD_DIR_LINK}" != "$${CURRENT_DIR}" ]; then \
	    echo "Fixing symlinks for build"; \
	    rm -f $${BUILD_DIR}; \
	    ln -s $${CURRENT_DIR} $${BUILD_DIR}; \
	fi

myserver:
	cd ${BUILD_DIR}; \
	mkdir -p bin; \
	GOOS=${GOOS} go build -ldflags "${LDFLAGS}" -gcflags "${GCFLAGS}" -o bin/${BINARY} . ; \
	cd - >/dev/null

myserver-linux:
	cd ${BUILD_DIR}; \
	GOOS=linux go build -ldflags "${LDFLAGS}" -gcflags "${GCFLAGS}" -o bin/${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null

myserver-windows:
	cd ${BUILD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build -ldflags "${LDFLAGS}" -gcflags "${GCFLAGS}" -o bin/${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

test:
	if ! hash go2xunit 2>/dev/null; then go install github.com/tebeka/go2xunit; fi
	cd ${BUILD_DIR}; \
	godep go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT} ; \
	cd - >/dev/null

vet:
	-cd ${BUILD_DIR}; \
	godep go vet ./... > ${VET_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VET_REPORT}
	-rm -rf bin

.PHONY: link myserver linux windows test vet fmt clean