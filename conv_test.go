package conv

import "testing"

func TestInt2Str(t *testing.T) {
	var a1 int = -10
	if res := Int2Str(a1); res != "-10" {
		t.Errorf("expect a1 is converted into \"-10\"")
	}

	var a2 int8 = -10
	if res := Int2Str(a2); res != "-10" {
		t.Errorf("expect a2 is converted into \"-10\"")
	}

	var a3 int16 = -10
	if res := Int2Str(a3); res != "-10" {
		t.Errorf("expect a3 is converted into \"-10\"")
	}

	var a4 int32 = -10
	if res := Int2Str(a4); res != "-10" {
		t.Errorf("expect a4 is converted into \"-10\"")
	}

	var a5 int64 = -10
	if res := Int2Str(a5); res != "-10" {
		t.Errorf("expect a5 is converted into \"-10\"")
	}
}

func TestUint2Str(t *testing.T) {
	var a1 uint = 10
	if res := Uint2Str(a1); res != "10" {
		t.Errorf("expect a1 is converted into \"10\"")
	}

	var a2 uint8 = 10
	if res := Uint2Str(a2); res != "10" {
		t.Errorf("expect a2 is converted into \"10\"")
	}

	var a3 uint16 = 10
	if res := Uint2Str(a3); res != "10" {
		t.Errorf("expect a3 is converted into \"10\"")
	}

	var a4 uint32 = 10
	if res := Uint2Str(a4); res != "10" {
		t.Errorf("expect a4 is converted into \"10\"")
	}

	var a5 uint64 = 10
	if res := Uint2Str(a5); res != "10" {
		t.Errorf("expect a5 is converted into \"10\"")
	}
}

func TestStr2Int(t *testing.T) {
	var a1 string = "10"
	if res, err := Str2Int[int](a1); err != nil || res != 10 {
		t.Errorf("expect a1 is converted into 10")
	}

	var a2 string = "-10"
	if res, err := Str2Int[int64](a2); err != nil || res != int64(-10) {
		t.Errorf("expect a2 is converted into -10")
	}

	var a3 string = "a-10"
	if _, err := Str2Int[int32](a3); err == nil {
		t.Errorf("expect a3 convertion failed")
	}
}

func TestStr2IntOrElse(t *testing.T) {
	var a1 string = "10"
	if res := Str2IntOrElse(a1, 20); res != 10 {
		t.Errorf("expect a1 is converted into 10")
	}

	var a2 string = "-10"
	if res := Str2IntOrElse(a2, int64(20)); res != int64(-10) {
		t.Errorf("expect a2 is converted into -10")
	}

	var a3 string = "a-10"
	if res := Str2IntOrElse(a3, int16(20)); res != int16(20) {
		t.Errorf("expect a3 is converted into 20 when convertion failed")
	}
}
