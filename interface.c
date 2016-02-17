#include "_cgo_export.h"
#include <wlc/wlc.h>

/* output */
bool handle_output_created(wlc_handle output) {
	return _go_handle_output_created(output);
}

void handle_output_destroyed(wlc_handle output) {
	_go_handle_output_destroyed(output);
}

void handle_output_focus(wlc_handle output, bool focus) {
	_go_handle_output_focus(output, focus);
}

void handle_output_resolution(wlc_handle output, const struct wlc_size *from, const struct wlc_size *to) {
	struct wlc_size nc_from = {
		.w = from->w,
		.h = from->h
	};
	struct wlc_size nc_to = {
		.w = to->w,
		.h = to->h
	};
	_go_handle_output_resolution(output, &nc_from, &nc_to);
}

void handle_output_pre_render(wlc_handle output) {
	_go_handle_output_pre_render(output);
}

void handle_output_post_render(wlc_handle output) {
	_go_handle_output_post_render(output);
}

/* view */
bool handle_view_created(wlc_handle view) {
	return _go_handle_view_created(view);
}

void handle_view_destroyed(wlc_handle view) {
	_go_handle_view_destroyed(view);
}

void handle_view_focus(wlc_handle view, bool focus) {
	_go_handle_view_focus(view, focus);
}

void handle_view_move_to_output(wlc_handle view, wlc_handle from_output, wlc_handle to_output) {
	_go_handle_view_move_to_output(view, from_output, to_output);
}

void handle_view_geometry_request(wlc_handle view, const struct wlc_geometry *geometry) {
	struct wlc_geometry nc_geometry = {
		.origin = {
			.x = geometry->origin.x,
			.y = geometry->origin.y
		},
		.size = {
			.w = geometry->size.w,
			.h = geometry->size.h
		}
	};
	_go_handle_view_geometry_request(view, &nc_geometry);
}

void handle_view_state_request(wlc_handle view, enum wlc_view_state_bit bit, bool toggle) {
	_go_handle_view_state_request(view, bit, toggle);
}

void handle_view_move_request(wlc_handle view, const struct wlc_point *point) {
	struct wlc_point nc_point = {
		.x = point->x,
		.y = point->y
	};
	_go_handle_view_move_request(view, &nc_point);
}

void handle_view_resize_request(wlc_handle view, uint32_t edges, const struct wlc_point *point) {
	struct wlc_point nc_point = {
		.x = point->x,
		.y = point->y
	};
	_go_handle_view_resize_request(view, edges, &nc_point);
}

void handle_view_pre_render(wlc_handle view) {
	_go_handle_view_pre_render(view);
}

void handle_view_post_render(wlc_handle view) {
	_go_handle_view_post_render(view);
}

/* keyboard */
bool handle_keyboard_key(wlc_handle view, uint32_t time, const struct wlc_modifiers *modifiers, uint32_t key, enum wlc_key_state state) {
	struct wlc_modifiers nc_modifiers = {
		.leds = modifiers->leds,
		.mods = modifiers->mods
	};
	return _go_handle_keyboard_key(view, time, &nc_modifiers, key, state);
}

/* pointer */
bool handle_pointer_button(wlc_handle view, uint32_t time, const struct wlc_modifiers *modifiers, uint32_t button, enum wlc_button_state state, const struct wlc_point *point) {
	struct wlc_modifiers nc_modifiers = {
		.leds = modifiers->leds,
		.mods = modifiers->mods
	};
	struct wlc_point nc_point = {
		.x = point->x,
		.y = point->y
	};
	return _go_handle_pointer_button(view, time, &nc_modifiers, button, state, &nc_point);
}

bool handle_pointer_scroll(wlc_handle view, uint32_t time, const struct wlc_modifiers *modifiers, uint8_t axis_bits, double amount[2]) {
	struct wlc_modifiers nc_modifiers = {
		.leds = modifiers->leds,
		.mods = modifiers->mods
	};
	return _go_handle_pointer_scroll(view, time, &nc_modifiers, axis_bits, amount);
}

bool handle_pointer_motion(wlc_handle view, uint32_t time, const struct wlc_point *point) {
	struct wlc_point nc_point = {
		.x = point->x,
		.y = point->y
	};
	return _go_handle_pointer_motion(view, time, &nc_point);
}

