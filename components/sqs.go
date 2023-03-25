package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

var (
	firstQueue  awssqs.Queue
	secondQueue awssqs.Queue
)

func ConfigQueues(stack *awscdk.Stack) {

	firstQueueDlq := awssqs.NewQueue(*stack, jsii.String("queue-for-test-1-dlq"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Fifo:              jsii.Bool(true),
	})

	secondQueueDlq := awssqs.NewQueue(*stack, jsii.String("queue-for-test-2-dlq"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Fifo:              jsii.Bool(true),
	})

	maxReceiveCount := 3.0

	firstQueue = awssqs.NewQueue(*stack, jsii.String("queue-for-test-1"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Fifo:              jsii.Bool(true),
		DeadLetterQueue:   &awssqs.DeadLetterQueue{MaxReceiveCount: &maxReceiveCount, Queue: firstQueueDlq},
	})

	secondQueue = awssqs.NewQueue(*stack, jsii.String("queue-for-test-2"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Fifo:              jsii.Bool(true),
		DeadLetterQueue:   &awssqs.DeadLetterQueue{MaxReceiveCount: &maxReceiveCount, Queue: secondQueueDlq},
	})

}
