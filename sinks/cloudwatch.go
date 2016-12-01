package sinks

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ejholmes/cloudwatch"
	"github.com/pborman/uuid"
)

// CloudwatchSink represents a filesystem sink

type CloudwatchSink struct {
	writer *io.Writer
}

// NewCloudwatchSink creates a CloudwatchSink instance
func NewCloudwatchSink(logGroup string) (Sink, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	group := cloudwatch.NewGroup(logGroup, cloudwatchlogs.New(sess))
	stream := uuid.New()
	writer, err := group.Create(stream)
	if err != nil {
		return nil, err
	}

	return Sink(&CloudwatchSink{writer: &writer}), nil
}

// Write writes to a file
func (cs *CloudwatchSink) Write(output []byte) error {
	_, err := (*cs.writer).Write(output)
	return err
}
