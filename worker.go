package libra

import (
	"sync"
)

/*
Worker !
*/
type Worker interface {
	ExecTask(job Job, context WorkerContext)
}

/*
WorkerContext is context of worker.
*/
type WorkerContext struct {
	Reporter         Reporter
	InitializeRunner Runner
	Runner           Runner
}

/*
SequentialWorker works sequentially.
*/
type SequentialWorker struct {
}

/*
ExecTask execute subtasks one by one.
*/
func (worker SequentialWorker) ExecTask(job Job, context WorkerContext) {
	reporter := context.Reporter
	if reporter == nil {
		reporter = nothingReporter{}
	}
	initRunner := context.InitializeRunner
	runner := context.Runner

	reporter.ReportStart(job)
	defer reporter.ReportEnd()
	initializer := job.Initializer()
	status := initRunner.Exec(initializer)
	if status.Code != OK {
		reporter.Report(initializer, status)
		return
	}
	tasks := job.Subtasks()
	for _, v := range tasks {
		status := runner.Exec(v)
		reporter.Report(v, status)
	}
}

/*
ConcurrentWorker works cuncurrently.
*/
type ConcurrentWorker struct {
	Num int
}

/*
ExecTask execute subtasks concurrently.
*/
func (worker ConcurrentWorker) ExecTask(job Job, context WorkerContext) {
	reporter := context.Reporter
	if reporter == nil {
		reporter = nothingReporter{}
	}
	initRunner := context.InitializeRunner
	runner := context.Runner
	reporter.ReportStart(job)
	defer reporter.ReportEnd()
	initializer := job.Initializer()
	status := initRunner.Exec(initializer)
	if status.Code != OK {
		reporter.Report(initializer, status)
		return
	}
	tasks := job.Subtasks()

	lim := make(chan struct{}, worker.Num)
	wg := &sync.WaitGroup{}
	for _, u := range tasks {
		wg.Add(1)
		go func(v Task) {
			lim <- struct{}{}
			defer wg.Done()
			status := runner.Exec(v)
			reporter.Report(v, status)
			<-lim
		}(u)
	}
	wg.Wait()
	close(lim)
}
