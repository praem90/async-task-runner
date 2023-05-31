package main

type Emitter struct {
	listeners map[string][]func(*Event)
}

type Event struct {
	topic string
	args  []interface{}
}

func (e *Emitter) On(topic string, listeners ...func(*Event)) {
	if _, ok := e.listeners[topic]; !ok {
		e.listeners[topic] = listeners
	} else {
		e.listeners[topic] = append(e.listeners[topic], listeners...)
	}
}

func (e *Emitter) Emit(topic string, args ...interface{}) {
	if listeners, ok := e.listeners[topic]; ok {
		for _, fn := range listeners {
			event := &Event{
				topic: topic,
				args:  args,
			}

			fn(event)
		}
	}
}

func (e *Emitter) GetListeners(topic *string) []func(*Event) {
	return e.listeners[*topic]
}

func NewEmitter() *Emitter {
	emitter := &Emitter{
		listeners: make(map[string][]func(*Event)),
	}

	return emitter
}
