// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/filter/http/gzip/v2/gzip.proto

package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Gzip) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Gzip)
	if !ok {
		that2, ok := that.(Gzip)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMemoryLevel()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMemoryLevel()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMemoryLevel(), target.GetMemoryLevel()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetContentLength()).(equality.Equalizer); ok {
		if !h.Equal(target.GetContentLength()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetContentLength(), target.GetContentLength()) {
			return false
		}
	}

	if m.GetCompressionLevel() != target.GetCompressionLevel() {
		return false
	}

	if m.GetCompressionStrategy() != target.GetCompressionStrategy() {
		return false
	}

	if len(m.GetContentType()) != len(target.GetContentType()) {
		return false
	}
	for idx, v := range m.GetContentType() {

		if strings.Compare(v, target.GetContentType()[idx]) != 0 {
			return false
		}

	}

	if m.GetDisableOnEtagHeader() != target.GetDisableOnEtagHeader() {
		return false
	}

	if m.GetRemoveAcceptEncodingHeader() != target.GetRemoveAcceptEncodingHeader() {
		return false
	}

	if h, ok := interface{}(m.GetWindowBits()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWindowBits()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWindowBits(), target.GetWindowBits()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Gzip_CompressionLevel) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Gzip_CompressionLevel)
	if !ok {
		that2, ok := that.(Gzip_CompressionLevel)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}
