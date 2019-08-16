package eventDispatcher_test

import (
	"github.com/dev2choiz/f7k/eventDispatcher"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDispatch(t *testing.T) {
	want := "first listener"
	eventName := "test-dispatch"

	d := eventDispatcher.Instance()
	d.Listen( eventName, func(e interfaces.Event) {
		event := e.(*events.Event)
		event.SetData(event.Data().(string) + "first listener")
		event.SetStopPropagation(true)
	})

	d.Listen( eventName, func(e interfaces.Event) {
		event := e.(*events.Event)
		event.SetData(event.Data().(string) + "second listener")
	})

	e := &events.Event{}
	e.SetData("")
	d.Dispatch(eventName, e)
	got := e.Data()

	assert.Equal(t, want, got, "%s should be equal to %s", got, want)
}