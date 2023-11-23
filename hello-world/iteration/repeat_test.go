package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if expected != repeated {
		t.Errorf("Expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 5)
	fmt.Println(repeat)
	// Expected Output: bbbbb
}
