package wlc

/*
#cgo LDFLAGS: -lwlc
#include <wlc/wlc.h>

struct wlc_point *init_point(int32_t x, int32_t y) {
	struct wlc_point *point = malloc(sizeof(struct wlc_point));
	point->x = x;
	point->y = y;
	return point;
}

struct wlc_size *init_size(uint32_t w, uint32_t h) {
	struct wlc_size *size = malloc(sizeof(struct wlc_size));
	size->w = w;
	size->h = h;
	return size;
}

struct wlc_geometry *init_geometry(int32_t x, int32_t y, uint32_t w, uint32_t h) {
	struct wlc_geometry *geometry = malloc(sizeof(struct wlc_geometry));
	geometry->origin.x = x;
	geometry->origin.y = y;
	geometry->size.w = w;
	geometry->size.h = h;
	return geometry;
}
*/
import "C"

type Point struct {
	X, Y int32
}

func (p *Point) c() *C.struct_wlc_point {
	return C.init_point(C.int32_t(p.X), C.int32_t(p.Y))
}

func pointCtoGo(c *C.struct_wlc_point) *Point {
	if c != nil {
		return &Point{
			X: int32((*c).x),
			Y: int32((*c).y),
		}
	}

	return nil
}

type Size struct {
	W, H uint32
}

func (s *Size) c() *C.struct_wlc_size {
	return C.init_size(C.uint32_t(s.W), C.uint32_t(s.H))
}

func sizeCtoGo(c *C.struct_wlc_size) *Size {
	if c != nil {
		return &Size{
			W: uint32((*c).w),
			H: uint32((*c).h),
		}
	}

	return nil
}

type Geometry struct {
	Origin Point
	Size   Size
}

func (g *Geometry) c() *C.struct_wlc_geometry {
	return C.init_geometry(
		C.int32_t(g.Origin.X),
		C.int32_t(g.Origin.Y),
		C.uint32_t(g.Size.W),
		C.uint32_t(g.Size.H),
	)
}

func geometryCtoGo(c *C.struct_wlc_geometry) *Geometry {
	if c != nil {
		return &Geometry{
			Origin: *pointCtoGo((*C.struct_wlc_point)(&(*c).origin)),
			Size:   *sizeCtoGo((*C.struct_wlc_size)(&(*c).size)),
		}
	}

	return nil
}
