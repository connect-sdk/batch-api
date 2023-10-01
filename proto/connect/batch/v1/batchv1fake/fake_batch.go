package batchv1fake

import (
	"fmt"
	"time"

	batchpb "cloud.google.com/go/batch/apiv1/batchpb"
	batchv1 "github.com/connect-sdk/batch-api/proto/connect/batch/v1"
)

func NewFakeBatchJob() *batchpb.Job {
	name := batchv1.ParsedJobName{
		ProjectID:  "prj-ci-fake",
		LocationID: "us-central1",
		JobID:      fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	return &batchpb.Job{
		Name: name.String(),
		Uid:  name.JobID,
		Status: &batchpb.JobStatus{
			State: batchpb.JobStatus_SUCCEEDED,
		},
	}
}
