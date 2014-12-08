# Workman

Define and run parallel workers in go

## Installation

```
$ go get github.com/jwaldrip/workman
```

## Usage

Define a function to run when a worker defines a task. It takes an interface, so
you will have to assert its type.

```go
package main

import (
  "fmt"
  "time"

  "github.com/jwaldrip/workman"
)

var worker = workman.DefineWorker(func(context workman.Task) {
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
```

## Contributing

See [CONTRIBUTING](https://github.com/jwaldrip/odin/blob/master/CONTRIBUTING.md) for details on how to contribute.
