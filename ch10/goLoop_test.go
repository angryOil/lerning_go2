package ch10

import "testing"

func Test_buggedLoop(t *testing.T) {
	buggedLoop()
}

func Test_fixedLoopByGoParam(t *testing.T) {
	fixedLoopByGoParam()
}

func Test_fixedLoopByShadowing(t *testing.T) {
	fixedLoopByShadowing()
}
