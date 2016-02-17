package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mikkeloscar/go-wlc"
)

// constants you would normally get from xkbcommon and linux/input.h
const (
	XKBKeyq      = 0x0071
	XKBKeyDown   = 0xff54
	XKBKeyEscape = 0xff1b
	XKBKeyReturn = 0x0ff0d
	btnLeft      = 0x110
	btnRight     = 0x111
)

type Compositor struct {
	action *action
}

type action struct {
	view  wlc.Handle
	grab  wlc.Point
	edges uint32
}

func (c *Compositor) startInteractiveAction(view wlc.Handle, origin wlc.Point) bool {
	if c.action != nil {
		return false
	}

	c.action = &action{
		view: view,
		grab: origin,
	}
	wlc.ViewBringToFront(view)
	return true
}

func (c *Compositor) startInteractiveMove(view wlc.Handle, origin wlc.Point) {
	c.startInteractiveAction(view, origin)
}

func (c *Compositor) startInteractiveResize(view wlc.Handle, edges uint32, origin wlc.Point) {
	g := wlc.ViewGetGeometry(view)
	if g == nil {
		return
	}

	if !c.startInteractiveAction(view, origin) {
		return
	}

	halfw := g.Origin.X + int32(g.Size.W)/2
	halfh := g.Origin.Y + int32(g.Size.H)/2

	if edges == 0 {
		if origin.X < halfw {
			edges = wlc.ResizeEdgeLeft
		} else {
			if origin.X > halfw {
				edges = wlc.ResizeEdgeRight
			}
		}

		if origin.Y < halfh {
			edges |= wlc.ResizeEdgeTop
		} else {
			if origin.Y > halfh {
				edges |= wlc.ResizeEdgeBottom
			}
		}
	}

	c.action.edges = edges

	wlc.ViewSetState(view, wlc.BitResizing, true)
}

func (c *Compositor) stopInteractiveAction() {
	if c.action == nil {
		return
	}

	wlc.ViewSetState(c.action.view, wlc.BitResizing, false)
	c.action = nil
}

func getTopmost(output wlc.Handle, offset int) wlc.Handle {
	views := wlc.OutputGetViews(output)
	if offset < len(views) {
		return views[offset]
	}

	return 0
}

func (c *Compositor) relayout(output wlc.Handle) {
	r := wlc.OutputGetResolution(output)
	if r == nil {
		return
	}

	views := wlc.OutputGetViews(output)

	toggle := false
	y := 0
	w := r.W / 2
	h := r.H / uint32(math.Max(float64((1+len(views))/2), 1))
	for i, view := range views {
		geometry := wlc.Geometry{
			Origin: wlc.Point{
				X: 0,
				Y: int32(y),
			},
			Size: wlc.Size{
				W: w,
				H: h,
			},
		}

		if toggle {
			geometry.Origin.X = int32(w)
		}

		if !toggle && i == len(views)-1 {
			geometry.Size.W = r.W
		}

		wlc.ViewSetGeometry(view, 0, geometry)

		toggle = !toggle
		if !toggle {
			y += int(h)
		}
	}
}

func (c *Compositor) OutputResolution(output wlc.Handle, from *wlc.Size, to *wlc.Size) {
	c.relayout(output)
}

func (c *Compositor) ViewCreated(view wlc.Handle) bool {
	wlc.ViewSetMask(view, wlc.OutputGetMask(wlc.ViewGetOutput(view)))
	wlc.ViewBringToFront(view)
	wlc.ViewFocus(view)
	c.relayout(wlc.ViewGetOutput(view))
	return true
}

func (c *Compositor) ViewDestroyed(view wlc.Handle) {
	wlc.ViewFocus(getTopmost(wlc.ViewGetOutput(view), 0))
	c.relayout(wlc.ViewGetOutput(view))
}

func (c *Compositor) ViewFocus(view wlc.Handle, focus bool) {
	wlc.ViewSetState(view, wlc.BitActivated, focus)
}

func (c *Compositor) ViewRequestMove(view wlc.Handle, origin *wlc.Point) {
	c.startInteractiveMove(view, *origin)
}

