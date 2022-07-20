CPPFLAGS = -DIMGUI_DISABLE_WIN32_DEFAULT_IME_FUNCTIONS -DIMGUI_USER_CONFIG='"../wrapper/ConfigOverride.h"'
CFLAGS   = -g -O2
CXXFLAGS = -g -O2 -Wno-unused-result -std=c++11
LDFLAGS  = -s -Wl,-O2
LDLIBS   = -lm -lstdc++

objects_imgui  = imgui.o imgui_demo.o imgui_draw.o imgui_tables.o imgui_widgets.o

all: libimgui.a


libimgui.a: $(objects_imgui)
	$(AR) cr $@ $(objects_imgui)

clean:
	rm -f *.o *.a

.PHONY: all clean

