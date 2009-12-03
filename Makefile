include $(GOROOT)/src/Make.$(GOARCH)

TARG=libusb
CGOFILES=libusb.go

CGO_LDFLAGS=-lusb

include $(GOROOT)/src/Make.pkg

bench: install
	$(GC) bench.go
	$(LD) -o $@ bench.$O
	./bench
cleanall:clean
	rm bench
