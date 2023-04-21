package testmemory

import (
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
	if meminfo.Total < 512 {
		Colorize(RedBrightFont, "ReadMemInfo -- FAIL!")
		t.Errorf("Result was incorrect.")
	} else {
		Colorize(CyanBrightFont, "ReadMemInfo -- OK!")
	}
}

func BenchmarkStringToInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        StringToInt("1001")
    }
}

func BenchmarkParseRaw(b *testing.B) {
    for i := 0; i < b.N; i++ {
        key, value := ParseRaw("MemTotal:        7815804 kB")
    }
}

func BenchmarkReadMemInfo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        meminfo := ReadMemInfo()
    }
}
