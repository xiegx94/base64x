// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package base64x

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__b64decode uintptr = __native_entry__() + 1064
    _subr__b64encode uintptr = __native_entry__() + 96
)

const (
    _stack__b64decode = 0
    _stack__b64encode = 0
)

var (
    _ = _subr__b64decode
    _ = _subr__b64encode
)

const (
    _ = _stack__b64decode
    _ = _stack__b64encode
)
