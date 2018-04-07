package main

import "testing"

func TestHello(t *testing.T) {
    expected := `Hello, World!`
    actual   := Hello()

    if actual != expected {
        t.Errorf("[FAILED]\nExpected:\t%s\nActual:\t\t%s\n", expected, actual)
    }
}
