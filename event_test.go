package main

import "testing"

func TestEmitter(t *testing.T) {
	emitter := NewEmitter()

	called := false

	emitter.On("test", func(e *Event) {
		called = true
	})

	emitter.Emit("test", "Emit")

	if called == false {
		t.Fatal("Callback is not getting called as expected")
	}
}

func TestEmitterFail(t *testing.T) {
	emitter := NewEmitter()

	called := false

	emitter.On("emit", func(e *Event) {
		called = true
	})

	emitter.Emit("test", "Emit")

	if called {
		t.Fatal("Callback is getting called for every topic")
	}
}
