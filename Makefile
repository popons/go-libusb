include $(GOROOT)/src/Make.$(GOARCH)

TARG=main

CGOFILES=main.go

CGO_LDFLAGS=-llibusb

include $(GOROOT)/src/Make.pkg

