include $(GOROOT)/src/Make.$(GOARCH)

TARG=libusb

CGOFILES=libusb.go

CGO_LDFLAGS=-llibusb

include $(GOROOT)/src/Make.pkg

