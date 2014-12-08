package workman

import "sync"

// Worker Definition
type Worker struct {
	tasks     chan Task
	waitGroup sync.WaitGroup
	fn        func(Task)
}

// DefineWorker defines a new worker group
func DefineWorker(fn func(Task)) *Worker {
	return &Worker{
		tasks: make(chan Task),
		fn:    fn,
	}
}

// Spawn some workers
func (w *Worker) Spawn(count int) *Worker {
	for i := 0; i < count; i++ {
		w.waitGroup.Add(1)
		go func(worker int) {
			for task := range w.tasks {
				w.fn(task)
			}
			w.waitGroup.Done()
		}(i)
	}
	return w
}

// AddTask to the worker queue
func (w *Worker) AddTask(task Task) *Worker {
	w.tasks <- task
	return w
}

// CloseTasks tells us that we are donw with our tasks
func (w *Worker) CloseTasks() *Worker {
	close(w.tasks)
	return w
}

// Finish waits for all the workers to finish
func (w *Worker) Finish() {
	w.CloseTasks().Wait()
}

// Wait for the workers to finish
func (w *Worker) Wait() *Worker {
	w.waitGroup.Wait()
	return w
}
