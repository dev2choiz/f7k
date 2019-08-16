package eventDispatcher_test

import (
	"fmt"
	"github.com/dev2choiz/f7k/eventDispatcher"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"testing"
	"time"
)

func TestDispatchAsync(t *testing.T) {
	want := "1-2-3-4-"
	var got string
	eventName := "test-dispatch-async"

	go func() {
		got = launchListening(eventName, "listenerId")
	}()
	time.Sleep(10 * time.Millisecond)
	launchDispatching(eventName)
	time.Sleep(200 * time.Millisecond)
	if want != got {
		t.Errorf("Incorrect result, want: %s, got: %s.", want, got)
	}
}

func TestDispatchAsyncFunc(t *testing.T) {
	want := "1-2-3-4-"
	eventName := "test-dispatch-async-func"
	var got string

	go func() {
		got = launchListening(eventName, "listenerId")
	}()
	time.Sleep(10 * time.Millisecond)

	eventDispatcher.Instance().DispatchAsyncFunc(eventName, func(ch chan interfaces.Event) {
		for i := 1; i <= 4; i++ {
			time.Sleep(10 * time.Millisecond)
			ev := &events.Event{}
			ev.SetData(i)
			ch <- ev
		}
		ch <- nil
	})

	time.Sleep(200 * time.Millisecond)
	if want != got {
		t.Errorf("Result incorrect, want: %s, got: %s.", want, got)
	}
}

func launchDispatching(eventName string) {
	d := eventDispatcher.Instance()
	go func() {
		for i := 1; i <= 4; i++ {
			time.Sleep(10 * time.Millisecond)
			ev := &events.Event{}
			ev.SetData(i)
			d.DispatchAsync(eventName, ev)
		}
		d.DispatchAsync(eventName, nil)
	}()
}

func launchListening(eventName, listenerName string) string {
	r := ""
	for {
		event, err := eventDispatcher.Instance().ListenAsync(eventName, listenerName)
		if nil != err {
			panic(err)
		}
		if nil == event {
			break
		}
		r += fmt.Sprintf("%v-", event.(*events.Event).Data())
	}
	return r
}
