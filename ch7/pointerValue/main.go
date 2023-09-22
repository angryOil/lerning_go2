package main

import (
	"fmt"
	"time"
)

func main() {
	c := counter{}
	fmt.Println(c)
	c.inc()
	fmt.Println(c.String())
}

type counter struct {
	total      int
	lastUpdate time.Time
}

func (c *counter) inc() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c counter) String() string {
	return fmt.Sprintf("total :%d,last update%v", c.total, c.lastUpdate)
}
