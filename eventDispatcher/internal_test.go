package eventDispatcher

import (
	"github.com/stretchr/testify/assert"
	_ "net/http/pprof"

	"testing"
)

func TestPostListen(t *testing.T) {
	em := &asyncEventMetadata{dispatchDone: false, listeningDone: false,}
	lm := &AsyncListenerMetadata{eventMetadata: em, done: false, nbReq: 1,}
	postListen(lm)
	assert.False(t, lm.done)
	assert.False(t, em.listeningDone)

	em = &asyncEventMetadata{dispatchDone: true, listeningDone: false,}
	lm = &AsyncListenerMetadata{eventMetadata: em, done: false, nbReq: 1,}
	postListen(lm)
	assert.False(t, lm.done)
	assert.False(t, em.listeningDone)

	em = &asyncEventMetadata{dispatchDone: true, listeningDone: false,}
	lm = &AsyncListenerMetadata{eventMetadata: em, done: false, nbReq: 0,}
	postListen(lm)
	assert.True(t, lm.done)
	assert.True(t, em.listeningDone)

}

func TestPostListenWithActiveDispatchers(t *testing.T) {
	em := &asyncEventMetadata{
		dispatchDone:  true,
		listeningDone: false,
		dispatchers: map[string]*dispatcher{
			"disp1": &dispatcher{done: false},
			"disp2": &dispatcher{done: false},
		},
	}
	lm := &AsyncListenerMetadata{eventMetadata: em, done: false, nbReq: 0,}
	postListen(lm)
	assert.True(t, lm.done)
	assert.False(t, em.listeningDone)
}

func TestPostListenWithoutActiveDispatchers(t *testing.T) {
	em := &asyncEventMetadata{
		dispatchDone: true,
		listeningDone: false,
		dispatchers: map[string]*dispatcher{
			"disp1": &dispatcher{done: true},
			"disp2": &dispatcher{done: true},
		},
	}
	lm := &AsyncListenerMetadata{eventMetadata: em, done: false, nbReq: 0,}
	postListen(lm)
	assert.True(t, lm.done)
	assert.True(t, em.listeningDone)
}
