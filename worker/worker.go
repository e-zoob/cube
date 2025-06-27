package worker

import (
	"github.com/google/uuid"

	"cube/task"
	"cube/utils"
)

type Worker struct {
	Name      string
	Queue     *queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}
