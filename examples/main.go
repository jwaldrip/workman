package main

import (
	"fmt"
	"time"

	"github.com/jwaldrip/workman"
)

var worker = workman.New(func(context workman.Task) {
	str := context.(string)
	fmt.Println(str)
	time.Sleep(time.Second)
})

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"kelly",
	"paul",
	"dina",
	"chris",
	"lisa",
	"tom",
	"travis",
}

func init() {
	worker.Spawn(2)
	for _, name := range names {
		worker.AddTask(name)
	}
}

func main() {
	worker.Finish()
}
