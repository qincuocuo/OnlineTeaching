package mbase

import (
	"testing"
)

func TestGetPageStart(t *testing.T) {
	testPageStruct := []struct {
		Page      int
		PageSize  int
		PageStart int
	}{
		{
			Page:      108,
			PageSize:  10,
			PageStart: 1070,
		},
	}
	for _, item := range testPageStruct {
		if item.PageStart != GetPageStart(item.Page, item.PageSize) {
			t.Error("TestGetPageStart result error", item)
		}
	}
}
