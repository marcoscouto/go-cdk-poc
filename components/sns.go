package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/jsii-runtime-go"
)

var (
	firstTopic  awssns.Topic
	secondTopic awssns.Topic
)

func ConfigTopics(stack *awscdk.Stack) {

	firstTopic = awssns.NewTopic(*stack, jsii.String("topic-1"), &awssns.TopicProps{
		ContentBasedDeduplication: jsii.Bool(true),
		DisplayName:               jsii.String("topic for test 1"),
		Fifo:                      jsii.Bool(true),
		TopicName:                 jsii.String("topic-test-1"),
	})

	secondTopic = awssns.NewTopic(*stack, jsii.String("topic-2"), &awssns.TopicProps{
		ContentBasedDeduplication: jsii.Bool(true),
		DisplayName:               jsii.String("topic for test 2"),
		Fifo:                      jsii.Bool(true),
		TopicName:                 jsii.String("topic-test-2"),
	})

}
