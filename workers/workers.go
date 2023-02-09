package workers

import (
	"context"
	"sync"

	"gitlab.id.vin/gami/ps2-gami-common/adapters/queue"
)

// Task represents a task
type Task struct {
	ctx      context.Context
	msg      *queue.ConsumerMessage
	executor func(context.Context, *queue.ConsumerMessage) error
}

// NewTask create new task
func NewTask(ctx context.Context, msg *queue.ConsumerMessage, executor func(context.Context, *queue.ConsumerMessage) error) *Task {
	return &Task{
		ctx:      ctx,
		msg:      msg,
		executor: executor,
	}
}

// Execute task
func (t *Task) Execute() {
	if t.executor != nil {
		_ = t.executor(t.ctx, t.msg)
	}
}

// Pool of worker
type Pool struct {
	ctx          context.Context
	cancel       context.CancelFunc
	numberWorker int
	wg           sync.WaitGroup
	ch           chan *Task
}

// NewPool create new worker pool
func NewPool(ctx context.Context, numberWorker int) (p *Pool) {
	if numberWorker <= 0 {
		numberWorker = 1
	}

	if ctx == nil {
		ctx = context.Background()
	}

	p = &Pool{
		numberWorker: numberWorker,
		ch:           make(chan *Task, numberWorker),
	}
	p.ctx, p.cancel = context.WithCancel(ctx)

	return
}

// Start workers
func (p *Pool) Start() {
	p.wg.Add(p.numberWorker)
	for i := 0; i < p.numberWorker; i++ {
		go p.worker()
	}
}

// Do a task
func (p *Pool) Do(t *Task) {
	if p.ch != nil && t != nil {
		select {
		case <-p.ctx.Done():
		case p.ch <- t:
		}
	}
}

// Stop worker. Wait all task done.
func (p *Pool) Stop() {
	// cancel context
	p.cancel()

	// wait child workers
	p.wg.Wait()
}

func (p *Pool) worker() {
	defer p.wg.Done()

	var task *Task

	for {
		select {
		case <-p.ctx.Done():
			return

		case task = <-p.ch:
			if task != nil {
				task.Execute()
			}
		}
	}
}
