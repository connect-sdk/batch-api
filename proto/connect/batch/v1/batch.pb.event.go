package batchv1

import (
	"context"
	"errors"
)

// The error occurs when the event is unknown.
var ErrBatchEventTypeUnknown = errors.New("unknown batch event type")

//counterfeiter:generate -o batchv1fake . BatchJobEvent

// BatchJobEvent is the interface for batch job event type BatchJobEvent interface {
type BatchJobEvent interface {
	GetJob() *BatchJob
	SetJob(*BatchJob)
}

var _ BatchJobEvent = &BatchJobStateChangedEvent{}

// GetJob returns the job.
func (x *BatchJobStateChangedEvent) GetJob() *BatchJob {
	return &BatchJob{
		Name: x.JobName,
		Uid:  x.JobUid,
		Status: &BatchJobStatus{
			State: BatchJobStatus_State(x.NewJobState),
		},
	}
}

// SetJob implements BatchJobEvent.
func (x *BatchJobStateChangedEvent) SetJob(job *BatchJob) {
	x.JobName = job.Name
	x.JobUid = job.Uid
	x.NewJobState = BatchJobStateChangedEvent_State(job.Status.GetState())
}

// BatchJobStateChangedEvent_State is the enum type of BatchJobStateChangedEvent.State.
func NewBatchJobStateChangedEvent_State(text string) BatchJobStateChangedEvent_State {
	return BatchJobStateChangedEvent_State(BatchJobStateChangedEvent_State_value[text])
}

//counterfeiter:generate -o batchv1fake . BatchJobEventHandler

// BatchJobEventHandler handles connect.batch.v1.BatchJobEvent event.
type BatchJobEventHandler interface {
	// HandleBatchJobEvent handles connect.batch.v1.BatchJobEvent event.
	HandleBatchJobEvent(context.Context, BatchJobEvent) error
}

//counterfeiter:generate -o batchv1fake . BatchTaskEvent

// BatchTaskEvent is the interface for batch job event type BatchTaskEvent interface {
type BatchTaskEvent interface {
	GetTask() *BatchTask
	SetTask(*BatchTask)
}

var _ BatchTaskEvent = &BatchTaskStateChangedEvent{}

// GetTask returns the task.
func (x *BatchTaskStateChangedEvent) GetTask() *BatchTask {
	return &BatchTask{
		Name: x.TaskName,
		Status: &BatchTaskStatus{
			State: BatchTaskStatus_State(x.NewTaskState),
		},
	}
}

// SetTask implements BatchTaskEvent.
func (x *BatchTaskStateChangedEvent) SetTask(task *BatchTask) {
	x.TaskName = task.Name
	x.NewTaskState = BatchTaskStateChangedEvent_State(task.Status.GetState())
}

// BatchTaskStateChangedEvent_State is the enum type of BatchTaskStateChangedEvent.State.
func NewBatchTaskStateChangedEvent_State(text string) BatchTaskStateChangedEvent_State {
	return BatchTaskStateChangedEvent_State(BatchTaskStateChangedEvent_State_value[text])
}

//counterfeiter:generate -o batchv1fake . BatchTaskEventHandler

// BatchTaskEventHandler handles connect.batch.v1.BatchTaskEvent event.
type BatchTaskEventHandler interface {
	// HandleBatchTaskEvent handles connect.batch.v1.BatchTaskEvent event.
	HandleBatchTaskEvent(context.Context, BatchTaskEvent) error
}
