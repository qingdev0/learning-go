package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) Status() string {
	return fmt.Sprintf("Total: %d, Last Updated: %s", c.total, c.lastUpdated.Format(time.RFC3339))
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.Status())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.Status())
}

func main() {
	counter := Counter{}
	fmt.Println("in main:", counter.Status())
	doUpdateWrong(counter)
	fmt.Println("in main:", counter.Status())
	doUpdateRight(&counter)
	fmt.Println("in main:", counter.Status())
}
