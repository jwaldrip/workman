package workman

import "sync"

// Worker Definition
type Worker struct {
	tasks chan Task
	wg    sync.WaitGroup
	fn    func(Task)
}

// New defines a new workman group
func New(fn func(Task)) *Worker {
	return &Worker{
		tasks: make(chan Task),
		fn:    fn,
	}
}

// Spawn some workers
func (w *Worker) Spawn(count int) *Worker {
	w.wg.Add(count)
	for i := 0; i < count; i++ {
		go func(worker int) {
			for task := range w.tasks {
				w.fn(task)
			}
			w.wg.Done()
		}(i)
	}
	return w
}

// AddTask to the worker queue
func (w *Worker) AddTask(task Task) {
	w.tasks <- task
}

// Finish waits for all the workers to finish
func (w *Worker) Finish() {
	close(w.tasks)
	w.wg.Wait()
}
