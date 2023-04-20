package testmemory

import (
	"testing"
)

func TestSet(t *testing.T) {
	Set("test", "qwerty")
	len_map_cache := len(c.Items)
	if len_map_cache != 1 {
		Colorize(RedBrightFont, "Set -- FAIL!")
		t.Errorf("Result was incorrect, got: %d, want: %d.", len_map_cache, 1)
	} else {
		Colorize(GreenBrightFont, "Set -- OK!")
	}
}

func TestGet(t *testing.T) {
	get_cache := Get("test")
	if get_cache != "qwerty" {
		Colorize(RedBrightFont, "Get -- FAIL!")
		t.Errorf("Result was incorrect, got: %s, want: %s.", get_cache, "qwerty")
	} else {
		Colorize(GreenBrightFont, "Get -- OK!")
	}
}

func TestDelete(t *testing.T) {
	Delete("test")
	get_cache := Get("test")
	if get_cache != nil {
		Colorize(RedBrightFont, "Delete -- FAIL!")
		t.Errorf("Result was incorrect, got: %s, want: %s.", get_cache, "nil")
	} else {
		Colorize(GreenBrightFont, "Delete -- OK!")
	}
}

func BenchmarkCache(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Set("bench", "qwerty")
	Get("bench")
	Delete("bench")
    }
}

func BenchmarkSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Set("kira", 8)
    }
}

func BenchmarkGet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Get("kira")
    }
}
