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
	ParseJobName() (ParsedJobName, error)
	GetJobName() string
	SetJobName(string)
	GetJobUid() string
	SetJobUid(string)
}

var _ BatchJobEvent = &BatchJobStateChangedEvent{}

// SetJobName implements BatchJobEvent.
func (x *BatchJobStateChangedEvent) SetJobName(name string) {
	x.JobName = name
}

// SetJobUid implements BatchJobEvent.
func (x *BatchJobStateChangedEvent) SetJobUid(uid string) {
	x.JobUid = uid
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
	ParseTaskName() (ParsedTaskName, error)
	GetTaskName() string
	SetTaskName(string)
	GetTaskUid() string
	SetTaskUid(string)
}

var _ BatchTaskEvent = &BatchTaskStateChangedEvent{}

// SetTaskName implements BatchTaskEvent.
func (x *BatchTaskStateChangedEvent) SetTaskName(name string) {
	x.TaskName = name
}

// SetTaskUid implements BatchTaskEvent.
func (x *BatchTaskStateChangedEvent) SetTaskUid(uid string) {
	x.TaskUid = uid
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
