package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>
*/
import "C"

import "unsafe"

// GetOutputs gets a list of outputs.
func GetOutputs() []Handle {
	var len C.size_t
	handles := C.wlc_get_outputs(&len)
	return handlesCArraytoGoSlice(handles, int(len))
}

// GetFocusedOutput gets focused output.
func GetFocusedOutput() Handle {
	return Handle(C.wlc_get_focused_output())
}

// OutputGetName gets output name.
func OutputGetName(output Handle) string {
	cname := C.wlc_output_get_name(C.wlc_handle(output))
	return C.GoString(cname)
}

// OutputGetSleep gets output sleep state.
func OutputGetSleep(output Handle) bool {
	return bool(C.wlc_output_get_sleep(C.wlc_handle(output)))
}

// OutputSetSleep sets sleep status: wake up / sleep.
func OutputSetSleep(output Handle, sleep bool) {
	C.wlc_output_set_sleep(C.wlc_handle(output), C._Bool(sleep))
}

// OutputGetResolution gets output resolution.
func OutputGetResolution(output Handle) *Size {
	csize := C.wlc_output_get_resolution(C.wlc_handle(output))
	return sizeCtoGo(csize)
}

// OutputSetResolution sets output resolution.
func OutputSetResolution(output Handle, resolution Size) {
	csize := resolution.c()
	defer C.free(unsafe.Pointer(csize))
	C.wlc_output_set_resolution(C.wlc_handle(output), csize)
}

// OutputGetMask gets current visibility bitmask.
func OutputGetMask(output Handle) uint32 {
	return uint32(C.wlc_output_get_mask(C.wlc_handle(output)))
}

// OuputSetMask sets visibility bitmask.
func OutputSetMask(output Handle, mask uint32) {
	C.wlc_output_set_mask(C.wlc_handle(output), C.uint32_t(mask))
}

// TODO: output_get_pixels

// OutputGetViews gets views in stack order.
func OutputGetViews(output Handle) []Handle {
	var len C.size_t
	handles := C.wlc_output_get_views(C.wlc_handle(output), &len)
	return handlesCArraytoGoSlice(handles, int(len))
}

// OutputGetMutableViews gets mutable views in creation order.
//This is mainly useful for wm's who need another view stack for inplace
//sorting. For example tiling wms, may want to use this to keep their tiling
//order separated from floating order.
func OutputGetMutableViews(output Handle) []Handle {
	var len C.size_t
	handles := C.wlc_output_get_mutable_views(C.wlc_handle(output), &len)
	return handlesCArraytoGoSlice(handles, int(len))
}

// OutputSetViews sets views in stack order. This will also change mutable
// views. Returns false on failure.
func OutputSetViews(output Handle, views []Handle) bool {
	// TODO: check that this works (passing views)
	return bool(C.wlc_output_set_views(C.wlc_handle(output), (*C.wlc_handle)(&views[0]), C.size_t(len(views))))
}

// OutputFocus focuses output. Pass zero for no focus.
func OutputFocus(output Handle) {
	C.wlc_output_focus(C.wlc_handle(output))
}
