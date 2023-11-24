package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeats letter 6 times", func(t *testing.T) {
		result := Repeat("a", 6)
		expected := "aaaaaa"

		assertionCall(t, result, expected)
	})
	t.Run("Repeats letter 3 times", func(t *testing.T) {
		result := Repeat("b", 3)
		expected := "bbb"

		assertionCall(t, result, expected)
	})
}

func assertionCall(t testing.TB, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
