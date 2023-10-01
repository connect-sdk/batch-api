package batchv1sdk

import (
	"context"

	connect "connectrpc.com/connect"
	pubsubv1 "github.com/connect-sdk/pubsub-api/proto/connect/pubsub/v1"

	batchv1 "github.com/connect-sdk/batch-api/proto/connect/batch/v1"
)

var _ pubsubv1.PubsubService = &BatchPubsubService{}

// BatchPubsubService represents a batch pubsub service.
type BatchPubsubService struct {
	// BatchJobEventHandler is a batch job event handler.
	BatchJobEventHandler batchv1.BatchJobEventHandler
	// BatchTaskEventHandler is a batch task event handler.
	BatchTaskEventHandler batchv1.BatchTaskEventHandler
}

// PushPubsubMessage implements pubsubv1.PubsubService.
func (x *BatchPubsubService) PushPubsubMessage(ctx context.Context, r *pubsubv1.PushPubsubMessageRequest) (*pubsubv1.PushPubsubMessageResponse, error) {
	switch r.Message.Attributes["Type"] {
	case batchv1.BatchJobNotification_JOB_STATE_CHANGED.String():
		event := &batchv1.BatchJobStateChangedEvent{
			JobName:     r.Message.Attributes["JobName"],
			JobUid:      r.Message.Attributes["JobUID"],
			NewJobState: batchv1.NewBatchJobStateChangedEvent_State(r.Message.Attributes["NewJobState"]),
			OldJobState: batchv1.NewBatchJobStateChangedEvent_State(r.Message.Attributes["OldJobState"]),
		}

		if err := x.BatchJobEventHandler.HandleBatchJobEvent(ctx, event); err != nil {
			return nil, err
		}

	case batchv1.BatchJobNotification_TASK_STATE_CHANGED.String():
		event := &batchv1.BatchTaskStateChangedEvent{
			TaskName:     r.Message.Attributes["TaskName"],
			TaskUid:      r.Message.Attributes["TaskUID"],
			NewTaskState: batchv1.NewBatchTaskStateChangedEvent_State(r.Message.Attributes["NewTaskState"]),
			OldTaskState: batchv1.NewBatchTaskStateChangedEvent_State(r.Message.Attributes["OldTaskState"]),
		}

		if err := x.BatchTaskEventHandler.HandleBatchTaskEvent(ctx, event); err != nil {
			return nil, err
		}
	default:
		return nil, connect.NewError(connect.CodeInternal, batchv1.ErrBatchEventTypeUnknown)
	}

	return &pubsubv1.PushPubsubMessageResponse{}, nil
}
