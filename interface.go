package wlc

/*
#cgo LDFLAGS: -lwlc
#include <wlc/wlc.h>

// output
extern bool handle_output_created(wlc_handle output);
extern void handle_output_destroyed(wlc_handle output);
extern void handle_output_focus(wlc_handle output, bool focus);
extern void handle_output_resolution(wlc_handle output, const struct wlc_size *from, const struct wlc_size *to);
extern void handle_output_pre_render(wlc_handle output);
extern void handle_output_post_render(wlc_handle output);
// view
extern bool handle_view_created(wlc_handle view);
extern void handle_view_destroyed(wlc_handle view);
extern void handle_view_focus(wlc_handle view, bool focus);
extern void handle_view_move_to_output(wlc_handle view, wlc_handle from_output, wlc_handle to_output);
extern void handle_view_geometry_request(wlc_handle view, const struct wlc_geometry*);
extern void handle_view_state_request(wlc_handle view, enum wlc_view_state_bit, bool toggle);
extern void handle_view_move_request(wlc_handle view, const struct wlc_point*);
extern void handle_view_resize_request(wlc_handle view, uint32_t edges, const struct wlc_point*);
extern void handle_view_pre_render(wlc_handle view);
extern void handle_view_post_render(wlc_handle view);
// keyboard
extern bool handle_keyboard_key(wlc_handle view, uint32_t time, const struct wlc_modifiers*, uint32_t key, enum wlc_key_state);
// pointer
extern bool handle_pointer_button(wlc_handle view, uint32_t time, const struct wlc_modifiers*, uint32_t button, enum wlc_button_state, const struct wlc_point*);
extern bool handle_pointer_scroll(wlc_handle view, uint32_t time, const struct wlc_modifiers*, uint8_t axis_bits, double amount[2]);
extern bool handle_pointer_motion(wlc_handle view, uint32_t time, const struct wlc_point*);
// touch
extern bool handle_touch_touch(wlc_handle view, uint32_t time, const struct wlc_modifiers*, enum wlc_touch_type, int32_t slot, const struct wlc_point*);
// compositor
extern void handle_compositor_ready(void);
// input
extern bool handle_input_created(struct libinput_device *device);
extern void handle_input_destroyed(struct libinput_device *device);
*/
import "C"
import "unsafe"

var wlcInterface *Interface

// Interface is used for commication with wlc.
type Interface struct {
	Output struct {
		Created    func(Handle) bool
		Destroyed  func(Handle)
		Focus      func(Handle, bool)
		Resolution func(Handle, *Size, *Size)
		Render     struct {
			Pre  func(Handle)
			Post func(Handle)
		}
	}
	View struct {
		Created      func(Handle) bool
		Destroyed    func(Handle)
		Focus        func(Handle, bool)
		MoveToOutput func(Handle, Handle, Handle)
		Request      struct {
			Geometry func(Handle, *Geometry)
			State    func(Handle, ViewStateBit, bool)
			Move     func(Handle, *Point)
			Resize   func(Handle, uint32, *Point)
		}
		Render struct {
			Pre  func(Handle)
			Post func(Handle)
		}
	}
	Keyboard struct {
		Key func(Handle, uint32, Modifiers, uint32, KeyState) bool
	}
	Pointer struct {
		Button func(Handle, uint32, Modifiers, uint32, ButtonState, *Point) bool
		Scroll func(Handle, uint32, Modifiers, uint8, [2]float64) bool
		Motion func(Handle, uint32, *Point) bool
	}
	Touch struct {
		Touch func(Handle, uint32, Modifiers, TouchType, int32, *Point) bool
	}
	Compositor struct {
		Ready func()
	}
	Input struct {
		Created   func(*C.struct_libinput_device) bool
		Destroyed func(*C.struct_libinput_device)
	}
}

// output wrappers

//export _go_handle_output_created
func _go_handle_output_created(output C.wlc_handle) C._Bool {
	return C._Bool(wlcInterface.Output.Created(Handle(output)))
}

//export _go_handle_output_destroyed
func _go_handle_output_destroyed(output C.wlc_handle) {
	wlcInterface.Output.Destroyed(Handle(output))
}

//export _go_handle_output_focus
func _go_handle_output_focus(output C.wlc_handle, focus bool) {
	wlcInterface.Output.Focus(Handle(output), focus)
}

//export _go_handle_output_resolution
func _go_handle_output_resolution(output C.wlc_handle, from *C.struct_wlc_size, to *C.struct_wlc_size) {
	wlcInterface.Output.Resolution(Handle(output), sizeCtoGo(from), sizeCtoGo(to))
}

//export _go_handle_output_pre_render
func _go_handle_output_pre_render(output C.wlc_handle) {
	wlcInterface.Output.Render.Pre(Handle(output))
}

//export _go_handle_output_post_render
func _go_handle_output_post_render(output C.wlc_handle) {
	wlcInterface.Output.Render.Post(Handle(output))
}

// view wrappers

