package main

import (
	"authorizer/handler"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	tableName := os.Getenv("TABLE")
	if tableName == "" {
		logger.Fatal("TABLE is not set")
	}

	handler := handler.New(svc, tableName, logger)
	lambda.Start(handler.HandleRequest)
}
