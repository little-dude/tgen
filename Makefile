# Makefile inspired by Jon Eisen one for the Go part, but simplified since my
# use case is much simpler
#
# https://gist.github.com/dnishimura/2961173
# http://joneisen.me/post/25503842796

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

# Versionning
BUILD_DATE  := `date +%Y-%m-%d\ %H:%M`
GIT_BRANCH  := `git rev-parse --abbrev-ref HEAD`
GIT_TAG     := `git describe --abbrev=0`
GIT_COMMITS := `git rev-list --count $$(git describe --abbrev=0)..HEAD`
GIT_HASH    := `git rev-list --abbrev-commit -n 1 HEAD`
GIT_DIRTY   := `git diff-index --quiet HEAD -- || echo dirty`
ifeq ($(GIT_DIRTY),dirty)
BUILD_NUMBER := $(GIT_TAG)+$(GIT_COMMITS).g$(GIT_HASH).dirty
else
BUILD_NUMBER := $(GIT_TAG)+$(GIT_COMMITS).g$(GIT_HASH)
endif

all: go

version_info:
	@echo "generating build number"
	@echo "building for branch.......$(GIT_BRANCH)"
	@echo "last tag..................$(GIT_TAG)"
	@echo "commits since last tag....$(GIT_COMMITS)"
	@echo "last commit...............$(GIT_HASH)"
ifeq ($(GIT_DIRTY),dirty)
	@echo "WARNING: the current index is dirty"
endif
	@echo "build number..............$(BUILD_NUMBER)"
	@echo "build date................$(BUILD_DATE)"

GO_VERSION_FILE := version.go

# http://stackoverflow.com/a/25003729/1836144
go: version_info
	@echo "building server"
	@echo "creating $(GO_VERSION_FILE)"
	@rm -f $(GO_VERSION_FILE)
	@echo "package main" 					>  $(GO_VERSION_FILE)
	@echo "const (" 						>> $(GO_VERSION_FILE)
	@echo "	VERSION = \"$(BUILD_NUMBER)\"" 	>> $(GO_VERSION_FILE)
	@echo "	BUILD_DATE = \"$(BUILD_DATE)\"" >> $(GO_VERSION_FILE)
	@echo ")" 								>> $(GO_VERSION_FILE)
	@echo "compiling..."
	@go build
	@echo "done building server"