func (c *Compositor) ViewRequestResize(view wlc.Handle, edges uint32, origin *wlc.Point) {
	c.startInteractiveResize(view, edges, *origin)
}

func (c *Compositor) KeyboardKey(view wlc.Handle, time uint32, modifiers wlc.Modifiers, key uint32, state wlc.KeyState) bool {
	sym := wlc.KeyboardGetKeysymForKey(key, nil)

	if state == wlc.KeyStatePressed {
		if view != 0 {
			if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == XKBKeyq {
				wlc.ViewClose(view)
				return true
			}

			if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == XKBKeyDown {
				wlc.ViewSendToBack(view)
				wlc.ViewFocus(getTopmost(wlc.ViewGetOutput(view), 0))
				return true
			}
		}

		if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == XKBKeyEscape {
			wlc.Terminate()
			return true
		}

		if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == XKBKeyReturn {
			term := os.Getenv("TERMINAL")
			if len(term) == 0 {
				term = "weston-terminal"
			}
			wlc.Exec(term, nil)
			return true
		}
	}

	return false
}

func (c *Compositor) PointerButton(view wlc.Handle, time uint32, modifiers wlc.Modifiers, button uint32, state wlc.ButtonState, pos *wlc.Point) bool {
	if state == wlc.ButtonStatePressed {
		wlc.ViewFocus(view)
		if view != 0 {
			if (modifiers.Mods&wlc.BitModCtrl != 0) && button == btnLeft {
				c.startInteractiveMove(view, *pos)
			}

			if (modifiers.Mods&wlc.BitModCtrl != 0) && button == btnRight {
				c.startInteractiveResize(view, 0, *pos)
			}
		}
	} else {
		c.stopInteractiveAction()
	}

	if c.action != nil {
		return true
	}

	return false
}

func (c *Compositor) PointerMotion(view wlc.Handle, time uint32, pos *wlc.Point) bool {
	if c.action != nil {
		dx := pos.X - c.action.grab.X
		dy := pos.Y - c.action.grab.Y
		g := wlc.ViewGetGeometry(c.action.view)

		if c.action.edges != 0 {
			min := wlc.Size{80, 40}
			n := *g
			if c.action.edges&wlc.ResizeEdgeLeft != 0 {
				n.Size.W -= uint32(dx)
				n.Origin.X += dx
			} else if c.action.edges&wlc.ResizeEdgeRight != 0 {
				n.Size.W += uint32(dx)
			}

			if c.action.edges&wlc.ResizeEdgeTop != 0 {
				n.Size.H -= uint32(dy)
				n.Origin.Y += dy
			} else if c.action.edges&wlc.ResizeEdgeBottom != 0 {
				n.Size.H += uint32(dy)
			}

			if n.Size.W >= min.W {
				g.Origin.X = n.Origin.X
				g.Size.W = n.Size.W
			}

			if n.Size.H >= min.H {
				g.Origin.Y = n.Origin.Y
				g.Size.H = n.Size.H
			}

			wlc.ViewSetGeometry(c.action.view, c.action.edges, *g)
		} else {
			g.Origin.X += dx
			g.Origin.Y += dy
			wlc.ViewSetGeometry(c.action.view, 0, *g)
		}

		c.action.grab = *pos
	}

	wlc.PointerSetPosition(*pos)
	if c.action != nil {
		return true
	}

	return false
}

func cbLog(typ wlc.LogType, str string) {
	fmt.Printf("%d: %s\n", typ, str)
}

func main() {
	wlc.LogSetHandler(cbLog)

	compositor := Compositor{}

	interf := wlc.Interface{}
	interf.Output.Resolution = compositor.OutputResolution
	interf.View.Created = compositor.ViewCreated
	interf.View.Destroyed = compositor.ViewDestroyed
	interf.View.Focus = compositor.ViewFocus
	interf.View.Request.Move = compositor.ViewRequestMove
	interf.View.Request.Resize = compositor.ViewRequestResize
	interf.Keyboard.Key = compositor.KeyboardKey
	interf.Pointer.Button = compositor.PointerButton
	interf.Pointer.Motion = compositor.PointerMotion

	if !wlc.Init(&interf) {
		os.Exit(1)
	}

	wlc.Run()
	os.Exit(0)
}
