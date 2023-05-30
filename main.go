package main

import (
	"fmt"
)

func main() {
	emitter := New()

	emitter.On("test", func(e *Event) {
		fmt.Println(e.args...)
	})

	topic := "test"
	listeners := emitter.GetListeners(&topic)
	fmt.Println("No of listeners", len(listeners))
	emitter.Emit(&topic, "Testing event emiiter")
}
