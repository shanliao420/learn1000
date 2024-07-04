package utils

import (
	"fmt"
	"testing"
)

func TestIDMaker(t *testing.T) {
	idm := NewIDMaker()
	fmt.Println(idm.MakeID())
	// Output: HbV1t0jxorS8EBPOKPIG
}

func TestIDMaker2(t *testing.T) {
	idm := NewIDMaker()
	id := idm.MakeID()
	fmt.Println(id)
	tm := idm.GetTimeFromID(id)
	fmt.Println(tm)
	fmt.Println(tm.Format("2006-01-02"))
}
