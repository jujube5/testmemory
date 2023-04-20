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
