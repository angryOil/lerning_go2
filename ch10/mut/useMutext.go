package mut

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type mutexScoreboardManger struct {
	l     sync.RWMutex
	board map[string]int
}

type scoreboard map[string]int

func newMutexScoreBoardManger() *mutexScoreboardManger {
	return &mutexScoreboardManger{
		board: scoreboard{},
	}
}

func (msm *mutexScoreboardManger) update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.board[name] = val
}

func (msm *mutexScoreboardManger) read(name string) (int, bool) {
	time.Now().Equal(time.Now())
	var a = &http.Client{
		Timeout: 30 * time.Second,
	}
	fmt.Println(a)
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.board[name]
	return val, ok
}
