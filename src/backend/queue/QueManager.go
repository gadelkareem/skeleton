package queue

import (
	"encoding/json"

	"github.com/gadelkareem/que"
)

type QueManager struct {
	c  *que.Client
	wm que.WorkMap
}

type Worker interface {
	Type() string
	Run(j *que.Job) error
}

func NewQueManager(c *que.Client) *QueManager {
	return &QueManager{c: c, wm: make(que.WorkMap)}
}

func (q QueManager) AddWorker(ws ...Worker) {
	for _, w := range ws {
		q.wm[w.Type()] = w.Run
	}
}

func (q QueManager) StartWorkers() *que.WorkerPool {
	p := que.NewWorkerPool(q.c, q.wm, 5)
	p.Start()

	return p
}

func (q QueManager) Enqueue(t string, j interface{}) error {
	enc, err := json.Marshal(j)
	if err != nil {
		return err
	}

	return q.c.Enqueue(&que.Job{
		Type: t,
		Args: enc,
	})
}
