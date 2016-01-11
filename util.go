package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>

char **char_array_init(size_t len) {
	char **arr = malloc(len * sizeof(char*));
	return arr;
}

void char_array_insert(char **arr, char *item, int index) {
	arr[index] = item;
}

void char_array_free(char **arr) {
	int i = 0;
	for (char **ptr = arr; *ptr; ++ptr) {
		free(*ptr);
	}
	free(arr);
}
*/
import "C"

import "unsafe"

// Initialize a C NULL terminated *char[] from a []string
func strSlicetoCArray(arr []string) **C.char {
	carr := C.char_array_init(C.size_t(len(arr) + 1))
	for i, s := range arr {
		C.char_array_insert(carr, C.CString(s), C.int(i))
	}
	C.char_array_insert(carr, nil, C.int(len(arr)))
	return carr
}

// Free a *char[]
func freeCStrArray(arr **C.char) {
	C.char_array_free(arr)
}

func handlesCArraytoGoSlice(handles *C.wlc_handle, len int) []Handle {
	goHandles := make([]Handle, 0, len)
	size := int(unsafe.Sizeof(*handles))
	for i := 0; i < len; i++ {
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(handles)) + uintptr(size*i))
		goHandles[i] = *(*Handle)(ptr)
	}

	return goHandles
}
