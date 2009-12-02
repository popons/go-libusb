include $(GOROOT)/src/Make.$(GOARCH)

TARG=libusb
CGOFILES=libusb.go

CGO_LDFLAGS=-lusb

include $(GOROOT)/src/Make.pkg

test: install
	$(GC) test.go
	$(LD) -o $@ test.$O
cleanall:clean
	rm test
