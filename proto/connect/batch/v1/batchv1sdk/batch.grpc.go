package batchv1sdk

import (
	"context"

	batch "cloud.google.com/go/batch/apiv1"
	batchpb "cloud.google.com/go/batch/apiv1/batchpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	iterator "google.golang.org/api/iterator"
	option "google.golang.org/api/option"

	batchv1 "github.com/connect-sdk/batch-api/proto/connect/batch/v1"
	gax "github.com/googleapis/gax-go/v2"
)

// WithCredentialsFile returns a ClientOption that authenticates API calls
// with the given service account or refresh token JSON credentials file.
var BatchServiceClientWithCredentialsFile = option.WithCredentialsFile

var _ batchv1.BatchServiceClient = &BatchServiceClient{}

// BatchServiceClient is a client for interacting with Connect Batch API.
type BatchServiceClient struct {
	client *batch.Client
}

// NewBatchServiceClient creates a new BatchServiceClient.
func NewBatchServiceClient(ctx context.Context, options ...option.ClientOption) (*BatchServiceClient, error) {
	client, err := batch.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &BatchServiceClient{client: client}, nil
}

// CancelOperation implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) CancelOperation(ctx context.Context, r *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return x.client.CancelOperation(ctx, r, opts...)
}

// CreateJob implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) CreateJob(ctx context.Context, r *batchpb.CreateJobRequest, opts ...gax.CallOption) (*batchpb.Job, error) {
	return x.client.CreateJob(ctx, r, opts...)
}

// DeleteJob implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) DeleteJob(ctx context.Context, r *batchpb.DeleteJobRequest, opts ...gax.CallOption) (batchv1.DeleteBatchJobOperation, error) {
	return x.client.DeleteJob(ctx, r, opts...)
}

// DeleteJobOperation implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) DeleteJobOperation(name string) batchv1.DeleteBatchJobOperation {
	return x.client.DeleteJobOperation(name)
}

// DeleteOperation implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) DeleteOperation(ctx context.Context, r *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return x.client.DeleteOperation(ctx, r, opts...)
}

// GetJob implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) GetJob(ctx context.Context, r *batchpb.GetJobRequest, opts ...gax.CallOption) (*batchpb.Job, error) {
	return x.client.GetJob(ctx, r, opts...)
}

// GetTask implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) GetTask(ctx context.Context, r *batchpb.GetTaskRequest, opts ...gax.CallOption) (*batchpb.Task, error) {
	return x.client.GetTask(ctx, r, opts...)
}

// ListJobs implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) ListJobs(ctx context.Context, r *batchpb.ListJobsRequest, opts ...gax.CallOption) batchv1.BatchJobIterator {
	traverser := x.client.ListJobs(ctx, r, opts...)
	return &BatchJobIterator{iterator: traverser}
}

// ListTasks implements batchv1.BatchServiceClient.
func (x *BatchServiceClient) ListTasks(ctx context.Context, r *batchpb.ListTasksRequest, opts ...gax.CallOption) batchv1.BatchTaskIterator {
	traverser := x.client.ListTasks(ctx, r, opts...)
	return &BatchTaskIterator{iterator: traverser}
}

var _ batchv1.BatchJobIterator = &BatchJobIterator{}

// BatchJobIterator is a wrapper around batch.JobIterator.
type BatchJobIterator struct {
	iterator *batch.JobIterator
}

// Next implements cdk.BatchJobIterator.
func (x *BatchJobIterator) Next() (*batchpb.Job, error) {
	return x.iterator.Next()
}

// PageInfo implements cdk.BatchJobIterator.
func (x *BatchJobIterator) PageInfo() batchv1.BatchPageInfo {
	info := x.iterator.PageInfo()
	return &BatchPageInfo{info: info}
}

var _ batchv1.BatchTaskIterator = &BatchTaskIterator{}

// BatchTaskIterator is a wrapper around batch.TaskIterator.
type BatchTaskIterator struct {
	iterator *batch.TaskIterator
}

// Next implements cdk.BatchTaskIterator.
func (x *BatchTaskIterator) Next() (*batchpb.Task, error) {
	return x.iterator.Next()
}

// PageInfo implements cdk.BatchTaskIterator.
func (x *BatchTaskIterator) PageInfo() batchv1.BatchPageInfo {
	info := x.iterator.PageInfo()
	return &BatchPageInfo{info: info}
}

var _ batchv1.BatchPageInfo = &BatchPageInfo{}

// BatchPageInfo is a wrapper around batch.PageInfo.
type BatchPageInfo struct {
	info *iterator.PageInfo
}

// MaxSize implements batchv1.BatchPageInfo.
func (x *BatchPageInfo) MaxSize() int {
	return x.info.MaxSize
}

// Remaining implements batchv1.BatchPageInfo.
func (x *BatchPageInfo) Remaining() int {
	return x.info.Remaining()
}

// Token implements batchv1.BatchPageInfo.
func (x *BatchPageInfo) Token() string {
	return x.info.Token
}
