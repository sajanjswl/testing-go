package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("adding two integer", func(t *testing.T) {

		got := Add(2, 3)
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
