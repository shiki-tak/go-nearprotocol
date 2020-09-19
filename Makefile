.PHONY:build build-rust

DLL_EXT = ""
ifeq ($(OS),Windows_NT)
	DLL_EXT = dll
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		DLL_EXT = so
	endif
	ifeq ($(UNAME_S),Darwin)
		DLL_EXT = dylib
	endif
endif

build: build-rust

# don't strip for now, for better error reporting
# build-rust: build-rust-release strip
build-rust: build-rust-release

# use release build to actually ship - smaller and much faster
build-rust-release:
	cargo build --release --features backtraces
	cp target/release/libgo_nearprotocol.$(DLL_EXT) api
	@ #this pulls out ELF symbols, 80% size reduction!
