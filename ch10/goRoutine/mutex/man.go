package main

import "sync"

func main() {

}

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type channelScoreBoardManager chan func(map[string]int)

func newChannelScoreManger() (channelScoreBoardManager, func()) {
	ch := make(channelScoreBoardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm channelScoreBoardManager) update(name string, score int) {
	csm <- func(m map[string]int) {
		m[name] = score
	}
}

func (csm channelScoreBoardManager) read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

type mutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func newMutexScoreboardManger() *mutexScoreboardManager {
	return &mutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *mutexScoreboardManager) update(name string, score int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = score
}

func (msm *mutexScoreboardManager) read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
