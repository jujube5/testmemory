package testmemory

import (
	"reflect"
	"testing"
)

func TestStringToInt(t *testing.T) {
	str_to_int := StringToInt("1001")
	if str_to_int != 1001 {
		Colorize(RedBrightFont, "StringToInt -- FAIL!")
		t.Errorf("Result was incorrect, got: %d, want: %d.", str_to_int, 1001)
	} else {
		Colorize(CyanBrightFont, "StringToInt -- OK!")
	}
}

func TestParseRaw(t *testing.T) {
	key, value := ParseRaw("MemTotal:        7815804 kB")
	if key == "MemTotal" || value == 7815804 {
		Colorize(CyanBrightFont, "ParseRaw -- OK!")
	} else {
		Colorize(RedBrightFont, "ParseRaw -- FAIL!")
		t.Errorf("Result was incorrect.")
	}
}

func TestReadMemInfo(t *testing.T) {
	meminfo := ReadMemInfo()
	extract_type := reflect.TypeOf(meminfo).Elem()
    extract_name := extract_type.Name()
	if extract_name != "MemInfo" {
		Colorize(RedBrightFont, "ReadMemInfo -- FAIL!")
		t.Errorf("Result was incorrect, got: %s, want: %s.", extract_name, "MemInfo")
	} else {
		Colorize(CyanBrightFont, "ReadMemInfo -- OK!")
	}
}
