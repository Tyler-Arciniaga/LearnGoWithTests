package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("t")
	}
}
