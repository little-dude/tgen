# Makefile inspired by Jon Eisen one for the Go part, but simplified since my
# use case is much simpler
#
# https://gist.github.com/dnishimura/2961173
# http://joneisen.me/post/25503842796


all: go py

BUILD_DATE  = $(shell date +%Y-%m-%d.%H:%M)
GIT_BRANCH  = $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG     = $(shell git describe --abbrev=0 --tags)
GIT_COMMITS = $(shell git rev-list --count $$(git describe --abbrev=0 --tags)..HEAD)
GIT_HASH    = $(shell git rev-list --abbrev-commit -n 1 HEAD)
GIT_DIRTY   = $(shell git diff-index --quiet HEAD -- || echo dirty)
ifeq ($(GIT_DIRTY),dirty)
BUILD_NUMBER := $(GIT_TAG)+$(GIT_COMMITS).g$(GIT_HASH).dirty
else
BUILD_NUMBER := $(GIT_TAG)+$(GIT_COMMITS).g$(GIT_HASH)
endif
version_info:
	@echo "generating build number"
	@echo "building for $(GIT_BRANCH) master"
	@echo "last tag: $(GIT_TAG)"
	@echo "commits since last tag: $(GIT_COMMITS)"
	@echo "last commit: $(GIT_HASH)"
ifeq ($(GIT_DIRTY),dirty)
	@echo "WARNING: the current index is dirty"
endif
	@echo "build number: $(BUILD_NUMBER)"
	@echo "build date: $(BUILD_DATE)"

GO_VERSION_FILE = version.go
GO_CMD 			= go
GO_BUILD 		= $(GO_CMD) build
GO_CLEAN 		= $(GO_CMD) clean
GO_INSTALL 		= $(GO_CMD) install
GO_TEST			= $(GO_CMD) test
GO_DEP			= $(GO_TEST) -i
GO_FMT			= gofmt -w
go: version_info capnp
	@echo "building server"
	cd server ; $(GO_BUILD) -ldflags "-X main.version=$(BUILD_NUMBER) -X main.build=$(BUILD_DATE)" -o ../tgen
	@echo "done building server"

PY_CMD 	   		= python setup.py
PY_BUILD   		= $(PY_CMD) bdist
PY_INSTALL 		= $(PY_CMD) install
PY_VERSION_FILE = client/tgenpy/__version__
py: version_info capnp
	@echo "building python client"
	rm -f $(PY_VERSION_FILE)
	echo $(BUILD_NUMBER) > $(PY_VERSION_FILE)
	cd client && $(PY_BUILD)
	@echo "done building python client"

capnp:
	cp -f schemas.capnp client/tgenpy/schemas.capnp
	rm -f tmp.capnp
	echo "using Go = import \"/zombiezen.com/go/capnproto2/go.capnp\";" > tmp.capnp
	echo "\$$Go.package(\"schemas\");" >> tmp.capnp
	echo "\$$Go.import(\"github.com/little-dude/tgen/server/schemas\");" >> tmp.capnp
	cat schemas.capnp >> tmp.capnp
	capnp compile -ogo tmp.capnp -I $(GOPATH)/src/
	mkdir -p server/schemas/
	mv tmp.capnp.go server/schemas/main.go
	rm -f tmp.capnp

clean:
	rm -f tmp.capnp*
	rm -f schemas/*
	rm -f tgen
	rm -fr client/{build,dist,tgen.egg-info}
