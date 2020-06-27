package taskrunner

import "time"


type Worker struct {
	ticker *time.Ticker
	runner *Runner



}

func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker:time.NewTicker(interval*time.Second),
		runner:r,
	}
}
//
func (w *Worker) startWorker()  {
	//for c = range w.ticker.C{ 错误写法
	//
	//}

	for {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()


		}
	}


}
//
func Start()  {
	// Start video file cleaning
	r:= NewRuner(3,true,VideoClearDispatcher,VideoClearExecutor)
	w := NewWorker(3,r)
	go w.startWorker()


}
//


//1 user- apiservice - dilete video
//2 api service - scheduler - write video deletion record
//3 timeer
//4 timer - runner - read wvdr - exec - delete video from folder







