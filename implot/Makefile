CPPFLAGS = -DIMGUI_DISABLE_WIN32_DEFAULT_IME_FUNCTIONS
CFLAGS   = -g -O2
CXXFLAGS = -g -O2 -Wno-unused-result -std=c++11
LDFLAGS  = -s -Wl,-O2
LDLIBS   = -lm -lstdc++

objects_implot = implot.o implot_demo.o implot_items.o

all: libimplot.a


libimplot.a: $(objects_implot)
	$(AR) cr $@ $(objects_implot)

clean:
	rm -f *.o *.a

.PHONY: all clean