//export _go_handle_view_created
func _go_handle_view_created(view C.wlc_handle) C._Bool {
	return C._Bool(wlcInterface.View.Created(Handle(view)))
}

//export _go_handle_view_destroyed
func _go_handle_view_destroyed(view C.wlc_handle) {
	wlcInterface.View.Destroyed(Handle(view))
}

//export _go_handle_view_focus
func _go_handle_view_focus(view C.wlc_handle, focus bool) {
	wlcInterface.View.Focus(Handle(view), focus)
}

//export _go_handle_view_move_to_output
func _go_handle_view_move_to_output(view C.wlc_handle, fromOutput C.wlc_handle, toOutput C.wlc_handle) {
	wlcInterface.View.MoveToOutput(Handle(view), Handle(fromOutput), Handle(toOutput))
}

//export _go_handle_view_geometry_request
func _go_handle_view_geometry_request(view C.wlc_handle, geometry *C.struct_wlc_geometry) {
	wlcInterface.View.Request.Geometry(Handle(view), geometryCtoGo(geometry))
}

//export _go_handle_view_state_request
func _go_handle_view_state_request(view C.wlc_handle, state C.enum_wlc_view_state_bit, toggle bool) {
	wlcInterface.View.Request.State(Handle(view), ViewStateBit(state), toggle)
}

//export _go_handle_view_move_request
func _go_handle_view_move_request(view C.wlc_handle, point *C.struct_wlc_point) {
	wlcInterface.View.Request.Move(Handle(view), pointCtoGo(point))
}

//export _go_handle_view_resize_request
func _go_handle_view_resize_request(view C.wlc_handle, edges C.uint32_t, point *C.struct_wlc_point) {
	wlcInterface.View.Request.Resize(Handle(view), uint32(edges), pointCtoGo(point))
}

//export _go_handle_view_pre_render
func _go_handle_view_pre_render(view C.wlc_handle) {
	wlcInterface.View.Render.Pre(Handle(view))
}

//export _go_handle_view_post_render
func _go_handle_view_post_render(view C.wlc_handle) {
	wlcInterface.View.Render.Post(Handle(view))
}

// keyboard wrapper

//export _go_handle_keyboard_key
func _go_handle_keyboard_key(view C.wlc_handle, time C.uint32_t, modifiers *C.struct_wlc_modifiers, key C.uint32_t, state C.enum_wlc_key_state) C._Bool {
	return C._Bool(wlcInterface.Keyboard.Key(
		Handle(view),
		uint32(time),
		modsCtoGo(modifiers),
		uint32(key),
		KeyState(state),
	))
}

// pointer wrapper

//export _go_handle_pointer_button
func _go_handle_pointer_button(view C.wlc_handle, time C.uint32_t, modifiers *C.struct_wlc_modifiers, button C.uint32_t, state C.enum_wlc_button_state, point *C.struct_wlc_point) C._Bool {
	return C._Bool(wlcInterface.Pointer.Button(
		Handle(view),
		uint32(time),
		modsCtoGo(modifiers),
		uint32(button),
		ButtonState(state),
		pointCtoGo(point),
	))
}

//export _go_handle_pointer_scroll
func _go_handle_pointer_scroll(view C.wlc_handle, time C.uint32_t, modifiers *C.struct_wlc_modifiers, axisBits C.uint8_t, amount *C.double) C._Bool {
	// convert double[2] to [2]float64
	goAmount := [2]float64{
		*(*float64)(amount),
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(amount)) + unsafe.Sizeof(*amount))),
	}
	return C._Bool(wlcInterface.Pointer.Scroll(
		Handle(view),
		uint32(time),
		modsCtoGo(modifiers),
		uint8(axisBits),
		goAmount,
	))
}

//export _go_handle_pointer_motion
func _go_handle_pointer_motion(view C.wlc_handle, time C.uint32_t, point *C.struct_wlc_point) C._Bool {
	return C._Bool(wlcInterface.Pointer.Motion(
		Handle(view),
		uint32(time),
		pointCtoGo(point),
	))
}

// touch wrapper

//export _go_handle_touch_touch
func _go_handle_touch_touch(view C.wlc_handle, time C.uint32_t, modifiers *C.struct_wlc_modifiers, touch C.enum_wlc_touch_type, slot C.int32_t, point *C.struct_wlc_point) C._Bool {
	return C._Bool(wlcInterface.Touch.Touch(
		Handle(view),
		uint32(time),
		modsCtoGo(modifiers),
		TouchType(touch),
		int32(slot),
		pointCtoGo(point),
	))
}

// compositor wrapper

//export _go_handle_compositor_ready
func _go_handle_compositor_ready() {
	wlcInterface.Compositor.Ready()
}

// input wrapper

//export _go_handle_input_created
func _go_handle_input_created(device *C.struct_libinput_device) C._Bool {
	return C._Bool(wlcInterface.Input.Created(device))
}

//export _go_handle_input_destroyed
func _go_handle_input_destroyed(device *C.struct_libinput_device) {
	wlcInterface.Input.Destroyed(device)
}
