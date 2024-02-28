package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/goombaio/namegenerator"
)

type MessageAttributes struct {
	MyIntField    int    `json:"myIntField"`
	MyStringField string `json:"myStringField"`
}

func PublishToSNS(ctx context.Context) error {

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	message := nameGenerator.Generate()

	// Retrieve SNS topic ARN from environment variable
	snsTopicArn := os.Getenv("SNS_TOPIC_ARN")
	if snsTopicArn == "" {
		return errors.New("SNS_TOPIC_ARN environment variable is not set")
	}

	// Initialize AWS session
	conf, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("Error loading AWS configuration: %v", err)
	}

	// Create SNS service client
	snsSvc := sns.NewFromConfig(conf)

	v := 123
	if time.Now().UTC().UnixNano()%2 == 0 {
		v = 456
	}

	messageAttributes := MessageAttributes{
		MyIntField:    v,
		MyStringField: "Hello World",
	}

	// Publish message to SNS
	_, err = snsSvc.Publish(ctx, &sns.PublishInput{
		TopicArn:          aws.String(snsTopicArn),
		Message:           aws.String(message),
		MessageAttributes: getMessageAttributes(messageAttributes),
	})
	if err != nil {
		return fmt.Errorf("Error publishing message to SNS: %v", err)
	}

	fmt.Printf("Published message to SNS [%v]\n", messageAttributes)

	return nil
}

func getMessageAttributes(messageAttributes MessageAttributes) map[string]types.MessageAttributeValue {
	return map[string]types.MessageAttributeValue{
		"myIntField": {
			DataType:    aws.String("Number"),
			StringValue: aws.String(fmt.Sprintf("%d", messageAttributes.MyIntField)),
		},
		"myStringField": {
			DataType:    aws.String("String"),
			StringValue: aws.String(messageAttributes.MyStringField),
		},
	}
}

func main() {
	lambda.Start(PublishToSNS)
}
