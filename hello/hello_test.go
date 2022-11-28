package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world"
	if got := Hello(); got != want {
		t.Errorf("Excpected %s but got %s", want, got)
	}
}

func TestSayHello(t *testing.T) {
    want := "Bonjour"
    if got := SayHello("fr"); want != got {
        t.Errorf("Expected %q but got %q", want, got)
    }
}