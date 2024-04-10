package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeating a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepetition(t, repeated, expected)
	})

	t.Run("repeating foo 3 times", func(t *testing.T) {
		repeated := Repeat("foo", 3)
		expected := "foofoofoo"
		assertCorrectRepetition(t, repeated, expected)
	})
}

func assertCorrectRepetition(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}