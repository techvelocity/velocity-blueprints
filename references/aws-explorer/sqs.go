package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

func testSQS() {
	sqsSvc := sqs.New(awsSession)

	sqsResult, err := sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(mustGetenv("QUEUE_NAME")),
	})
	if err != nil {
		log.Fatalf("sqsSvc.GetQueueUrl: %v", err)
	}

	log.Printf("Queue URL: %v\n", *sqsResult.QueueUrl)
}
