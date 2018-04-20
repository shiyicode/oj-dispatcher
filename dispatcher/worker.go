package dispatcher

import "fmt"

type Worker struct {
	JobChannel JobChan       // 用于接收job
	quit       chan struct{} // 用于接收停止信号
	workerPool WorkerChan    // worker池
}

// worker里面需要包含docker相关
// 初始化时创建docker(docker,judger 的包装)
// 每次收到任务，用docker执行
// 开始任务前，进行检测docker是否可用，不可用即销毁并重新生成  reCreate
// 任务结束时，如果不可用，即销毁容器

func NewWorker(workPool WorkerChan) Worker {
	return Worker{
		JobChannel: make(chan Job),
		quit:       make(chan struct{})}
}

func (w *Worker) Start() {
	go func() {
		for {
			// 将任务接收队列放入worker池
			w.workerPool <- w.JobChannel
			select {
			// 接收到了任务
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					fmt.Printf("excute job failed with err: %v", err)
					// 任务失败，删除docker，重新建立docker
				}

			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- struct{}{}
	}()
}
