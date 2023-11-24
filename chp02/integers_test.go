package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Sum 2+2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertionCall(t, sum, expected)
	})
}

func assertionCall(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("Expected: %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
