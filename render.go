package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc-render.h>
*/
import "C"

import "unsafe"

// PixelFormat describes the pixelformat used when writing/reading pixels.
type PixelFormat C.enum_wlc_pixel_format

const (
	RGBA8888 PixelFormat = iota
)

// PixelsWrite write pixel data with the specific format to outputs
// framebuffer. If geometry is out of bounds, it will be automatically clamped.
// TODO: make more go friendly
func PixelsWrite(format PixelFormat, geometry Geometry, data unsafe.Pointer) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	C.wlc_pixels_write(C.enum_wlc_pixel_format(format), cgeometry, data)
}

// PixelsRead read pixel data from output's framebuffer.
// If the geometry is out of bounds, it will be automatically clamped.
// Potentially clamped geometry will be stored in out_geometry, to indicate
// width / height of the returned data.
// TODO: make more go friendly
func PixelsRead(format PixelFormat, geometry Geometry, outGeometry *Geometry, out_data unsafe.Pointer) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	var cgOut C.struct_wlc_geometry
	C.wlc_pixels_read(C.enum_wlc_pixel_format(format), cgeometry, &cgOut, out_data)
	geometryCtoGo(outGeometry, &cgOut)
}

// SurfaceRender for rendering surfaces inside post / pre render hooks.
func SurfaceRender(surface Resource, geometry Geometry) {
	cgeometry := geometry.c()
	defer C.free(unsafe.Pointer(cgeometry))
	C.wlc_surface_render(C.wlc_resource(surface), cgeometry)
}

// ScheduleRender schedules output for rendering next frame.
// If output was already scheduled this is no-op, if output is currently
// rendering, it will render immediately after.
func (o Output) ScheduleRender() {
	C.wlc_output_schedule_render(C.wlc_handle(o))
}
