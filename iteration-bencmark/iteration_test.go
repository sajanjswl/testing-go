package iter

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q want %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
