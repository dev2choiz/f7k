package eventDispatcher

import (
	"errors"
	"fmt"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"io"
	"math/rand"
	"sync"
)

var EOD = io.EOF
var EOD_WAIT = errors.New("need wait another dispatcher")

var muxDis = &sync.Mutex{}

type asyncEventMetadata struct {
	asyncListeners  map[string]interfaces.AsyncListenerMetadata
	waitGroup       *sync.WaitGroup
	payload         interface{}
	hasListeners    bool
	listenersWaiter chan bool
	listenersWaiterMux *sync.Mutex
	dispatchers     map[string]*dispatcher
	name            string
	mux             *sync.Mutex
	dispatchDone    bool
	listeningDone   bool
	nbDispatchers   int
}

type AsyncListenerMetadata struct {
	id            string
	pipelines     map[string]*pipeline
	eventMetadata interfaces.AsyncEventMetadata
	validTokens   []string
	tokenWaiter   chan bool
	waitGroup     *sync.WaitGroup
	mux           *sync.Mutex
	done          bool
	nbReq         int
	nbAllReq      int
}

type pipeline struct {
	ch       chan interfaces.AsyncEvent
	token    *pipToken
	consumed bool
}

type pipToken struct {
	value    string
	owned    bool
}

type dispatcher struct {
	id                 string
	firstDisWaiter2    chan struct{}
	firstDisWaiter    *sync.WaitGroup
	done    		  bool
	firstReq  bool
}

var StopDispatcherEvent *events.AsyncEvent

func init() {
	StopDispatcherEvent = &events.AsyncEvent{}
	StopDispatcherEvent.SetStopPropagation(true)
	StopDispatcherEvent.SetData("StopDispatcherEvent")
}

/**
Dispatch an event struct through a event channel.
You can dispatch several events struct for a same event.
To tell the event is finished, give nil instead of event struct as second argument.
Nothing will be done if there isn't listener registered for the event.
*/
func (ed *EventDispatcher) DispatchAsync(eventName string, dispName string, event interfaces.AsyncEvent) (interfaces.AsyncEventMetadata, error) {
	ed.mux.Lock()
	defer ed.mux.Unlock()
	ed.initEvent(eventName)

	tmp := ed.AsyncEventsMetadata[eventName]
	aem := tmp.(*asyncEventMetadata)

	aem.mux.Lock()

	dis := aem.dispatchers[dispName]
	if !dis.firstReq {
		dis.firstReq = true
	}

	if event == nil || event == StopDispatcherEvent {
		dis.done = true
		aem.nbDispatchers = aem.countActiveDispatchers()
		if aem.AllDispatchersEnd() {
			aem.dispatchDone = true
		}
		aem.mux.Unlock()

		return nil, nil
	}

	aem.nbDispatchers = aem.countActiveDispatchers()
	aem.SetPayload(event.Data())
	aem.mux.Unlock()

	ed.mux.Unlock()
	wg := ed.doDispatchAsync(aem, event, eventName, dispName)
	ed.mux.Lock()

	_ = wg
	return aem, nil
}

func (ed *EventDispatcher) doDispatchAsync(aem *asyncEventMetadata, event interfaces.AsyncEvent, eventName, dispName string) *sync.WaitGroup {
	muxDis.Lock()
	defer muxDis.Unlock()

	wg := &sync.WaitGroup{}
	for _, l := range aem.asyncListeners {
		listener := l.(*AsyncListenerMetadata)
		listener.nbReq++
		listener.nbAllReq++
		aem.WaitGroup().Add(1)

		wg.Add(1)
		ed.initListener(eventName, listener.Id())
		go func(listener *AsyncListenerMetadata) {
			<-listener.tokenWaiter

			listener.mux.Lock()
			tok := listener.validTokens[0]
			pip := listener.pipelines[tok]
			delete(listener.pipelines, tok)
			listener.validTokens = RemoveString(listener.validTokens, tok)
			pip.consumed = true
			listener.mux.Unlock()

			pip.ch <- event
			wg.Done()
		}(listener)
	}

	return wg
}

