package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleSQSMessage(ctx context.Context, sqsEvent events.SQSEvent) error {
	fmt.Println("===============================================================================")
	fmt.Printf("%+v\n", sqsEvent)
	fmt.Println("===============================================================================")

	for _, message := range sqsEvent.Records {
		fmt.Printf("Message ID: %s\n", message.MessageId)
		fmt.Printf("Message Body: %s\n", message.Body)
		fmt.Printf("message.EventSource: %v\n", message.EventSource)
		fmt.Printf("message.EventSourceARN: %v\n", message.EventSourceARN)
		fmt.Printf("message.AWSRegion: %v\n", message.AWSRegion)

		for k, v := range message.Attributes {
			fmt.Printf("message.Attributes[%s]: %v\n", k, v)
		}

		for k, v := range message.MessageAttributes {
			fmt.Println("...........................................................................")
			if v.StringValue != nil {
				fmt.Printf("message.MessageAttributes[%s].StringValue: %v\n", k, *v.StringValue)
			}

			if v.BinaryValue != nil {
				fmt.Printf("message.MessageAttributes[%s].BinaryValue: %v\n", k, v.BinaryValue)
			}

			if v.StringListValues != nil {
				fmt.Printf("message.MessageAttributes[%s].StringListValues: %v\n", k, v.StringListValues)
			}

			if v.BinaryListValues != nil {
				fmt.Printf("message.MessageAttributes[%s].BinaryListValues: %v\n", k, v.BinaryListValues)
			}

			fmt.Printf("message.MessageAttributes[%s].DataType: %v\n", k, v.DataType)

			fmt.Println("...........................................................................")
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleSQSMessage)
}
