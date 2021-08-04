package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello World,v0.1.3"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