/**
This function return an event received from the service dispatching.
You need to call it several times until you receive a nil value instead of event.
Do not listen a event if you are not sure it will be dispatched, because it's blocking.
*/
func (ed *EventDispatcher) ListenAsync(eventName, listenerId string) (chan interfaces.AsyncEvent, interfaces.AsyncListenerMetadata, error) {
	ed.mux.Lock()
	defer ed.mux.Unlock()
	ed.initListener(eventName, listenerId)

	aem := ed.AsyncEventsMetadata[eventName].(*asyncEventMetadata)
	tmp := aem.asyncListeners[listenerId]
	l := tmp.(*AsyncListenerMetadata)

	if l.done {
		ch := make(chan interfaces.AsyncEvent)
		close(ch)

		return ch, l, ListenDoneError(aem)
	}

	l.nbReq--

	b := make([]byte, 4)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	l.mux.Lock()
	pip := &pipeline{
		ch: make(chan interfaces.AsyncEvent),
		token :   &pipToken{
			value: token,
		},
	}
	l.pipelines[token] = pip
	l.validTokens = append(l.validTokens, token)

	ch := pip.ch
	l.mux.Unlock()

	l.WaitGroup().Add(1)

	aem.mux.Lock()
	ok := aem.HasListeners()
	aem.mux.Unlock()
	if !ok {
		go func() {
			aem.ListenersWaiter() <- true
		}()
	}

	ed.mux.Unlock()
	go func() {
		l.tokenWaiter <- true
	}()
	ed.mux.Lock()

	defer postListen(l)
	return ch, l, nil
}

func postListen(l *AsyncListenerMetadata) {
	aem := l.eventMetadata.(*asyncEventMetadata)
	if !aem.dispatchDone {
		return
	}

	if l.nbReq > 0 {
		return
	}
	l.done = true

	if aem.countActiveDispatchers() > 0 {
		return
	}
	aem.listeningDone = true
}

func (ed *EventDispatcher) WaitDisAndListen(eventName, listenerId string) (chan interfaces.AsyncEvent, interfaces.AsyncListenerMetadata, error) {
	ed.mux.Lock()
	defer ed.mux.Unlock()

	ch := make(chan struct{})
	aem := ed.AsyncEventsMetadata[eventName].(*asyncEventMetadata)
	for _, d := range aem.dispatchers {
		if d.done {
			continue
		}
		cwg := *d.firstDisWaiter
		go func(cwg sync.WaitGroup) {
			cwg.Wait()
			ch <- struct{}{}
		}(cwg)
	}

	<-ch
	close(ch)
	return ed.ListenAsync(eventName, listenerId)
}

func (ed *EventDispatcher) initEvent(eventName string) {
	if _, ok := ed.AsyncEventsMetadata[eventName]; ok {
		return
	}
	wg := &sync.WaitGroup{}
	ed.AsyncEventsMetadata[eventName] = &asyncEventMetadata{
		asyncListeners:  map[string]interfaces.AsyncListenerMetadata{},
		waitGroup:       wg,
		hasListeners:    false,
		listenersWaiter: make(chan bool),
		listenersWaiterMux: &sync.Mutex{},
		mux: 			 &sync.Mutex{},
		name:            eventName,
		dispatchers:     make(map[string]*dispatcher),
	}
}

func (ed *EventDispatcher) initListener(eventName, listenerName string) {
	ed.initEvent(eventName)
	em := ed.AsyncEventsMetadata[eventName]

	if _, ok := em.AsyncListenerMetadata()[listenerName]; ok {
		return
	}

	wg := &sync.WaitGroup{}
	em.AsyncListenerMetadata()[listenerName] = &AsyncListenerMetadata{
		id:       listenerName,
		pipelines: make(map[string]*pipeline, 0),
		waitGroup:  wg,
		eventMetadata: em,
		tokenWaiter: make(chan bool),
		mux: 			 &sync.Mutex{},
	}
}

func (ed *EventDispatcher) StopDispatcher(eventName, disp string) error {
	_, err := ed.DispatchAsync(eventName, disp, StopDispatcherEvent)
	return err
}

func (ed *EventDispatcher) WaitUntilAsyncListeners(e string) {
	ed.mux.Lock()
	ed.initEvent(e)
	tmp := ed.AsyncEventsMetadata[e]
	aem := tmp.(*asyncEventMetadata)
	ed.mux.Unlock()


	aem.listenersWaiterMux.Lock()
	defer aem.listenersWaiterMux.Unlock()

	aem.mux.Lock()
	has := aem.HasListeners()
	aem.mux.Unlock()

	if has {
		return
	}

	<-aem.ListenersWaiter()

	aem.mux.Lock()
	defer aem.mux.Unlock()
	aem.SetHasListeners(true)
}

func (ed *EventDispatcher) InitDispatcher(eventName, disId  string) {
	ed.initEvent(eventName)
	aem := ed.AsyncEventsMetadata[eventName].(*asyncEventMetadata)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	dis := &dispatcher{
		id: disId,
		done: false,
		firstReq: false,
		firstDisWaiter: wg,
	}
	aem.dispatchers[disId] = dis
}

func (ae *asyncEventMetadata) Wait() {
	ae.waitGroup.Wait()
}

func (ae *asyncEventMetadata) WaitGroup() *sync.WaitGroup {
	return ae.waitGroup
}

func (ae *asyncEventMetadata) Payload() interface{} {
	return ae.payload
}

func (ae *asyncEventMetadata) SetPayload(v interface{}) {
	ae.payload = v
}

func (al AsyncListenerMetadata) Id() string {
	return al.id
}

func (ae *asyncEventMetadata) HasListeners() bool {
	return ae.hasListeners
}

func (ae *asyncEventMetadata) SetHasListeners(v bool) {
	ae.hasListeners = v
}

func (ae *asyncEventMetadata) ListenersWaiter() chan bool {
	return ae.listenersWaiter
}

func (ae *asyncEventMetadata) AllDispatchersEnd() bool {
	for _, d := range ae.dispatchers {
		if !d.done {
			return false
		}
	}
	return true
}

func (ae *asyncEventMetadata) countActiveDispatchers() int {
	n := 0
	for _, d := range ae.dispatchers {
		if d.done {
			continue
		}
		n++
	}
	return n
}

func (ae *asyncEventMetadata) AsyncListenerMetadata() map[string]interfaces.AsyncListenerMetadata {
	return ae.asyncListeners
}

func (al AsyncListenerMetadata) Done() {
	al.waitGroup.Done()
	al.eventMetadata.WaitGroup().Done()
}

func (ed *EventDispatcher) CloseListener(e, l string) {
	ed.mux.Lock()
	defer ed.mux.Unlock()
	aem, ok := ed.AsyncEventsMetadata[e].(*asyncEventMetadata)
	if !ok {
		return
	}
	lis, ok := aem.asyncListeners[l].(*AsyncListenerMetadata)
	if !ok {
		return
	}
	for _, pip := range lis.pipelines {
		close(pip.ch)
	}
}

func (al AsyncListenerMetadata) Payload() interface{} {
	return al.eventMetadata.Payload()
}

func (al AsyncListenerMetadata) SetPayload(v interface{}) {
	al.eventMetadata.SetPayload(v)
}

func (al AsyncListenerMetadata) WaitGroup() *sync.WaitGroup {
	return al.waitGroup
}

func ListenDoneError(aem interfaces.AsyncEventMetadata) error {
	if aem.AllDispatchersEnd() {
		return EOD
	}

	return EOD_WAIT
}
