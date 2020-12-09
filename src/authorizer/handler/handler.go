package handler

import (
	"authorizer/policy"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/sirupsen/logrus"
)

// AuthHandler implements a dummy API gateway authorizer
type AuthHandler struct {
	logger    *logrus.Logger
	ddb       dynamodbiface.DynamoDBAPI
	tableName string
}

type token struct {
	Name  string
	Email string
}

// New returns a new API gateway authorizer handler
func New(ddb dynamodbiface.DynamoDBAPI, tableName string, logger *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		logger:    logger,
		ddb:       ddb,
		tableName: tableName,
	}
}

// HandleRequest processes a single API gateway authorizer request
func (h *AuthHandler) HandleRequest(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	// Extract token from "Bearer <token>" string
	tokenSlice := strings.Split(request.AuthorizationToken, " ")
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[len(tokenSlice)-1]
	}

	// Lookup token in DynamoDB to see if it exists
	token, err := h.getToken(ctx, bearerToken)
	if err != nil {
		// Not found or any other error: unauthorized!
		h.logger.WithError(err).Info("Unauthorized")
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	// Return policy that authorizes the request
	return policy.Generate(token.Name, "Allow", request.MethodArn), nil
}

func (h *AuthHandler) getToken(ctx context.Context, id string) (*token, error) {
	result, err := h.ddb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(h.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Token": {S: aws.String(id)},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, fmt.Errorf("Not found: %s", id)
	}
	var t token
	if err := dynamodbattribute.UnmarshalMap(result.Item, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
