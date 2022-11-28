package greeting

import "testing"

func testHello(t *testing.T) {
	want := "Bonjour"
	if got := Hello("fr"); got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}
