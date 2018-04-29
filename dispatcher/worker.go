package dispatcher

import (
	"github.com/open-fightcoder/oj-dispatcher/judger"
)

type Worker struct {
	JobChannel JobChan        // 用于接收job
	quit       chan struct{}  // 用于接收停止信号
	judger     *judger.Judger // 判题实例
	id         int            // 编号
}

// worker里面需要包含docker相关
// 初始化时创建docker(docker,judger 的包装)
// 每次收到任务，用docker执行
// 开始任务前，进行检测docker是否可用，不可用即销毁并重新生成  reCreate
// 任务结束时，如果不可用，即销毁容器

func NewWorker(id int) *Worker {
	worker := new(Worker)
	worker.JobChannel = make(chan *judger.Job)
	worker.quit = make(chan struct{})
	worker.judger = judger.NewJudger(id)
	worker.id = id
	return worker
}

func (w *Worker) Start() {
	go func() {
		for {
			// 将任务接收队列放入worker池
			workerPool <- w.JobChannel
			select {
			// 接收到了任务
			case job := <-w.JobChannel:
				w.judger.Do(job)

			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- struct{}{}
	w.judger.DropDocker()
}
