package dispatcher

import (
	"fmt"

	"github.com/open-fightcoder/oj-dispatcher/judger"
	log "github.com/sirupsen/logrus"
)

type JobChan chan *judger.Job // 任务队列
type WorkerChan chan JobChan  // worker池

var (
	quit       chan struct{} // 用于接收停止信号
	workerPool WorkerChan
	jobQueue   JobChan
	workers    []*Worker
)

func Start(maxWorkers int, maxJobs int) {
	workerPool = make(WorkerChan, maxWorkers)
	workers = make([]*Worker, 0)
	jobQueue = make(JobChan, maxJobs)
	quit = make(chan struct{})

	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker()
		workers = append(workers, worker)
		worker.Start()
	}

	go dispatch()
}

func Stop() {

	quit <- struct{}{}

	for _, worker := range workers {
		worker.Stop()
	}

	log.Info("dispatcher stoped")
}

func AddJob(job *judger.Job) {
	jobQueue <- job
}

func dispatch() {
	for {
		fmt.Println("ttt")
		select {
		case job := <-jobQueue:
			go func(job *judger.Job) {
				jobChannel := <-workerPool
				jobChannel <- job
			}(job)
			fmt.Println("xxx")

		case <-quit:
			fmt.Println("yyy")
			return
		}
	}
}
