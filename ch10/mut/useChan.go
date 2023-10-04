package mut

type channelScoreboardManager chan func(map[string]int)

func newChannelScoreManager() (channelScoreboardManager, func()) {
	ch := make(channelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManger(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm channelScoreboardManager) read(name string) (int, bool) {
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

func (csm channelScoreboardManager) update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func scoreboardManger(in <-chan func(map[string]int), done <-chan struct{}) {
	board := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(board)
		}
	}
}
