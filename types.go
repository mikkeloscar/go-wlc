package wlc

/*
#cgo LDFLAGS: -lwlc
#include <wlc/wlc.h>
*/
import "C"

type Handle C.wlc_handle

type EventSource *C.struct_wlc_event_source

// type XKBState *C.struct_xkb_state
// type XKBKeymap *C.struct_xkb_keymap

type InputDevice *C.struct_libinput_device

type LogType C.enum_wlc_log_type

const (
	LogInfo LogType = iota
	LogWarn
	LogError
	LogWayland
)

type BackendType C.enum_wlc_backend_type

const (
	BackendNone BackendType = iota
	BackendDrm
	BackendX11
)

type EventBit C.enum_wlc_event_bit

const (
	EventReadable  EventBit = 0x01
	EventWriteable          = 0x02
	EventHangup             = 0x04
	EventError              = 0x08
)

type ViewStateBit C.enum_wlc_view_state_bit

const (
	BitMaximized  ViewStateBit = 1 << 0
	BitFullscreen              = 1 << 1
	BitResizing                = 1 << 2
	BitMoving                  = 1 << 3
	BitActivated               = 1 << 4
)

type ViewTypeBit C.enum_wlc_view_type_bit

const (
	BitOverrideRedirect ViewTypeBit = 1 << 0
	BitUnmanaged                    = 1 << 1
	BitSplash                       = 1 << 2
	BitModal                        = 1 << 3
	BitPopup                        = 1 << 4
)

type ResizeEdge C.enum_wlc_resize_edge

const (
	ResizeEdgeNone        ResizeEdge = 0
	ResizeEdgeTop                    = 1
	ResizeEdgeBottom                 = 2
	ResizeEdgeLeft                   = 4
	ResizeEdgeTopLeft                = 5
	ResizeEdgeBottomLeft             = 6
	ResizeEdgeRight                  = 8
	ResizeEdgeTopRight               = 9
	ResizeEdgeBottomRight            = 10
)

type ModifierBit C.enum_wlc_modifier_bit

const (
	BitModShift ModifierBit = 1 << 0
	BitModCaps              = 1 << 1
	BitModCtrl              = 1 << 2
	BitModAlt               = 1 << 3
	BitModMod2              = 1 << 4
	BitModMod3              = 1 << 5
	BitModLogo              = 1 << 6
	BitModMod5              = 1 << 7
)

type LedBit C.enum_wlc_led_bit

const (
	BitLedNum    LedBit = 1 << 0
	BitLedCaps          = 1 << 1
	BitLedScroll        = 1 << 2
)

type KeyState C.enum_wlc_key_state

const (
	KeyStateReleased KeyState = 0
	KeyStatePressed           = 1
)

type ButtonState C.enum_wlc_button_state

const (
	ButtonStateReleased = 0
	ButtonStatePressed  = 1
)

type ScrollAxisBit C.enum_wlc_scroll_axis_bit

const (
	ScrollAxisVertical   ScrollAxisBit = 1 << 0
	ScrollAxisHorizontal               = 1 << 1
)

type TouchType C.enum_wlc_touch_type

const (
	TouchDown TouchType = iota
	TouchUp
	TouchMotion
	TouchFrame
	TouchCancel
)

type Modifiers struct {
	Leds uint32
	Mods uint32
}

func (m *Modifiers) c() *C.struct_wlc_modifiers {
	return &C.struct_wlc_modifiers{
		leds: C.uint32_t(m.Leds),
		mods: C.uint32_t(m.Mods),
	}
}

func modsCtoGo(c *C.struct_wlc_modifiers) Modifiers {
	return Modifiers{
		Leds: uint32((*c).leds),
		Mods: uint32((*c).mods),
	}
}
