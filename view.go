package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>
*/
import "C"

import "unsafe"

// View is a wlc_handle describing a view object in wlc.
type View C.wlc_handle

// Focus focuses view. Pass zero for no focus.
func (v View) Focus() {
	C.wlc_view_focus(C.wlc_handle(v))
}

// Close closes view.
func (v View) Close() {
	C.wlc_view_close(C.wlc_handle(v))
}

// GetOutput gets output of view.
func (v View) GetOutput() Output {
	return Output(C.wlc_view_get_output(C.wlc_handle(v)))
}

// SetOutput sets output for view. Alternatively OutputSetViews can be used.
func (v View) SetOutput(output Output) {
	C.wlc_view_set_output(C.wlc_handle(v), C.wlc_handle(output))
}

// SendToBack sends view behind everything.
func (v View) SendToBack() {
	C.wlc_view_send_to_back(C.wlc_handle(v))
}

// SendBelow sends view below another view.
func (v View) SendBelow(other View) {
	C.wlc_view_send_below(C.wlc_handle(v), C.wlc_handle(other))
}

// BringAbove brings view above another view.
func (v View) BringAbove(other View) {
	C.wlc_view_bring_above(C.wlc_handle(v), C.wlc_handle(other))
}

// BringToFront brings view to front of everything.
func (v View) BringToFront() {
	C.wlc_view_bring_to_front(C.wlc_handle(v))
}

// GetMask gets current visibility bitmask.
func (v View) GetMask() uint32 {
	return uint32(C.wlc_view_get_mask(C.wlc_handle(v)))
}

// SetMask sets visibility bitmask.
func (v View) SetMask(mask uint32) {
	C.wlc_view_set_mask(C.wlc_handle(v), C.uint32_t(mask))
}

// GetGeometry gets current geometry.
func (v View) GetGeometry() *Geometry {
	cgeometry := C.wlc_view_get_geometry(C.wlc_handle(v))
	return geometryCtoGo(&Geometry{}, cgeometry)
}

// SetGeometry sets geometry. Set edges if the geometry change is caused by
// interactive resize.
func (v View) SetGeometry(edges uint32, geometry Geometry) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	C.wlc_view_set_geometry(C.wlc_handle(v), C.uint32_t(edges), cgeometry)
}

// GetType gets type bitfield for view.
func (v View) GetType() uint32 {
	return uint32(C.wlc_view_get_type(C.wlc_handle(v)))
}

// SetType sets type bit. TOggle indicates whether it is set or not.
func (v View) SetType(typ ViewTypeBit, toggle bool) {
	C.wlc_view_set_type(C.wlc_handle(v), uint32(typ), C._Bool(toggle))
}

// GetState gets current state bitfield.
func (v View) GetState() uint32 {
	return uint32(C.wlc_view_get_state(C.wlc_handle(v)))
}

// SetState sets state bit. Toggle indicates whether it is set or not.
func (v View) SetState(state ViewStateBit, toggle bool) {
	C.wlc_view_set_state(C.wlc_handle(v), uint32(state), C._Bool(toggle))
}

// GetParent gets parent view.
func (v View) GetParent() View {
	return View(C.wlc_view_get_parent(C.wlc_handle(v)))
}

// SetParent sets parent view.
func (v View) SetParent(parent View) {
	C.wlc_view_set_parent(C.wlc_handle(v), C.wlc_handle(parent))
}

// GetTitle gets title.
func (v View) GetTitle() string {
	ctitle := C.wlc_view_get_title(C.wlc_handle(v))
	return C.GoString(ctitle)
}

// GetClass gets class. (shell-surface only).
func (v View) GetClass() string {
	cclass := C.wlc_view_get_class(C.wlc_handle(v))
	return C.GoString(cclass)
}

// GetAppID gets app id. (xdg-surface only).
func (v View) GetAppID() string {
	capp := C.wlc_view_get_app_id(C.wlc_handle(v))
	return C.GoString(capp)
}
