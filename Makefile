GO = go
GOFLAGS = -ldflags="-s -w"
UPX = upx
UPXFLAGS = --quiet --best --lzma

# The default target:

all: github-oracle

.PHONY: all

# Output executables:

github-oracle: main.go GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(GOFLAGS) -o $@ $^ && $(UPX) $(UPXFLAGS) $@

# Rules for development:

clean: @rm -Rf github-oracle *~

distclean: clean

mostlyclean: clean

maintainer-clean: clean

.PHONY: clean distclean mostlyclean maintainer-clean

.SECONDARY:
.SUFFIXES:
