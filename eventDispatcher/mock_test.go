package eventDispatcher_test

import (
	"github.com/dev2choiz/f7k/eventDispatcher"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/stretchr/testify/mock"
	"sync"
)

type AsyncListenerMetadataMock struct {
	*eventDispatcher.AsyncListenerMetadata
	PipelineClosedMock bool
}

type AsyncEventMetadataMock struct {
	mock.Mock
	HasListener bool
	mAllDispatchersEnd bool
}

func (a AsyncEventMetadataMock) WaitGroup() *sync.WaitGroup {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) Wait() {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) Payload() interface{} {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) SetPayload(interface{}) {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) HasListeners() bool {
	return a.HasListeners()
}

func (a AsyncEventMetadataMock) SetHasListeners(v bool) {
	a.HasListener = v
}

func (a AsyncEventMetadataMock) ListDispatchers() []string {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) SetListDispatchers([]string) {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) ListenersWaiter() chan bool {
	_ = a.Mock.Called()
	return nil
}

func (a AsyncEventMetadataMock) AsyncListenerMetadata() map[string]interfaces.AsyncListenerMetadata {
	panic("should not be called")
}

func (a AsyncEventMetadataMock) AllDispatchersEnd() bool {
	return a.mAllDispatchersEnd
}
