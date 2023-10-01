package batchv1

import (
	"context"

	"cloud.google.com/go/batch/apiv1/batchpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/timestamppb"
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

var _ BatchServiceClient = &NopBatchServiceClient{}

// NopBatchServiceClient is a no-op implementation of BatchServiceClient.
type NopBatchServiceClient struct{}

// CancelOperation implements BatchServiceClient.
func (*NopBatchServiceClient) CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error {
	return nil
}

// CreateJob implements BatchServiceClient.
func (*NopBatchServiceClient) CreateJob(context.Context, *batchpb.CreateJobRequest, ...gax.CallOption) (*batchpb.Job, error) {
	return &batchpb.Job{}, nil
}

// DeleteJob implements BatchServiceClient.
func (*NopBatchServiceClient) DeleteJob(context.Context, *batchpb.DeleteJobRequest, ...gax.CallOption) (DeleteBatchJobOperation, error) {
	panic("unimplemented")
}

// DeleteJobOperation implements BatchServiceClient.
func (*NopBatchServiceClient) DeleteJobOperation(string) DeleteBatchJobOperation {
	panic("unimplemented")
}

// DeleteOperation implements BatchServiceClient.
func (*NopBatchServiceClient) DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error {
	return nil
}

// GetJob implements BatchServiceClient.
func (*NopBatchServiceClient) GetJob(context.Context, *batchpb.GetJobRequest, ...gax.CallOption) (*batchpb.Job, error) {
	return &batchpb.Job{
		Name:             "",
		Uid:              "",
		Priority:         0,
		TaskGroups:       []*batchpb.TaskGroup{},
		AllocationPolicy: &batchpb.AllocationPolicy{},
		Labels:           map[string]string{},
		Status:           &batchpb.JobStatus{},
		CreateTime:       &timestamppb.Timestamp{},
		UpdateTime:       &timestamppb.Timestamp{},
		LogsPolicy:       &batchpb.LogsPolicy{},
		Notifications:    []*batchpb.JobNotification{},
	}, nil
}

// GetTask implements BatchServiceClient.
func (*NopBatchServiceClient) GetTask(context.Context, *batchpb.GetTaskRequest, ...gax.CallOption) (*batchpb.Task, error) {
	return &batchpb.Task{}, nil
}

// ListJobs implements BatchServiceClient.
func (*NopBatchServiceClient) ListJobs(context.Context, *batchpb.ListJobsRequest, ...gax.CallOption) BatchJobIterator {
	return &NopBatchJobIterator{}
}

// ListTasks implements BatchServiceClient.
func (*NopBatchServiceClient) ListTasks(context.Context, *batchpb.ListTasksRequest, ...gax.CallOption) BatchTaskIterator {
	return &NopBatchTaskIterator{}
}

var _ DeleteBatchJobOperation = &NopDeleteBatchJobOperation{}

// NopDeleteBatchJobOperation is a no-op implementation of DeleteBatchJobOperation.
type NopDeleteBatchJobOperation struct{}

// Done implements DeleteBatchJobOperation.
func (*NopDeleteBatchJobOperation) Done() bool {
	return true
}

// Metadata implements DeleteBatchJobOperation.
func (*NopDeleteBatchJobOperation) Metadata() (*batchpb.OperationMetadata, error) {
	return &batchpb.OperationMetadata{}, nil
}

// Name implements DeleteBatchJobOperation.
func (*NopDeleteBatchJobOperation) Name() string {
	return ""
}

// Poll implements DeleteBatchJobOperation.
func (*NopDeleteBatchJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return nil
}

// Wait implements DeleteBatchJobOperation.
func (*NopDeleteBatchJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return nil
}

var _ BatchJobIterator = &NopBatchJobIterator{}

// NopBatchJobIterator is a no-op implementation of BatchJobIterator.
type NopBatchJobIterator struct{}

// Next implements BatchJobIterator.
func (*NopBatchJobIterator) Next() (*batchpb.Job, error) {
	return nil, iterator.Done
}

// PageInfo implements BatchJobIterator.
func (*NopBatchJobIterator) PageInfo() BatchPageInfo {
	return &NoBatchPageInfo{}
}

var _ BatchTaskIterator = &NopBatchTaskIterator{}

// NopBatchTaskIterator is a no-op implementation of BatchTaskIterator.
type NopBatchTaskIterator struct{}

// Next implements BatchTaskIterator.
func (*NopBatchTaskIterator) Next() (*batchpb.Task, error) {
	return nil, iterator.Done
}

// PageInfo implements BatchTaskIterator.
func (*NopBatchTaskIterator) PageInfo() BatchPageInfo {
	return &NoBatchPageInfo{}
}

var _ BatchPageInfo = &NoBatchPageInfo{}

// NoBatchPageInfo is a no-op implementation of BatchPageInfo.
type NoBatchPageInfo struct{}

// MaxSize implements BatchPageInfo.
func (*NoBatchPageInfo) MaxSize() int {
	return 0
}

// Remaining implements BatchPageInfo.
func (*NoBatchPageInfo) Remaining() int {
	return 0
}

// Token implements BatchPageInfo.
func (*NoBatchPageInfo) Token() string {
	return ""
}
