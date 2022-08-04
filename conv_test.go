package conv

import (
	"fmt"
	"testing"
)

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
		t.Errorf("expect a3 conversion failed")
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
		t.Errorf("expect a3 is converted into 20 when conversion failed")
	}
}

func TestStr2Uint(t *testing.T) {
	var a1 string = "10"
	if res, err := Str2Uint[uint](a1); err != nil || res != 10 {
		t.Errorf("expect a1 is converted into 10")
	}

	var a2 string = "10"
	if res, err := Str2Uint[uint64](a2); err != nil || res != uint64(10) {
		t.Errorf("expect a2 is converted into 10")
	}

	var a3 string = "a-10"
	if _, err := Str2Uint[uint32](a3); err == nil {
		t.Errorf("expect a3 conversion failed")
	}
}

func TestStr2UintOrElse(t *testing.T) {
	var a1 string = "10"
	if res := Str2UintOrElse(a1, uint(20)); res != 10 {
		t.Errorf("expect a1 is converted into 10")
	}

	var a2 string = "10"
	if res := Str2UintOrElse(a2, uint64(20)); res != uint64(10) {
		t.Errorf("expect a2 is converted into 10")
	}

	var a3 string = "a-10"
	if res := Str2UintOrElse(a3, uint16(20)); res != uint16(20) {
		t.Errorf("expect a3 is converted into 20 when conversion failed")
	}
}

func TestFloat2Str(t *testing.T) {
	var a1 float32 = 0.645
	if res := Float2Str(a1); res != "0.645" {
		t.Errorf("expect a1 is converted into \"0.645\"")
	}

	var a2 float32 = -0.654353534534
	if res := Float2Str(a2, 4); res != "-0.6544" {
		fmt.Println(res)
		t.Errorf("expect a2 is converted into \"-0.6543\"")
	}

	var a3 float64 = -34534.9823567982375
	if res := Float2Str(a3, 4); res != "-34534.9824" {
		fmt.Println(res)
		t.Errorf("expect a3 is converted into \"-34534.9824\"")
	}
}

func TestStr2Float(t *testing.T) {
	var a1 string = "10.567"
	if res, err := Str2Float[float32](a1); err != nil || res != 10.567 {
		t.Errorf("expect a1 is converted into 10.567")
	}

	var a2 string = "-4325.1241414"
	if res, err := Str2Float[float64](a2); err != nil || res != float64(-4325.1241414) {
		t.Errorf("expect a2 is converted into -4325.1241414")
	}
}

func TestStr2FloatOrElse(t *testing.T) {
	var a1 string = "-10.567"
	if res := Str2FloatOrElse(a1, 3.0); res != -10.567 {
		t.Errorf("expect a1 is converted into -10.567")
	}

	var a2 string = "-aa325.1241414"
	if res := Str2FloatOrElse(a2, float64(3.0)); res != float64(3.0) {
		t.Errorf("expect a2 is converted into 3.0 when conversion failed")
	}
}

func TestBool2Str(t *testing.T) {
	var a1 bool = true
	if res := Bool2Str(a1); res != "true" {
		t.Errorf("expect a1 is converted into \"true\"")
	}

	var a2 bool = false
	if res := Bool2Str(a2); res != "false" {
		t.Errorf("expect a2 is converted into \"false\"")
	}
}

func TestStr2Bool(t *testing.T) {
	var a1 string = "true"
	if res, err := Str2Bool(a1); err != nil || res != true {
		t.Errorf("expect a1 is converted into true")
	}

	var a2 string = "flase"
	if _, err := Str2Bool(a2); err == nil {
		t.Errorf("expect a2 conversion failed")
	}
}

func TestStr2BoolOrElse(t *testing.T) {
	var a1 string = "true"
	if res := Str2BoolOrElse(a1, false); res != true {
		t.Errorf("expect a1 is converted into true")
	}

	var a2 string = "flase"
	if res := Str2BoolOrElse(a2, true); res != true {
		t.Errorf("expect a2 is converted into true when conversion failed")
	}
}
