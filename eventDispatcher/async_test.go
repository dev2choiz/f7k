package eventDispatcher_test

import (
	"fmt"
	"github.com/dev2choiz/f7k/eventDispatcher"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/stretchr/testify/assert"
	_ "net/http/pprof"
	"sort"
	"strings"
	"sync"
	"testing"
)

var TYPE_INT = 1
var TYPE_ALPHA = 2
var TYPE_ALPHA_MIN = 3

func TestDispatchAsync(t *testing.T) {
	want := "ok"
	eventName := "test-dispatch-async"
	d := svc()

	go func() {
		ch, l, err := d.ListenAsync(eventName, "listenerOne")
		if err != nil {
			panic(err)
		}

		event := <- ch
		l.Done()
		d.CloseListener(eventName, "listenerOne")

		assert.Equal(t, "ko", event.Data().(string))
		assert.Equal(t, "ko", l.Payload())
		l.SetPayload(want)
	}()

	svc().InitDispatcher(eventName, "dispatcherOne")
	d.WaitUntilAsyncListeners(eventName)
	ev := &events.AsyncEvent{}
	ev.SetData("ko")
	evMd, _ := d.DispatchAsync(eventName, "dispatcherOne", ev)
	evMd.Wait()

	_ = d.StopDispatcher(eventName, "dispatcherOne")
	got := evMd.Payload().(string)

	assert.Equal(t, want, got)
}

func TestDispatchAsyncMultiple(t *testing.T) {
	want := []string{"1", "2", "3", "4"}
	eventName := "test-dispatch-async-multiple"

	c := launchListening(eventName, "listenerOne")
	svc().InitDispatcher(eventName, "dispOne")
	_ = loopDispatch(eventName, "dispOne", 4, TYPE_INT, false, false)

	r := <-c
	got := format(r)
	assert.Equal(t, want, got)
}

// Three goroutines will listen same event
// Three goroutines will will dispatch on the same event, one with numbers, the others with letters
func TestDispatchAsyncWithConcurrency(t *testing.T) {
	eventName := "test-dispatch-async-with-concurrency"
	nbLoopOne := 4
	dispOne := "dispatcherOne"
	nbLoopTwo := 8
	dispTwo := "dispatcherTwo"
	nbLoopThree := 8
	dispThree := "dispatcherThree"

	c1 := launchListening(eventName, "listenerOne")
	c2 := launchListening(eventName, "listenerTwo")
	c3 := launchListening(eventName, "listenerThree")

	svc().InitDispatcher(eventName, dispOne)
	svc().InitDispatcher(eventName, dispTwo)
	svc().InitDispatcher(eventName, dispThree)
	wg1 := loopDispatch(eventName, dispOne, nbLoopOne, TYPE_INT, false, false)
	wg2 := loopDispatch(eventName, dispTwo, nbLoopTwo, TYPE_ALPHA,  false, false)
	wg3 := loopDispatch(eventName, dispThree, nbLoopThree, TYPE_ALPHA_MIN,  false, false)

	wg1.Wait()
	wg2.Wait()
	wg3.Wait()

	r1:= <-c1
	r2 := <-c2
	r3 := <-c3

	want := []string{"1", "2", "3", "4", "A", "B", "C", "D", "E", "F", "G", "H", "a", "b", "c", "d", "e", "f", "g", "h"}
	got1 := format(r1)
	got2 := format(r2)
	got3 := format(r3)

	nb := nbLoopOne + nbLoopTwo + nbLoopThree
	assert.Equal(t, nb, len(got1))
	assert.Equal(t, nb, len(got2))
	assert.Equal(t, nb, len(got2))
	assert.Equal(t, want, got1)
	assert.Equal(t, want, got2)
	assert.Equal(t, want, got3)
}

func launchListening(eventName, listenerId string) chan string {
	d := svc()
	mux := &sync.Mutex{}

	c := make(chan string)
	go func() {
		r := ""
		for i := 0; i <= 32; i++ {
			ch, l, err := d.ListenAsync(eventName, listenerId)
			if err != nil {
				if err == eventDispatcher.EOD {
					break
				} else if err == eventDispatcher.EOD_WAIT {
					ch, l, err = d.WaitDisAndListen(eventName, listenerId)
				} else {
					panic(err)
				}
			}

			mux.Lock()
			event := <-ch
			if nil == event {
				mux.Unlock()
				break
			}
			r += fmt.Sprintf("%v ", event.(*events.AsyncEvent).Data())
			l.Done()
			mux.Unlock()
		}
		d.CloseListener(eventName, listenerId)
		mux.Lock()
		c <- r
		mux.Unlock()
	}()

	return c
}

func loopDispatch(eventName, dispId string, loop, typ int, wait, stopAfter bool) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		d := svc()
		d.WaitUntilAsyncListeners(eventName)
		for i := 1; i <= loop; i++ {
			ev := &events.AsyncEvent{}
			switch typ {
			case TYPE_INT:
				ev.SetData(i) // dispatch 1, 2, 3, 4, ...
				break
			case TYPE_ALPHA:
				ev.SetData(string(64 + i)) // dispatch A, B, C, D, ...
				break
			case TYPE_ALPHA_MIN:
				ev.SetData(string(96 + i)) // dispatch a, b, c, d, ...
				break
			}

			aem, err := d.DispatchAsync(eventName, dispId, ev)
			if err != nil {
				panic(err)
			}

			if wait {
				aem.Wait()
			}
		}
		_ = d.StopDispatcher(eventName, dispId)

		wg.Done()
	}(wg)

	return wg
}

func TestWaitUntilAsyncListenersWithNoListener(t *testing.T) {
	d := svc()
	em := &AsyncEventMetadataMock{
		HasListener: true,
	}
	d.AsyncEventsMetadata["dummy"] = em
	_, _, _ = d.ListenAsync("eventName", "listenerId")
	d.WaitUntilAsyncListeners("eventName")
	em.AssertNotCalled(t, "ListenersWaiter")
}

func TestWaitListenDoneError(t *testing.T) {
	em := &AsyncEventMetadataMock{
		mAllDispatchersEnd: true,
	}
	e := eventDispatcher.ListenDoneError(em)

	assert.Equal(t, e, eventDispatcher.EOD)
}

func TestWaitListenDoneErrorWithADispatcherWaiting(t *testing.T) {
	em := &AsyncEventMetadataMock{
		mAllDispatchersEnd: false,
	}
	e := eventDispatcher.ListenDoneError(em)

	assert.Equal(t, e, eventDispatcher.EOD_WAIT)
}

func format(in string) []string {
	r := strings.Split(strings.TrimSpace(in), " ")
	sort.Strings(r)

	return r
}