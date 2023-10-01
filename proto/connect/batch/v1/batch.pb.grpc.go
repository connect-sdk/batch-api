package batchv1

import (
	"context"

	"cloud.google.com/go/batch/apiv1/batchpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/googleapis/gax-go"
)

//go:generate counterfeiter -generate

//counterfeiter:generate -o batchv1fake . BatchServiceClient

// BatchServiceClient is the client API for BatchService service.
type BatchServiceClient interface {
	CreateJob(context.Context, *CreateBatchJobRequest, ...gax.CallOption) (*BatchJob, error)
	GetJob(context.Context, *GetBatchJobRequest, ...gax.CallOption) (*BatchJob, error)
	DeleteJob(context.Context, *DeleteBatchJobRequest, ...gax.CallOption) (DeleteBatchJobOperation, error)
	DeleteJobOperation(string) DeleteBatchJobOperation
	ListJobs(context.Context, *ListBatchJobsRequest, ...gax.CallOption) BatchJobIterator
	GetTask(context.Context, *GetBatchTaskRequest, ...gax.CallOption) (*BatchTask, error)
	ListTasks(context.Context, *ListBatchTasksRequest, ...gax.CallOption) BatchTaskIterator
	CancelOperation(context.Context, *CancelBatchOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *DeleteBatchOperationRequest, ...gax.CallOption) error
}

type (
	BatchJob             = batchpb.Job
	BatchJobStatus       = batchpb.JobStatus
	BatchJobStatus_State = batchpb.JobStatus_State
)

var (
	BatchJobStatus_STATE_UNSPECIFIED = batchpb.JobStatus_STATE_UNSPECIFIED
	// Notify users that the job state has changed.
	BatchJobNotification_JOB_STATE_CHANGED = batchpb.JobNotification_JOB_STATE_CHANGED
)

type (
	BatchTask             = batchpb.Task
	BatchTaskStatus       = batchpb.TaskStatus
	BatchTaskStatus_State = batchpb.TaskStatus_State
)

var (
	BatchTaskStatus_STATE_UNSPECIFIED = batchpb.TaskStatus_STATE_UNSPECIFIED
	// Notify users that the task state has changed.
	BatchJobNotification_TASK_STATE_CHANGED = batchpb.JobNotification_TASK_STATE_CHANGED
)

type (
	CreateBatchJobRequest = batchpb.CreateJobRequest
	GetBatchJobRequest    = batchpb.GetJobRequest
	DeleteBatchJobRequest = batchpb.DeleteJobRequest
	ListBatchJobsRequest  = batchpb.ListJobsRequest
	GetBatchTaskRequest   = batchpb.GetTaskRequest
	ListBatchTasksRequest = batchpb.ListTasksRequest
)

type (
	BatchOperation              = longrunningpb.Operation
	CancelBatchOperationRequest = longrunningpb.CancelOperationRequest
	DeleteBatchOperationRequest = longrunningpb.DeleteOperationRequest
)

//counterfeiter:generate -o batchv1fake . DeleteBatchJobOperation

// DeleteJobOperation manages a long-running operation from DeleteJob.
type DeleteBatchJobOperation interface {
	Name() string
	Wait(ctx context.Context, opts ...gax.CallOption) error
	Poll(ctx context.Context, opts ...gax.CallOption) error
	Metadata() (*batchpb.OperationMetadata, error)
	Done() bool
}

//counterfeiter:generate -o batchv1fake . BatchPageInfo

// PageInfo contains information about an iterator's paging state.
type BatchPageInfo interface {
	Token() string
	MaxSize() int
	Remaining() int
}

//counterfeiter:generate -o batchv1fake . BatchJobIterator

type BatchJobIterator interface {
	Next() (*batchpb.Job, error)
	PageInfo() BatchPageInfo
}

//counterfeiter:generate -o batchv1fake . BatchTaskIterator

type BatchTaskIterator interface {
	Next() (*batchpb.Task, error)
	PageInfo() BatchPageInfo
}

//counterfeiter:generate -o batchv1fake . BatchOperationIterator

type BatchOperationIterator interface {
	Next() (*BatchOperation, error)
	PageInfo() BatchPageInfo
}
