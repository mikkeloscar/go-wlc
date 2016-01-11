package wlc

/*
#cgo LDFLAGS: -lwlc
#include <wlc/wlc.h>
*/
import "C"

type Point struct {
	X, Y int32
}

func (p *Point) c() *C.struct_wlc_point {
	return &C.struct_wlc_point{
		x: C.int32_t(p.X),
		y: C.int32_t(p.Y),
	}
}

func pointCtoGo(c *C.struct_wlc_point) Point {
	return Point{
		X: int32((*c).x),
		Y: int32((*c).y),
	}
}

type Size struct {
	W, H uint32
}

func (s *Size) c() *C.struct_wlc_size {
	return &C.struct_wlc_size{
		w: C.uint32_t(s.W),
		h: C.uint32_t(s.H),
	}
}

func sizeCtoGo(c *C.struct_wlc_size) Size {
	return Size{
		W: uint32((*c).w),
		H: uint32((*c).h),
	}
}

type Geometry struct {
	Origin Point
	Size   Size
}

func (g *Geometry) c() *C.struct_wlc_geometry {
	return &C.struct_wlc_geometry{
		origin: *g.Origin.c(),
		size:   *g.Size.c(),
	}
}

func geometryCtoGo(c *C.struct_wlc_geometry) Geometry {
	return Geometry{
		Origin: pointCtoGo((*C.struct_wlc_point)(&(*c).origin)),
		Size:   sizeCtoGo((*C.struct_wlc_size)(&(*c).size)),
	}
}
