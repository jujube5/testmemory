package testmemory

import (
	"testing"
)

func TestRedisCache(t *testing.T) {
	RedisSet("test", "qwerty")
	get_value := RedisGet("test")
	if get_value != "qwerty" {
		Colorize(RedBrightFont, "RedisGet -- FAIL!")
		t.Errorf("Result was incorrect, got: %s, want: %s.", get_value, "test")
	} else {
		Colorize(MagentaBrightFont, "RedisSet -- OK!")
		Colorize(MagentaBrightFont, "RedisGet -- OK!")
	}
	RedisDelete("test")
	get_value_2 := RedisGet("test")
	if get_value_2 != "empty" || get_value_2 == "error" {
		Colorize(RedBrightFont, "RedisDelete -- FAIL!")
		t.Errorf("Result was incorrect, got: %s, want: %s.", get_value, "empty")
	} else {
		Colorize(MagentaBrightFont, "RedisDelete -- OK!")
	}
}

func BenchmarkRedisCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisSet("bench", "qwerty")
		RedisGet("bench")
		RedisDelete("bench")
	}
}
