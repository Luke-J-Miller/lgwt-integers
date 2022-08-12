package integers

import (
	"fmt"
	"testing"
)

// Tests the Add function in adder.go
func TestAdder(t *testing.T) {
	assertEqual := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}
	t.Run("2 and 2 makes 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertEqual(t, sum, expected)
	})

}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
