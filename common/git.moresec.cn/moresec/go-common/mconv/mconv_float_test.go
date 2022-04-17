package mconv

import (
	"reflect"
	"testing"
)

func TestFloat32(t *testing.T) {
	runTest := []struct {
		in   interface{}
		Type reflect.Type
		out  float32
	}{
		{"1.2", reflect.TypeOf(float32(0)), float32(1.2)},
		{"-123", reflect.TypeOf(float32(0)), float32(-123)},
		{123, reflect.TypeOf(float32(0)), float32(123)},
		{float64(12.12), reflect.TypeOf(float32(0)), float32(12.12)},
	}

	for i := range runTest {
		if Float32(runTest[i].in) != runTest[i].out {
			t.Errorf("Float32(%v) != %v", runTest[i].in, runTest[i].out)
		}
		if reflect.TypeOf(Float32(runTest[i].in)) != runTest[i].Type {
			t.Errorf("Float32(%v) != %v", runTest[i].in, runTest[i].Type)
		}
	}
}
