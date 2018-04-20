package dispatcher

type Job interface {
	Do() error
}
type JobChan chan Job        // 任务队列
type WorkerChan chan JobChan // worker池

type Dispatcher struct {
	workerPool WorkerChan
	jobQueue   JobChan
	maxWorkers int
	maxJobs    int
}

func NewDispatcher(maxWorkers int, maxJobs int) *Dispatcher {
	pool := make(WorkerChan, maxWorkers)
	return &Dispatcher{workerPool: pool}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) AddJob(job Job) {

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func(job Job) {
				jobChannel := <-d.workerPool
				jobChannel <- job
			}(job)
		}
	}
}