/* touch */
bool handle_touch_touch(wlc_handle view, uint32_t time, const struct wlc_modifiers *modifiers, enum wlc_touch_type touch, int32_t slot, const struct wlc_point *point) {
	struct wlc_modifiers nc_modifiers = {
		.leds = modifiers->leds,
		.mods = modifiers->mods
	};
	struct wlc_point nc_point = {
		.x = point->x,
		.y = point->y
	};
	return _go_handle_touch_touch(view, time, &nc_modifiers, touch, slot, &nc_point);
}

/* compositor */
void handle_compositor_ready(void) {
	_go_handle_compositor_ready();
}

void handle_compositor_terminate(void) {
	_go_handle_compositor_terminate();
}

/* input */
bool handle_input_created(struct libinput_device *device) {
	return _go_handle_input_created(device);
}

void handle_input_destroyed(struct libinput_device *device) {
	_go_handle_input_destroyed(device);
}


struct wlc_interface interface_wlc;

/**
 * Enable interface functions based on the mask.
 */
void init_interface(uint32_t mask) {
	/* output */

	if ((mask & (1 << 0)) != 0) {
		interface_wlc.output.created = handle_output_created;
	}

	if ((mask & (1 << 1)) != 0) {
		interface_wlc.output.destroyed = handle_output_destroyed;
	}

	if ((mask & (1 << 2)) != 0) {
		interface_wlc.output.focus = handle_output_focus;
	}

	if ((mask & (1 << 3)) != 0) {
		interface_wlc.output.resolution = handle_output_resolution;
	}

	if ((mask & (1 << 4)) != 0) {
		interface_wlc.output.render.pre = handle_output_pre_render;
	}

	if ((mask & (1 << 5)) != 0) {
		interface_wlc.output.render.post = handle_output_post_render;
	}

	/* view */

	if ((mask & (1 << 6)) != 0) {
		interface_wlc.view.created = handle_view_created;
	}

	if ((mask & (1 << 7)) != 0) {
		interface_wlc.view.destroyed = handle_view_destroyed;
	}

	if ((mask & (1 << 8)) != 0) {
		interface_wlc.view.focus = handle_view_focus;
	}

	if ((mask & (1 << 9)) != 0) {
		interface_wlc.view.move_to_output = handle_view_move_to_output;
	}

	if ((mask & (1 << 10)) != 0) {
		interface_wlc.view.request.geometry = handle_view_geometry_request;
	}

	if ((mask & (1 << 11)) != 0) {
		interface_wlc.view.request.state = handle_view_state_request;
	}

	if ((mask & (1 << 12)) != 0) {
		interface_wlc.view.request.move = handle_view_move_request;
	}

	if ((mask & (1 << 13)) != 0) {
		interface_wlc.view.request.resize = handle_view_resize_request;
	}

	if ((mask & (1 << 14)) != 0) {
		interface_wlc.view.render.pre = handle_view_pre_render;
	}

	if ((mask & (1 << 15)) != 0) {
		interface_wlc.view.render.post = handle_view_post_render;
	}

	/* keyboard */

	if ((mask & (1 << 16)) != 0) {
		interface_wlc.keyboard.key = handle_keyboard_key;
	}

	/* pointer */

	if ((mask & (1 << 17)) != 0) {
		interface_wlc.pointer.button = handle_pointer_button;
	}

	if ((mask & (1 << 18)) != 0) {
		interface_wlc.pointer.scroll = handle_pointer_scroll;
	}

	if ((mask & (1 << 19)) != 0) {
		interface_wlc.pointer.motion = handle_pointer_motion;
	}

	/* touch */

	if ((mask & (1 << 20)) != 0) {
		interface_wlc.touch.touch = handle_touch_touch;
	}

	/* compositor */

	if ((mask & (1 << 21)) != 0) {
		interface_wlc.compositor.ready = handle_compositor_ready;
	}

	if ((mask & (1 << 22)) != 0) {
		interface_wlc.compositor.terminate = handle_compositor_terminate;
	}

	/* input */

	if ((mask & (1 << 23)) != 0) {
		interface_wlc.input.created = handle_input_created;
	}

	if ((mask & (1 << 24)) != 0) {
		interface_wlc.input.destroyed = handle_input_destroyed;
	}
}
