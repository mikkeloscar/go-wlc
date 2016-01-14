package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>
*/
import "C"

import "unsafe"

// ViewFocus focuses view. Pass zero for no focus.
func ViewFocus(view Handle) {
	C.wlc_view_focus(C.wlc_handle(view))
}

// ViewClose closes view.
func ViewClose(view Handle) {
	C.wlc_view_close(C.wlc_handle(view))
}

// ViewGetOutput gets output of view.
func ViewGetOutput(view Handle) Handle {
	return Handle(C.wlc_view_get_output(C.wlc_handle(view)))
}

// ViewSetOutput sets output for view. Alternatively OutputSetViews can be used.
func ViewSetOutput(view Handle, output Handle) {
	C.wlc_view_set_output(C.wlc_handle(view), C.wlc_handle(output))
}

// ViewSendToBack sends view behind everything.
func ViewSendToBack(view Handle) {
	C.wlc_view_send_to_back(C.wlc_handle(view))
}

// ViewSendBelow sends view below another view.
func ViewSendBelow(view Handle, other Handle) {
	C.wlc_view_send_below(C.wlc_handle(view), C.wlc_handle(other))
}

// ViewBringAbove brings view above another view.
func ViewBringAbove(view Handle, other Handle) {
	C.wlc_view_bring_above(C.wlc_handle(view), C.wlc_handle(other))
}

// ViewBringToFront brings view to front of everything.
func ViewBringToFront(view Handle) {
	C.wlc_view_bring_to_front(C.wlc_handle(view))
}

// ViewGetMask gets current visibility bitmask.
func ViewGetMask(view Handle) uint32 {
	return uint32(C.wlc_view_get_mask(C.wlc_handle(view)))
}

// ViewSetMask sets visibility bitmask.
func ViewSetMask(view Handle, mask uint32) {
	C.wlc_view_set_mask(C.wlc_handle(view), C.uint32_t(mask))
}

// ViewGetGeometry gets current geometry.
func ViewGetGeometry(view Handle) *Geometry {
	cgeometry := C.wlc_view_get_geometry(C.wlc_handle(view))
	return geometryCtoGo(cgeometry)
}

// ViewSetGeometry sets geometry. Set edges if the geometry change is caused by
// interactive resize.
func ViewSetGeometry(view Handle, edges uint32, geometry Geometry) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	C.wlc_view_set_geometry(C.wlc_handle(view), C.uint32_t(edges), cgeometry)
}

// ViewGetType gets type bitfield for view.
func ViewGetType(view Handle) uint32 {
	return uint32(C.wlc_view_get_type(C.wlc_handle(view)))
}

// ViewSetType sets type bit. TOggle indicates whether it is set or not.
func ViewSetType(view Handle, typ ViewTypeBit, toggle bool) {
	C.wlc_view_set_type(C.wlc_handle(view), uint32(typ), C._Bool(toggle))
}

// ViewGetState gets current state bitfield.
func ViewGetState(view Handle) uint32 {
	return uint32(C.wlc_view_get_state(C.wlc_handle(view)))
}

// ViewSetState sets state bit. Toggle indicates whether it is set or not.
func ViewSetState(view Handle, state ViewStateBit, toggle bool) {
	C.wlc_view_set_state(C.wlc_handle(view), uint32(state), C._Bool(toggle))
}

// ViewGetParent gets parent view.
func ViewGetParent(view Handle) Handle {
	return Handle(C.wlc_view_get_parent(C.wlc_handle(view)))
}

// ViewSetParent sets parent view.
func ViewSetParent(view Handle, parent Handle) {
	C.wlc_view_set_parent(C.wlc_handle(view), C.wlc_handle(parent))
}

// ViewGetTitle gets title.
func ViewGetTitle(view Handle) string {
	ctitle := C.wlc_view_get_title(C.wlc_handle(view))
	return C.GoString(ctitle)
}

// ViewGetClass gets class. (shell-surface only).
func ViewGetClass(view Handle) string {
	cclass := C.wlc_view_get_class(C.wlc_handle(view))
	return C.GoString(cclass)
}

// ViewGetAppId gets app id. (xdg-surface only).
func ViewGetAppId(view Handle) string {
	capp := C.wlc_view_get_app_id(C.wlc_handle(view))
	return C.GoString(capp)
}
