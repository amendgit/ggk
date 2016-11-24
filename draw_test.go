package ggk

import "testing"

func TestBitmapXferProc(t *testing.T) {
	var xferProc tBitmapXferProc
	xferProc = new(tBitmapXferDst)
	var ok = false
	_, ok = xferProc.(*tBitmapXferDst)
	if !ok {
		t.Errorf(`BitmapXfer type assertion is not tBitmapXferDst`)
	}

	_, ok = xferProc.(*tBitmapXferClear)
	if ok {
		t.Errorf(`BitmapXfer type assertion is tBitmapXferClear`)
	}
}
