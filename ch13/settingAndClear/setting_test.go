package settingAndClear

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testingTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("set up in Main")
	testingTime = time.Now()
	fmt.Println("now ", testingTime)
	exitVal := m.Run()
	fmt.Println("clean up stuff after test here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("testFirst uses stuff set up in testMain ", testingTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("testSec uses stuff set up in testMain ", testingTime)
}
