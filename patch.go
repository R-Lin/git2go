package git

/*
#include <git2.h>
*/
import "C"
import (
	"runtime"
)

type Patch struct {
	ptr *C.git_patch
}

func newPatchFromC(ptr *C.git_patch) *Patch {
	if ptr == nil {
		return nil
	}

	patch := &Patch{
		ptr: ptr,
	}

	runtime.SetFinalizer(patch, (*Patch).Free)
	return patch
}

func (patch *Patch) Free() error {
	if patch.ptr == nil {
		return ErrInvalid
	}
	runtime.SetFinalizer(patch, nil)
	C.git_patch_free(patch.ptr)
	return nil
}

func (patch *Patch) String() (string, error) {
	if patch.ptr == nil {
		return "", ErrInvalid
	}
	var buf C.git_buf
	C.git_patch_to_buf(&buf, patch.ptr)
	return C.GoString(buf.ptr), nil
}
