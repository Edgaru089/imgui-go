package imgui

// #cgo CPPFLAGS: -DIMGUI_DISABLE_WIN32_DEFAULT_IME_FUNCTIONS
// #cgo CXXFLAGS: -std=c++11
// #cgo CXXFLAGS: -Wno-subobject-linkage
// #cgo prebuilt LDFLAGS:  -L./implot -L./imgui -limplot -limgui
// #cgo prebuilt CPPFLAGS: -DIMPLOT_GO_PREBUILT
import "C"
