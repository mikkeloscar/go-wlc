package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc-wayland.h>
*/
import "C"

import "unsafe"

type Resource C.wlc_resource

// GetWLDisplay returns wayland display.
func GetWLDisplay() *C.struct_wl_display {
	return C.wlc_get_wl_display()
}

// HandleFromWLSurface returns view handle from wl_surface resource.
func HandleFromWLSurface(resource *C.struct_wl_resource) Handle {
	return Handle(C.wlc_handle_from_wl_surface_resource(resource))
}

// HandleFromWLOutputResource returns output handle from wl_output resource.
func HandleFromWLOutputResource(resource *C.struct_wl_resource) Handle {
	return Handle(C.wlc_handle_from_wl_output_resource(resource))
}

// HandleFromWLSurfaceResource returns internal wlc surface from wl_surface
// resource.
func HandleFromWLSurfaceResource(resource *C.struct_wl_resource) Resource {
	return Resource(C.wlc_handle_from_wl_surface_resource(resource))
}

// ViewGetSurface returns internal wlc surface from view Handle.
func ViewGetSurface(view Handle) Resource {
	return Resource(C.wlc_view_get_surface(C.wlc_handle(view)))
}

// SurfaceGetSize gets surface size.
func SurfaceGetSize(surface Resource) *Size {
	csize := C.wlc_surface_get_size(C.wlc_resource(surface))
	return sizeCtoGo(csize)
}

// SurfaceRender for rendering surfaces inside post / pre render hooks.
func SurfaceRender(surface Resource, geometry Geometry) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	C.wlc_surface_render(C.wlc_resource(surface), cgeometry)
}
