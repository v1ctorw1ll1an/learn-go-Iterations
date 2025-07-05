package iteracaco

import (
	"fmt"
	"testing"
)

func TestIteracao(t *testing.T) {
	t.Run("test main functionally", func(t *testing.T) {
		want := "aaaaaaaaaa"
		got := Repeat("a", 10)

		if want != got {
			t.Errorf("want: '%s' got: '%s'", want, got)
		}
	})
}

func BenchmarkIteracao(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 10))
	// Output: aaaaaaaaaa
}
