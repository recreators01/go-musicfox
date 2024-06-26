// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignaturePlaybackMediaMarkerSequence string = "rc(Windows.Media.Playback.PlaybackMediaMarkerSequence;{f2810cee-638b-46cf-8817-1d111fe9d8c4})"

type PlaybackMediaMarkerSequence struct {
	ole.IUnknown
}

func (impl *PlaybackMediaMarkerSequence) GetSize() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarkerSequence))
	defer itf.Release()
	v := (*iPlaybackMediaMarkerSequence)(unsafe.Pointer(itf))
	return v.GetSize()
}

func (impl *PlaybackMediaMarkerSequence) Insert(value *PlaybackMediaMarker) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarkerSequence))
	defer itf.Release()
	v := (*iPlaybackMediaMarkerSequence)(unsafe.Pointer(itf))
	return v.Insert(value)
}

func (impl *PlaybackMediaMarkerSequence) Clear() error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiPlaybackMediaMarkerSequence))
	defer itf.Release()
	v := (*iPlaybackMediaMarkerSequence)(unsafe.Pointer(itf))
	return v.Clear()
}

const GUIDiPlaybackMediaMarkerSequence string = "f2810cee-638b-46cf-8817-1d111fe9d8c4"
const SignatureiPlaybackMediaMarkerSequence string = "{f2810cee-638b-46cf-8817-1d111fe9d8c4}"

type iPlaybackMediaMarkerSequence struct {
	ole.IInspectable
}

type iPlaybackMediaMarkerSequenceVtbl struct {
	ole.IInspectableVtbl

	GetSize uintptr
	Insert  uintptr
	Clear   uintptr
}

func (v *iPlaybackMediaMarkerSequence) VTable() *iPlaybackMediaMarkerSequenceVtbl {
	return (*iPlaybackMediaMarkerSequenceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iPlaybackMediaMarkerSequence) GetSize() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iPlaybackMediaMarkerSequence) Insert(value *PlaybackMediaMarker) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Insert,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in PlaybackMediaMarker
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iPlaybackMediaMarkerSequence) Clear() error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Clear,
		uintptr(unsafe.Pointer(v)), // this
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
