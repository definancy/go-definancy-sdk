UNAME := $(shell uname)

GOTAGSLIST  := ${GOTAGSCUSTOM}

ifeq ($(UNAME), Linux)
GOTAGSLIST	+= osusergo netgo static_build
GOBUILDMODE := -buildmode pie
EXTLDFLAGS	:= -static-libstdc++ -static-libgcc
# the following predicate is abit misleading; it tests if we're not in centos.
ifeq (,$(wildcard /etc/centos-release))
EXTLDFLAGS  += -static
endif
endif

# If build number already set, use it - to ensure same build number across multiple platforms being built
BUILDNUMBER		?= $(shell echo 1)
COMMITHASH		:= $(shell echo 1)

GOLDFLAGS_BASE	:= -extldflags \"$(EXTLDFLAGS)\"

GOMOD_DIRS := 

GOTRIMPATH	:= $(shell GOPATH=$(GOPATH) && go help build | grep -q .-trimpath && echo -trimpath)
GOTAGS      := --tags "$(GOTAGSLIST)"
GOLDFLAGS 	:= $(GOLDFLAGS_BASE)

default: examples

examples: clean tidy
	go build -o ./build/examples/ $(GOTRIMPATH) $(GOTAGS) $(GOBUILDMODE) -ldflags="$(GOLDFLAGS)" ./...

clean:
	go clean -i ./...
	rm -rf ./build

tidy:
	@echo "Tidying"
	go mod tidy
	@for dir in $(GOMOD_DIRS); do \
		echo "Tidying $$dir" && \
		(cd $$dir && go mod tidy); \
	done
