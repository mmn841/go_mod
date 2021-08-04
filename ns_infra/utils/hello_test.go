package utils

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello World,v0.1.11"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

