package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mikkeloscar/go-wlc"
	"github.com/mikkeloscar/go-xkbcommon"
)

// constants you would normally get from linux/input.h
const (
	btnLeft  = 0x110
	btnRight = 0x111
)

type Compositor struct {
	action *action
}

type action struct {
	view  wlc.View
	grab  wlc.Point
	edges uint32
}

func (c *Compositor) startInteractiveAction(view wlc.View, origin wlc.Point) bool {
	if c.action != nil {
		return false
	}

	c.action = &action{
		view: view,
		grab: origin,
	}
	view.BringToFront()
	return true
}

func (c *Compositor) startInteractiveMove(view wlc.View, origin wlc.Point) {
	c.startInteractiveAction(view, origin)
}

func (c *Compositor) startInteractiveResize(view wlc.View, edges uint32, origin wlc.Point) {
	g := view.GetGeometry()
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

	view.SetState(wlc.BitResizing, true)
}

func (c *Compositor) stopInteractiveAction() {
	if c.action == nil {
		return
	}

	c.action.view.SetState(wlc.BitResizing, false)
	c.action = nil
}

func getTopmost(output wlc.Output, offset int) wlc.View {
	views := output.GetViews()
	if offset < len(views) {
		return views[offset]
	}

	return 0
}

func (c *Compositor) relayout(output wlc.Output) {
	r := output.GetResolution()
	if r == nil {
		return
	}

	views := output.GetViews()

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

		view.SetGeometry(0, geometry)

		toggle = !toggle
		if !toggle {
			y += int(h)
		}
	}
}

func (c *Compositor) OutputResolution(output wlc.Output, from *wlc.Size, to *wlc.Size) {
	c.relayout(output)
}

func (c *Compositor) ViewCreated(view wlc.View) bool {
	view.SetMask(view.GetOutput().GetMask())
	view.BringToFront()
	view.Focus()
	c.relayout(view.GetOutput())
	return true
}

func (c *Compositor) ViewDestroyed(view wlc.View) {
	getTopmost(view.GetOutput(), 0).Focus()
	c.relayout(view.GetOutput())
}

func (c *Compositor) ViewFocus(view wlc.View, focus bool) {
	view.SetState(wlc.BitActivated, focus)
}

func (c *Compositor) ViewRequestMove(view wlc.View, origin *wlc.Point) {
	c.startInteractiveMove(view, *origin)
}

func (c *Compositor) ViewRequestResize(view wlc.View, edges uint32, origin *wlc.Point) {
	c.startInteractiveResize(view, edges, *origin)
}

func (c *Compositor) KeyboardKey(view wlc.View, time uint32, modifiers wlc.Modifiers, key uint32, state wlc.KeyState) bool {
	sym := wlc.KeyboardGetKeysymForKey(key, nil)

	if state == wlc.KeyStatePressed {
		if view != 0 {
			if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == xkb.Keyq {
				view.Close()
				return true
			}

			if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == xkb.KeyDown {
				view.SendToBack()
				getTopmost(view.GetOutput(), 0).Focus()
				return true
			}
		}

		if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == xkb.KeyEscape {
			wlc.Terminate()
			return true
		}

		if (modifiers.Mods&wlc.BitModCtrl != 0) && sym == xkb.KeyReturn {
			term := os.Getenv("TERMINAL")
			if len(term) == 0 {
				term = "weston-terminal"
			}
			wlc.Exec(term)
			return true
		}
	}

	return false
}

func (c *Compositor) PointerButton(view wlc.View, time uint32, modifiers wlc.Modifiers, button uint32, state wlc.ButtonState, pos *wlc.Point) bool {
	if state == wlc.ButtonStatePressed {
		view.Focus()
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

func (c *Compositor) PointerMotion(view wlc.View, time uint32, pos *wlc.Point) bool {
	if c.action != nil {
		dx := pos.X - c.action.grab.X
		dy := pos.Y - c.action.grab.Y
		g := c.action.view.GetGeometry()

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

			c.action.view.SetGeometry(c.action.edges, *g)
		} else {
			g.Origin.X += dx
			g.Origin.Y += dy
			c.action.view.SetGeometry(0, *g)
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

	wlc.SetOutputResolutionCb(compositor.OutputResolution)
	wlc.SetViewCreatedCb(compositor.ViewCreated)
	wlc.SetViewDestroyedCb(compositor.ViewDestroyed)
	wlc.SetViewFocusCb(compositor.ViewFocus)
	wlc.SetViewRequestMoveCb(compositor.ViewRequestMove)
	wlc.SetViewRequestResizeCb(compositor.ViewRequestResize)
	wlc.SetKeyboardKeyCb(compositor.KeyboardKey)
	wlc.SetPointerButtonCb(compositor.PointerButton)
	wlc.SetPointerMotionCb(compositor.PointerMotion)

	if !wlc.Init() {
		os.Exit(1)
	}

	wlc.Run()
	os.Exit(0)
}
