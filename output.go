package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>
*/
import "C"

import "unsafe"

// Output is a wlc_handle describing an output object in wlc.
type Output C.wlc_handle

// GetOutputs gets a list of outputs.
func GetOutputs() []Output {
	var len C.size_t
	handles := C.wlc_get_outputs(&len)
	return outputHandlesCArraytoGoSlice(handles, int(len))
}

// GetFocusedOutput gets focused output.
func GetFocusedOutput() Output {
	return Output(C.wlc_get_focused_output())
}

// GetName gets output name.
func (o Output) GetName() string {
	cname := C.wlc_output_get_name(C.wlc_handle(o))
	return C.GoString(cname)
}

// GetSleep gets output sleep state.
func (o Output) GetSleep() bool {
	return bool(C.wlc_output_get_sleep(C.wlc_handle(o)))
}

// SetSleep sets sleep status: wake up / sleep.
func (o Output) SetSleep(sleep bool) {
	C.wlc_output_set_sleep(C.wlc_handle(o), C._Bool(sleep))
}

// GetResolution gets output resolution.
func (o Output) GetResolution() *Size {
	csize := C.wlc_output_get_resolution(C.wlc_handle(o))
	return sizeCtoGo(csize)
}

// SetResolution sets output resolution.
func (o Output) SetResolution(resolution Size) {
	csize := resolution.c()
	defer C.free(unsafe.Pointer(csize))
	C.wlc_output_set_resolution(C.wlc_handle(o), csize)
}

// GetMask gets current visibility bitmask.
func (o Output) GetMask() uint32 {
	return uint32(C.wlc_output_get_mask(C.wlc_handle(o)))
}

// SetMask sets visibility bitmask.
func (o Output) SetMask(mask uint32) {
	C.wlc_output_set_mask(C.wlc_handle(o), C.uint32_t(mask))
}

// GetViews gets views in stack order.
func (o Output) GetViews() []View {
	var len C.size_t
	handles := C.wlc_output_get_views(C.wlc_handle(o), &len)
	return viewHandlesCArraytoGoSlice(handles, int(len))
}

// GetMutableViews gets mutable views in creation order.
//This is mainly useful for wm's who need another view stack for inplace
//sorting. For example tiling wms, may want to use this to keep their tiling
//order separated from floating order.
func (o Output) GetMutableViews() []View {
	var len C.size_t
	handles := C.wlc_output_get_mutable_views(C.wlc_handle(o), &len)
	return viewHandlesCArraytoGoSlice(handles, int(len))
}

// SetViews sets views in stack order. This will also change mutable
// views. Returns false on failure.
func (o Output) SetViews(views []View) bool {
	cviews, len := viewHandlesSliceToCArray(views)
	return bool(C.wlc_output_set_views(C.wlc_handle(o), cviews, len))
}

// Focus focuses output. Pass zero for no focus.
func (o Output) Focus() {
	C.wlc_output_focus(C.wlc_handle(o))
}
